package usecases

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryUseCase struct {
	CategoryService ports.CategoryService
}

func (uc CategoryUseCase) CreateCategoryUseCase(cat domain.Category) (domain.Category, error) {
	// Comprobar si la categoría ya existe
	_, err := uc.CategoryService.ListOne(cat)
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

		// Convertir el ID de MongoDB a string
		objectID, ok := createdID.(primitive.ObjectID)
		if !ok {
			return domain.Category{}, fmt.Errorf("expected primitive.ObjectID but got %T", createdID)
		}

		// Asignar el ID a la entidad de dominio
		newCat.ID = objectID.Hex()

		// Devolver la categoría con su nuevo ID
		return newCat, nil
	}

	// Si el error no es ErrCategoryNotFound, devolverlo tal cual
	return domain.Category{}, err
}

func (uc CategoryUseCase) ListAllCategoryUseCase() ([]domain.Category, error) {
	return uc.CategoryService.ListAll()
}

func (uc CategoryUseCase) ListOneCategoryUseCase(cat domain.Category) (domain.Category, error) {
	return uc.CategoryService.ListOne(cat)
}

func (uc CategoryUseCase) DeleteCategoryUseCase(cat domain.Category) (int64, error) {
	return uc.CategoryService.Delete(cat)
}

func (uc CategoryUseCase) UpdateCategoryUseCase(catName string, cat domain.Category) (int64, error) {
	// Verify that the category to update actually exists in the database
	_, err := uc.CategoryService.ListOne(domain.Category{Name: catName})
	if errors.Is(err, domain.ErrCategoryNotFound) {
		return 0, domain.ErrCategoryNotFound
	} else if err != nil {
		return 0, err
	}

	// Ensure the new category data does not conflict with an existing record
	_, err = uc.CategoryService.ListOne(cat)
	if err == nil {
		return 0, domain.ErrCategoryAlreadyExists
	} else if errors.Is(err, domain.ErrCategoryNotFound) {
		return uc.CategoryService.Update(catName, cat)
	}
	return 0, err
}
