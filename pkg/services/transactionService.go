package services

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
	"fmt"
	"log"
)

type TransactionService struct {
	Repo ports.TransactionRepository
}

func (s TransactionService) Create(tx domain.Transaction) (interface{}, error) {

	insertedId, err := s.Repo.Insert(tx)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating transaction %w", err)
	}

	return insertedId, nil
}
