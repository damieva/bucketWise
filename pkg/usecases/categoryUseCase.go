package usecases

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
)

type CategoryUseCase struct {
	CategoryService ports.CategoryService
}

func (uc CategoryUseCase) CreateCategoryUseCase(cat domain.Category) (id interface{}, err error) {
	// Aquí no hay lógica adicional: solo orquestamos
	return uc.CategoryService.Create(cat)
}

func (uc CategoryUseCase) ListAllCategoryUseCase() (id interface{}, err error) {
	return uc.CategoryService.ListAll()
}
