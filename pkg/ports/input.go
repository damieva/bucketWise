package ports

import (
	"bucketWise/pkg/domain"
)

type CategoryService interface {
	Create(cat domain.Category) (interface{}, error)
	List(name string) ([]domain.Category, error)
	Delete(cat domain.Category) (int64, error)
	Update(catName string, cat domain.Category) (int64, error)
}

type TransactionService interface {
	Create(tx domain.Transaction) (interface{}, error)
	List(cat string) ([]domain.Transaction, error)
}

type CategoryUseCase interface {
	CreateCategoryUseCase(cat domain.Category) (domain.Category, error)
	ListCategoriesUseCase(name string) ([]domain.Category, error)
	DeleteCategoryUseCase(cat domain.Category) (int64, error)
	UpdateCategoryUseCase(catName string, cat domain.Category) (int64, error)
}

type TransactionUseCase interface {
	CreateTransactionUseCase(tx domain.Transaction) (domain.Transaction, error)
	ListTransactionsUseCase(cat string) ([]domain.Transaction, error)
}
