package mongo

import (
	"bucketWise/pkg/domain"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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
		log.Println(err.Error())
		return nil, fmt.Errorf("error inserting category %w", err)
	}

	return insertResult.InsertedID, nil
}

func (r CategoryRepo) SelectAll() ([]domain.Category, error) {
	collection := r.Client.Database("bucketWise").Collection("categories")

	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error getting all the category documents %w", err)
	}

	var categories []domain.Category
	err = cursor.All(context.Background(), &categories)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error converting mongo documents from the cursor to a category array %w", err)
	}

	return categories, nil
}
