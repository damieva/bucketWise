package usecases

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/services"
	"context"
)

type CreateCategoryUseCase struct {
	categoryService services.CategoryService
}

func (uc CreateCategoryUseCase) Run(ctx context.Context, cat domain.Category) (id interface{}, err error) {
	// Aquí no hay lógica adicional: solo orquestamos
	return uc.categoryService.Create(cat)
}
