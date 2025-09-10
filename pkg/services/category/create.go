package category

import (
	"bucketWise/pkg/domain"
	"fmt"
	"log"
)

func (s Service) Create(category domain.Category) (id interface{}, err error) {

	insertedId, err := s.Repo.Insert(category)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating category %w", err)
	}

	return insertedId, nil
}
