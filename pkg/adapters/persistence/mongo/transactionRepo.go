package mongo

import (
	"bucketWise/pkg/domain"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepo struct {
	Client *mongo.Client
}

func (r TransactionRepo) Insert(tx domain.Transaction) (interface{}, error) {
	// Inicializamos un handler para trabajar con la collection transactions
	collection := r.Client.Database("go-l").Collection("transactions")
	// Insertamos un documento en la collection. El contexto (como bien inicializamos arriba) indica el tiempo y cancelación de la operación.
	// El insertResult nos devolverá el ID que Mongo asignará al documento
	insertResult, err := collection.InsertOne(context.Background(), tx)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error inserting transaction %w", err)
	}

	return insertResult.InsertedID, nil
}

func (r TransactionRepo) SelectAll() ([]domain.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (r TransactionRepo) SelectOne(cat domain.Transaction) (domain.Category, error) {
	//TODO implement me
	panic("implement me")
}
