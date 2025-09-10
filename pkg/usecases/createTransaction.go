package usecases

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/services/transaction"
	"context"
)

type CreateTransactionUseCase struct {
	transactionService transaction.Service
}

func (uc CreateTransactionUseCase) Run(ctx context.Context, tx domain.Transaction) (id interface{}, err error) {
	// Aquí no hay lógica adicional: solo orquestamos
	return uc.transactionService.Create(tx)
}
