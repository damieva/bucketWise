package ports

import (
	"bucketWise/pkg/domain"
)

type CategoryUseCase interface {
	CreateCategoryUseCase(cat domain.Category) (domain.Category, error)
	ListCategoriesUseCase(name string) ([]domain.Category, error)
	DeleteCategoryUseCase(IDs []domain.ID) (int64, error)
	UpdateCategoryUseCase(catName string, cat domain.Category) (int64, error)
}

type TransactionUseCase interface {
	CreateTransactionUseCase(tx domain.Transaction) (domain.Transaction, error)
	ListTransactionsUseCase(cat string) ([]domain.Transaction, error)
	DeleteTransactionsUseCase(IDs []domain.ID) (int64, error)
}
