package usecases

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
)

type CategoryUseCase struct {
	CategoryService ports.CategoryService
}

func (uc CategoryUseCase) CreateCategoryUseCase(cat domain.Category) (interface{}, error) {
	return uc.CategoryService.Create(cat)
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
