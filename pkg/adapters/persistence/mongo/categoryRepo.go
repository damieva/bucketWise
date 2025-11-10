package mongo

import (
	"bucketWise/pkg/domain"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepo struct {
	Client *mongo.Client
}

func (r CategoryRepo) Insert(cat domain.Category) (interface{}, error) {
	// Inicializamos un handler para trabajar con la collection categories
	collection := r.Client.Database("bucketWise").Collection("categories")
	// Insertamos un documento en la collection. El contexto (como bien inicializamos arriba) indica el tiempo y cancelación de la operación.
	// El insertResult nos devolverá el ID que Mongo asignará al documento
	insertResult, err := collection.InsertOne(context.Background(), cat)
	if err != nil {
		log.Printf("error inserting category %w", err)
		return nil, domain.ErrUnexpectedDatabase
	}

	return insertResult.InsertedID, nil
}

func (r CategoryRepo) Select(name string) ([]domain.Category, error) {
	collection := r.Client.Database("bucketWise").Collection("categories")

	var filter bson.M
	if name == "" {
		filter = bson.M{}
	} else {
		filter = bson.M{"name": name}
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("error getting category documents %v", err)
		return nil, domain.ErrUnexpectedDatabase
	}

	var categories []domain.Category
	err = cursor.All(context.Background(), &categories)
	if err != nil {
		log.Printf("error decoding mongo cursor: %v", err)
		return nil, fmt.Errorf("error converting mongo documents to a category array: %w", err)
	}

	// Si se buscaba una sola categoría y no existe, devolvemos un error semántico
	if name != "" && len(categories) == 0 {
		return nil, domain.ErrCategoryNotFound
	}

	return categories, nil
}

func (r CategoryRepo) Delete(IDs []string) (int64, error) {
	collection := r.Client.Database("bucketWise").Collection("categories")

	// Convertir los IDs de tipo string a ObjectID
	objectIDs := make([]primitive.ObjectID, 0, len(IDs))
	for _, idStr := range IDs {
		objID, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			log.Printf("invalid ObjectID: %s, error: %+v\n", idStr, err)
			return 0, fmt.Errorf("invalid ObjectID: %s", idStr)
		}
		objectIDs = append(objectIDs, objID)
	}

	// Crear el filtro para eliminar todas las categorías con esos _id
	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	res, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Printf("error deleting categories %v: %+v\n", IDs, err)
		return 0, domain.ErrUnexpectedDatabase
	}

	return res.DeletedCount, nil
}

func (r CategoryRepo) Update(catName string, cat domain.Category) (int64, error) {
	collection := r.Client.Database("bucketWise").Collection("categories")
	filter := bson.M{"name": catName}
	result, err := collection.ReplaceOne(context.Background(), filter, cat)
	if err != nil {
		log.Printf("error updating category %s: %+v\n", cat.Name, err)
		return 0, domain.ErrUnexpectedDatabase
	}

	return result.ModifiedCount, nil
}
