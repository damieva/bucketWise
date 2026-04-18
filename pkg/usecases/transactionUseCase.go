package usecases

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
	"fmt"
	"log"
)

type TransactionUseCase struct {
	TransactionRepo ports.TransactionRepository
}

func (uc TransactionUseCase) CreateTransactionUseCase(tx domain.Transaction) (domain.Transaction, error) {
	// Crear una nueva transacción basada en la recibida
	newTx := tx

	// Añadir/forzar los campos de categoría
	newTx.CategoryID = "691c3ac97db0505faae7015c"
	newTx.CategoryName = "Fixed costs"
	newTx.Type = domain.ExpenseCategory

	// Llamar al repo para crear la transacción
	createdID, err := uc.TransactionRepo.Insert(newTx)
	if err != nil {
		return domain.Transaction{}, err
	}

	// Asignar el ID devuelto por el repo (ya es domain.ID)
	newTx.ID = createdID

	// Devolver la transacción con su nuevo ID
	return newTx, nil
}

func (uc TransactionUseCase) ListTransactionsUseCase(cat string) ([]domain.Transaction, error) {
	transactionList, err := uc.TransactionRepo.Select(cat)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error listing the requested transactions %w", err)
	}

	return transactionList, nil
}

func (uc TransactionUseCase) DeleteTransactionsUseCase(IDs []domain.ID) (int64, error) {
	deletedCount, err := uc.TransactionRepo.Delete(IDs)
	if err != nil {
		log.Println(err.Error())
		return deletedCount, fmt.Errorf("error deleting transaction %w", err)
	}

	return deletedCount, nil
}
