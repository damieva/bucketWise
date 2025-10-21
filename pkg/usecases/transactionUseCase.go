package usecases

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
)

type TransactionUseCase struct {
	transactionService ports.TransactionService
}

func (uc TransactionUseCase) CreateTransactionUseCase(tx domain.Transaction) (domain.Transaction, error) {
	// Fase inicial: asignación estática de categoría
	tx.CategoryID = "1"
	tx.CategoryName = "Fixed costs"
	tx.Type = domain.ExpenseCategory

	return uc.transactionService.Create(tx)
}
