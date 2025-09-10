package ports

import "bucketWise/pkg/domain"

type TransactionService interface {
	Create(transaction domain.Transaction) (id interface{}, err error)
}

type TransactionRepository interface {
	Insert(transaction domain.Transaction) (id interface{}, err error)
}
