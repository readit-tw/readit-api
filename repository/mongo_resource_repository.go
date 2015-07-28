package repository

import (
	"errors"
	"github.com/readit-tw/readit-api/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoResourceRepository struct {
	db *mgo.Database
}

func NewMongoResourceRepository(db *mgo.Database) *MongoResourceRepository {
	return &MongoResourceRepository{db: db}
}

func (rr *MongoResourceRepository) Create(resource *model.Resource) (*model.Resource, error) {
	resource.Id = bson.NewObjectId()
	err := rr.db.C("resources").Insert(resource)
	if err != nil {
		return nil, errors.New("Failed to Insert ")
	}

	var createdResource *model.Resource
	err = rr.db.C("resources").Find(bson.M{"_id": resource.Id}).One(&createdResource)
	if err != nil {
		return nil, errors.New("Failed to Insert")
	}

	return createdResource, nil
}