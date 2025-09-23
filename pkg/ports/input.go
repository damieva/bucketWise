package ports

import (
	"bucketWise/pkg/domain"
)

type CategoryService interface {
	Create(cat domain.Category) (id interface{}, err error)
	ListAll() ([]domain.Category, error)
}

type TransactionService interface {
	Create(tx domain.Transaction) (id interface{}, err error)
}

type CategoryUseCase interface {
	CreateCategoryUseCase(cat domain.Category) (id interface{}, err error)
	ListAllCategoryUseCase() (id interface{}, err error)
}

type TransactionUseCase interface {
	CreateTransactionUseCase(tx domain.Transaction) (id interface{}, err error)
}
