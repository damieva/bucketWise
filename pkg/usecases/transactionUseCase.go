package usecases

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionUseCase struct {
	TransactionService ports.TransactionService
}

func (uc TransactionUseCase) CreateTransactionUseCase(tx domain.Transaction) (domain.Transaction, error) {
	// Crear una nueva transacción basada en la recibida
	newTx := tx

	// Añadir/forzar los campos de categoría
	newTx.CategoryID = "68fe71b6e080f7fd26ff6c87"
	newTx.CategoryName = "Fixed costs"
	newTx.Type = domain.ExpenseCategory

	// Llamar al servicio para crear la transacción
	createdID, err := uc.TransactionService.Create(newTx)
	if err != nil {
		return domain.Transaction{}, err
	}

	// Convertir el ID de MongoDB a string
	objectID, ok := createdID.(primitive.ObjectID)
	if !ok {
		return domain.Transaction{}, fmt.Errorf("expected primitive.ObjectID but got %T", createdID)
	}

	// Asignar el ID a la entidad de dominio
	newTx.ID = objectID.Hex()

	// Devolver la transacción con su nuevo ID
	return newTx, nil
}

func (uc TransactionUseCase) ListTransactionsUseCase(cat string) ([]domain.Transaction, error) {
	return uc.TransactionService.List(cat)
}
