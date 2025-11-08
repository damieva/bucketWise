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

type TransactionRepo struct {
	Client *mongo.Client
}

func (r TransactionRepo) Select(cat string) ([]domain.Transaction, error) {
	collection := r.Client.Database("bucketWise").Collection("transactions")

	// Filtra por categoryName y si categoryName es vacío devuelve todas las tx
	var filter bson.M
	if cat == "" {
		filter = bson.M{}
	} else {
		filter = bson.M{"categoryName": cat}
	}
	log.Println(cat)
	log.Println(filter)
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("error getting the requested transaction documents %v", err)
		return nil, domain.ErrUnexpectedDatabase
	}

	var transactions []domain.Transaction
	err = cursor.All(context.Background(), &transactions)
	log.Println(transactions)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error converting mongo documents from the cursor to a transactions array %w", err)
	}

	return transactions, nil
}

func (r TransactionRepo) Insert(tx domain.Transaction) (interface{}, error) {
	// Inicializamos un handler para trabajar con la collection transactions
	collection := r.Client.Database("bucketWise").Collection("transactions")
	// Insertamos un documento en la collection. El contexto (como bien inicializamos arriba) indica el tiempo y cancelación de la operación.
	// El insertResult nos devolverá el ID que Mongo asignará al documento
	insertResult, err := collection.InsertOne(context.Background(), tx)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error inserting transaction %w", err)
	}

	return insertResult.InsertedID, nil
}

func (r TransactionRepo) Delete(IDs []string) (int64, error) {
	collection := r.Client.Database("bucketWise").Collection("transactions")

	objectIDs := make([]primitive.ObjectID, 0, len(IDs))
	for _, id := range IDs {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return 0, fmt.Errorf("invalid transaction ID: %s", id)
		}
		objectIDs = append(objectIDs, objID)
	}

	filter := bson.M{"_id": bson.M{"$in": objectIDs}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Printf("error deleting transactions: %v", err)
		return 0, fmt.Errorf("error deleting transactions: %w", err)
	}

	return deleteResult.DeletedCount, nil
}
