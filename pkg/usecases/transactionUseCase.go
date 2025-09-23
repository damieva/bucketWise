package usecases

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/services"
)

type TransactionUseCase struct {
	transactionService services.TransactionService
}

func (uc TransactionUseCase) CreateTransactionUseCase(tx domain.Transaction) (id interface{}, err error) {
	// Aquí no hay lógica adicional: solo orquestamos
	return uc.transactionService.Create(tx)
}
