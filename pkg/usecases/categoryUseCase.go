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
	// First, check if the category already exists
	_, err := uc.CategoryService.ListOne(cat)
	if err == nil {
		// If no error, the category is already in the database
		return nil, domain.ErrCategoryAlreadyExists
	}

	// If the category was not found, proceed with creation
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
