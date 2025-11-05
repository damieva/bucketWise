package ports

import "bucketWise/pkg/domain"

type CategoryRepository interface {
	Insert(category domain.Category) (interface{}, error)
	SelectAll() ([]domain.Category, error)
	Delete(category domain.Category) (int64, error)
	SelectOne(cat domain.Category) (domain.Category, error)
	Update(catName string, cat domain.Category) (int64, error)
}

type TransactionRepository interface {
	Insert(tx domain.Transaction) (interface{}, error)
	Select(cat string) ([]domain.Transaction, error)
}
