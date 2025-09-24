package services

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
	"fmt"
	"log"
	"time"
)

type TransactionService struct {
	Repo ports.TransactionRepository
}

func (s TransactionService) Create(transaction domain.Transaction) (interface{}, error) {
	// Set creation time
	transaction.CreationTime = time.Now().UTC()

	insertedId, err := s.Repo.Insert(transaction)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating transaction %w", err)
	}

	return insertedId, nil
}
