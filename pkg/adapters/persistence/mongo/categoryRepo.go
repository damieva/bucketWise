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

func (r CategoryRepo) Insert(cat domain.Category) (domain.ID, error) {
	collection := r.Client.Database("bucketWise").Collection("categories")

	insertResult, err := collection.InsertOne(context.Background(), cat)
	if err != nil {
		log.Printf("error inserting category: %v", err)
		return "", domain.ErrUnexpectedDatabase
	}

	// Convertir el ID devuelto por Mongo a primitive.ObjectID
	oid, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("expected ObjectID but got %T", insertResult.InsertedID)
	}

	// Conversion directa a domain.ID
	return domain.ID(oid.Hex()), nil
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
		log.Println(err.Error())
		return nil, fmt.Errorf("%w: %v", domain.ErrEntityDecoding, err)
	}

	// Si se buscaba una sola categoría y no existe, devolvemos un error semántico
	if name != "" && len(categories) == 0 {
		return nil, domain.ErrCategoryNotFound
	}

	return categories, nil
}

func (r CategoryRepo) Delete(IDs []domain.ID) (int64, error) {
	collection := r.Client.Database("bucketWise").Collection("categories")

	// Convertir los IDs de domain.ID a ObjectID
	objectIDs := make([]primitive.ObjectID, 0, len(IDs))
	for _, id := range IDs {
		objID, err := primitive.ObjectIDFromHex(string(id))
		if err != nil {
			log.Printf("invalid ObjectID: %s, error: %+v\n", id, err)
			return 0, fmt.Errorf("invalid ObjectID: %s", id)
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
