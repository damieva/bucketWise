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
	Create(tx domain.Transaction) (domain.Transaction, error)
	ListAll() ([]domain.Transaction, error)
	ListOne(cat domain.Transaction) (domain.Transaction, error)
}

type CategoryUseCase interface {
	CreateCategoryUseCase(cat domain.Category) (interface{}, error)
	ListAllCategoryUseCase() ([]domain.Category, error)
	ListOneCategoryUseCase(cat domain.Category) (domain.Category, error)
	DeleteCategoryUseCase(cat domain.Category) (int64, error)
	UpdateCategoryUseCase(catName string, cat domain.Category) (int64, error)
}

type TransactionUseCase interface {
	CreateTransactionUseCase(tx domain.Transaction) (domain.Transaction, error)
	ListAllTransactionUseCase() ([]domain.Transaction, error)
	ListOneTransactionUseCase(cat domain.Transaction) (domain.Transaction, error)
}
