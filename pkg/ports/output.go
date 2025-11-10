package ports

import "bucketWise/pkg/domain"

type CategoryRepository interface {
	Insert(category domain.Category) (interface{}, error)
	Select(name string) ([]domain.Category, error)
	Delete(IDs []string) (int64, error)
	Update(catName string, cat domain.Category) (int64, error)
}

type TransactionRepository interface {
	Insert(tx domain.Transaction) (interface{}, error)
	Select(cat string) ([]domain.Transaction, error)
	Delete(IDs []string) (int64, error)
	ExistsByCategoryIDs(categoryIDs []string) (bool, error)
}
