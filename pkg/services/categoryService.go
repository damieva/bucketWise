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

func (s CategoryService) Create(cat domain.Category) (id interface{}, err error) {

	insertedId, err := s.Repo.Insert(cat)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating category %w", err)
	}

	return insertedId, nil
}

func (s CategoryService) ListAll() ([]domain.Category, error) {

	categoryCollection, err := s.Repo.SelectAll()
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error listing all the categories %w", err)
	}

	return categoryCollection, nil
}
