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

	// Filtra por category_name y si category_name es vacío devuelve todas las tx
	var filter bson.M
	if cat == "" {
		filter = bson.M{}
	} else {
		filter = bson.M{"category_name": cat}
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

func (r TransactionRepo) Insert(tx domain.Transaction) (domain.ID, error) {
	collection := r.Client.Database("bucketWise").Collection("transactions")

	// Convertir CategoryID string → ObjectID
	categoryOID, err := primitive.ObjectIDFromHex(string(tx.CategoryID))
	if err != nil {
		return "", fmt.Errorf("invalid CategoryID '%s': %w", tx.CategoryID, err)
	}

	// Documento real que se guarda en Mongo (mapper infra → mongo)
	mongoTx := bson.M{
		"category_id":   categoryOID,
		"category_name": tx.CategoryName,
		"type":          tx.Type,
		"amount":        tx.Amount,
		"date":          tx.Date,
		"description":   tx.Description,
	}

	insertResult, err := collection.InsertOne(context.Background(), mongoTx)
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("error inserting transaction: %w", err)
	}

	// Extraer ID generado por Mongo
	oid, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("expected ObjectID but got %T", insertResult.InsertedID)
	}

	return domain.ID(oid.Hex()), nil
}

func (r TransactionRepo) Delete(IDs []domain.ID) (int64, error) {
	collection := r.Client.Database("bucketWise").Collection("transactions")

	objectIDs := make([]primitive.ObjectID, 0, len(IDs))
	for _, id := range IDs {
		objID, err := primitive.ObjectIDFromHex(string(id))
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

func (r TransactionRepo) ExistsByCategoryIDs(categoryIDs []domain.ID) (bool, error) {
	collection := r.Client.Database("bucketWise").Collection("transactions")

	objectIDs := make([]primitive.ObjectID, 0, len(categoryIDs))
	for _, id := range categoryIDs {
		oid, err := primitive.ObjectIDFromHex(string(id))
		if err != nil {
			return false, fmt.Errorf("invalid ObjectID: %s", id)
		}
		objectIDs = append(objectIDs, oid)
	}

	filter := bson.M{"category_id": bson.M{"$in": objectIDs}}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
