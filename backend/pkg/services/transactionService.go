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

func (s TransactionService) Create(tx domain.Transaction) (domain.ID, error) {
	insertedID, err := s.Repo.Insert(tx)
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("error creating transaction: %w", err)
	}

	return insertedID, nil
}

func (s TransactionService) List(cat string) ([]domain.Transaction, error) {
	transactionList, err := s.Repo.Select(cat)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error listing the requested transactions %w", err)
	}

	return transactionList, nil
}

func (s TransactionService) Delete(IDs []domain.ID) (int64, error) {
	deletedCount, err := s.Repo.Delete(IDs)
	if err != nil {
		log.Println(err.Error())
		return deletedCount, fmt.Errorf("error deleting transaction %w", err)
	}

	return deletedCount, nil
}

func (s TransactionService) ExistsByCategoryIDs(categoryIDs []domain.ID) (bool, error) {
	return s.Repo.ExistsByCategoryIDs(categoryIDs)
}
