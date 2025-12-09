package ports

import (
	"bucketWise/pkg/domain"
)

type CategoryRepository interface {
	Insert(cat domain.Category) (domain.ID, error)
	Select(name string) ([]domain.Category, error)
	Delete(IDs []domain.ID) (int64, error)
	Update(catName string, cat domain.Category) (int64, error)
}

type TransactionRepository interface {
	Insert(tx domain.Transaction) (domain.ID, error)
	Select(cat string) ([]domain.Transaction, error)
	Delete(IDs []domain.ID) (int64, error)
	ExistsByCategoryIDs(categoryIDs []domain.ID) (bool, error)
}
