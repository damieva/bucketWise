package services

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
	"fmt"
	"log"
)

type CategoryService struct {
	Repo ports.CategoryRepository
}

func (s CategoryService) Create(cat domain.Category) (interface{}, error) {

	insertedId, err := s.Repo.Insert(cat)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating category %w", err)
	}

	return insertedId, nil
}

func (s CategoryService) List(name string) ([]domain.Category, error) {

	categoryList, err := s.Repo.Select(name)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error listing all the categories %w", err)
	}

	return categoryList, nil
}

func (s CategoryService) Delete(IDs []string) (int64, error) {
	deletedCount, err := s.Repo.Delete(IDs)
	if err != nil {
		log.Println(err.Error())
		return deletedCount, fmt.Errorf("error deleting category %w", err)
	}

	return deletedCount, nil
}

func (s CategoryService) Update(catName string, cat domain.Category) (int64, error) {
	modifiedCount, err := s.Repo.Update(catName, cat)
	if err != nil {
		log.Println(err.Error())
		return modifiedCount, fmt.Errorf("error modifying category %w", err)
	}

	return modifiedCount, nil
}
