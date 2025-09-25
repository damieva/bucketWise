package usecases

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
	"errors"
)

type CategoryUseCase struct {
	CategoryService ports.CategoryService
}

func (uc CategoryUseCase) CreateCategoryUseCase(cat domain.Category) (interface{}, error) {
	// Primero verificamos si la categoría ya existe
	_, err := uc.CategoryService.ListOne(cat)

	if err == nil {
		// Si ListOne devuelve nil, la categoría ya existe
		return nil, domain.ErrCategoryAlreadyExists
	}

	// Si la categoria no existe, la creamos
	if errors.Is(err, domain.ErrCategoryNotFound) {
		return uc.CategoryService.Create(cat)
	}

	return nil, err

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
