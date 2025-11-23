package usecases

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
)

type TransactionUseCase struct {
	TransactionService ports.TransactionService
}

func (uc TransactionUseCase) CreateTransactionUseCase(tx domain.Transaction) (domain.Transaction, error) {
	// Crear una nueva transacción basada en la recibida
	newTx := tx

	// Añadir/forzar los campos de categoría
	newTx.CategoryID = "691c3ac97db0505faae7015c"
	newTx.CategoryName = "Fixed costs"
	newTx.Type = domain.ExpenseCategory

	// Llamar al servicio para crear la transacción
	createdID, err := uc.TransactionService.Create(newTx)
	if err != nil {
		return domain.Transaction{}, err
	}

	// Asignar el ID devuelto por el servicio (ya es domain.ID)
	newTx.ID = createdID

	// Devolver la transacción con su nuevo ID
	return newTx, nil
}

func (uc TransactionUseCase) ListTransactionsUseCase(cat string) ([]domain.Transaction, error) {
	return uc.TransactionService.List(cat)
}

func (uc TransactionUseCase) DeleteTransactionsUseCase(IDs []domain.ID) (int64, error) {
	return uc.TransactionService.Delete(IDs)
}
