package repository

import (
	"github.com/readit-tw/readit-api/model"
)

type TagRepository interface {
	Create(tag *model.Tag) (*model.Tag, error)
	GetAll() ([]*model.Tag, error)
	SearchByTerm(term string) ([]*model.Tag, error)
	GetByName(name string) (*model.Tag, error)
}

