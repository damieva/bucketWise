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

func (s CategoryService) ListAll() ([]domain.Category, error) {

	categoryList, err := s.Repo.SelectAll()
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error listing all the categories %w", err)
	}

	return categoryList, nil
}

func (s CategoryService) ListOne(cat domain.Category) (domain.Category, error) {
	return s.Repo.SelectOne(cat)
}

func (s CategoryService) Delete(cat domain.Category) (int64, error) {
	deletedCount, err := s.Repo.Delete(cat)
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
