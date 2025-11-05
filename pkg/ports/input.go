package ports

import (
	"bucketWise/pkg/domain"
)

type CategoryService interface {
	Create(cat domain.Category) (interface{}, error)
	ListAll() ([]domain.Category, error)
	Delete(cat domain.Category) (int64, error)
	ListOne(cat domain.Category) (domain.Category, error)
	Update(catName string, cat domain.Category) (int64, error)
}

type TransactionService interface {
	Create(tx domain.Transaction) (interface{}, error)
	List(cat string) ([]domain.Transaction, error)
}

type CategoryUseCase interface {
	CreateCategoryUseCase(cat domain.Category) (domain.Category, error)
	ListAllCategoryUseCase() ([]domain.Category, error)
	ListOneCategoryUseCase(cat domain.Category) (domain.Category, error)
	DeleteCategoryUseCase(cat domain.Category) (int64, error)
	UpdateCategoryUseCase(catName string, cat domain.Category) (int64, error)
}

type TransactionUseCase interface {
	CreateTransactionUseCase(tx domain.Transaction) (domain.Transaction, error)
	ListTransactionsUseCase(cat string) ([]domain.Transaction, error)
}
