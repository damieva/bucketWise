package ports

import (
	"bucketWise/pkg/domain"
)

type CategoryService interface {
	Create(cat domain.Category) (domain.ID, error)
	List(name string) ([]domain.Category, error)
	Delete(IDs []domain.ID) (int64, error)
	Update(catName string, cat domain.Category) (int64, error)
}

type TransactionService interface {
	Create(tx domain.Transaction) (domain.ID, error)
	List(cat string) ([]domain.Transaction, error)
	Delete(IDs []domain.ID) (int64, error)
	ExistsByCategoryIDs(categoryIDs []domain.ID) (bool, error)
}

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
