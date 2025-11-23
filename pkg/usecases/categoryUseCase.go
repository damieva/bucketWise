package usecases

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
	"errors"
	"fmt"
)

type CategoryUseCase struct {
	CategoryService    ports.CategoryService
	TransactionService ports.TransactionService
}

func (uc CategoryUseCase) CreateCategoryUseCase(cat domain.Category) (domain.Category, error) {
	// Comprobar si la categoría ya existe
	_, err := uc.CategoryService.List(cat.Name)
	if err == nil {
		// Si no hay error, la categoría ya existe
		return domain.Category{}, domain.ErrCategoryAlreadyExists
	}

	// Si no se encuentra la categoría, crearla
	if errors.Is(err, domain.ErrCategoryNotFound) {
		createdID, err := uc.CategoryService.Create(cat)
		if err != nil {
			return domain.Category{}, err
		}

		// Crear una nueva categoría basada en la recibida
		newCat := cat

		// Asignar el ID devuelto por el servicio (ya es domain.ID)
		newCat.ID = createdID

		// Devolver la categoría con su nuevo ID
		return newCat, nil
	}

	// Si el error no es ErrCategoryNotFound, devolverlo tal cual
	return domain.Category{}, err
}

func (uc CategoryUseCase) ListCategoriesUseCase(name string) ([]domain.Category, error) {
	return uc.CategoryService.List(name)
}

func (uc CategoryUseCase) DeleteCategoryUseCase(IDs []domain.ID) (int64, error) {
	// Primero, verificar si alguna de las categorías tiene transacciones asociadas
	hasTx, err := uc.TransactionService.ExistsByCategoryIDs(IDs)
	if err != nil {
		return 0, fmt.Errorf("error checking transactions for categories: %w", err)
	}

	if hasTx {
		return 0, domain.ErrCategoryHasTransactions
	}

	// Si no hay transacciones, proceder a eliminar las categorías
	deletedCount, err := uc.CategoryService.Delete(IDs)
	if err != nil {
		return 0, err
	}

	return deletedCount, nil
}

func (uc CategoryUseCase) UpdateCategoryUseCase(catName string, cat domain.Category) (int64, error) {
	// Verify that the category to update actually exists in the database
	_, err := uc.CategoryService.List(catName)
	if errors.Is(err, domain.ErrCategoryNotFound) {
		return 0, domain.ErrCategoryNotFound
	} else if err != nil {
		return 0, err
	}

	// Ensure the new category data does not conflict with an existing record
	_, err = uc.CategoryService.List(cat.Name)
	if err == nil {
		return 0, domain.ErrCategoryAlreadyExists
	} else if errors.Is(err, domain.ErrCategoryNotFound) {
		return uc.CategoryService.Update(catName, cat)
	}
	return 0, err
}
