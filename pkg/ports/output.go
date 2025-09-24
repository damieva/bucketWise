package ports

import "bucketWise/pkg/domain"

type CategoryRepository interface {
	Insert(category domain.Category) (interface{}, error)
	SelectAll() ([]domain.Category, error)
	Delete(category domain.Category) (int64, error)
	SelectOne(cat domain.Category) (domain.Category, error)
}

type TransactionRepository interface {
	Insert(tx domain.Transaction) (interface{}, error)
	SelectAll() ([]domain.Transaction, error)
	Delete(tx domain.Transaction) (int64, error)
	SelectOne(cat domain.Transaction) (domain.Category, error)
}
