package ports

import "bucketWise/pkg/domain"

type CategoryRepository interface {
	Insert(category domain.Category) (id interface{}, err error)
}

type TransactionRepository interface {
	Insert(transaction domain.Transaction) (id interface{}, err error)
}
