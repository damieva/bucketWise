package ports

import "bucketWise/pkg/domain"

type CategoryService interface {
	Create(category domain.Category) (id interface{}, err error)
}

type CategoryRepository interface {
	Insert(category domain.Category) (id interface{}, err error)
}
