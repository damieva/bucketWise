package services

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
)

type TransactionService struct {
	Repo ports.TransactionRepository
}

func (s TransactionService) Create(tx domain.Transaction) (domain.Transaction, error) {

	/*insertedId, err := s.Repo.Insert(tx)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating transaction %w", err)
	}

	return insertedId, nil*/
	return domain.Transaction{}, nil
}
