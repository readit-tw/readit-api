package repository

import (
	"github.com/readit-tw/readit-api/model"
)

type ResourceRepository interface {
	Create(resource *model.Resource) (*model.Resource, error)
	GetAll() ([]*model.Resource, error)
	SearchByTerm(term string) ([]*model.Resource, error)
}
