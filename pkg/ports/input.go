package ports

import (
	"bucketWise/pkg/domain"
)

type CategoryService interface {
	Create(cat domain.Category) (interface{}, error)
	ListAll() ([]domain.Category, error)
	Delete(cat domain.Category) (int64, error)
	ListOne(cat domain.Category) (domain.Category, error)
}

type TransactionService interface {
	Create(tx domain.Transaction) (interface{}, error)
}

type CategoryUseCase interface {
	CreateCategoryUseCase(cat domain.Category) (interface{}, error)
	ListAllCategoryUseCase() ([]domain.Category, error)
	ListOneCategoryUseCase(cat domain.Category) (domain.Category, error)
	DeleteCategoryUseCase(cat domain.Category) (int64, error)
	//UpdateCategoryUseCase(cat domain.Category) (id interface{}, err error)
}

type TransactionUseCase interface {
	CreateTransactionUseCase(tx domain.Transaction) (interface{}, error)
}
