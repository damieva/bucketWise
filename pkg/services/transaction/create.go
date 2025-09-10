package transaction

import (
	"bucketWise/pkg/domain"
	"fmt"
	"log"
	"time"
)

func (s Service) Create(transaction domain.Transaction) (id interface{}, err error) {
	// Set creation time
	transaction.CreationTime = time.Now().UTC()

	insertedId, err := s.Repo.Insert(transaction)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating transaction %w", err)
	}

	return insertedId, nil
}
