package usecases

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/services"
	"context"
)

type CreateTransactionUseCase struct {
	transactionService services.TransactionService
}

func (uc CreateTransactionUseCase) Run(ctx context.Context, tx domain.Transaction) (id interface{}, err error) {
	// Aquí no hay lógica adicional: solo orquestamos
	return uc.transactionService.Create(tx)
}
