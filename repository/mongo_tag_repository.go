package repository

import (
	"errors"
	"github.com/readit-tw/readit-api/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"fmt"
)

type MongoTagRepository struct {
	db *mgo.Database
}

func NewMongoTagRepository(db *mgo.Database) *MongoTagRepository {
	return &MongoTagRepository{db: db}
}

func (rr *MongoTagRepository) GetByName(name string) (*model.Tag, error) {
	log.Printf("tag name:" + name)
	var tag *model.Tag
	
	err := rr.db.C("tags").Find(bson.M{"name": name}).One(&tag)
	if err != nil {
		return nil, errors.New("Failed to Retrieve")
		log.Printf("Failed to Retrieve search result for :" + name)
	}
	return tag, nil
}

func (rr *MongoTagRepository) GetAll() ([]*model.Tag, error) {
	var tags []*model.Tag
	err := rr.db.C("tags").Find(nil).All(&tags)
	if err != nil {
		return nil, errors.New("Failed to Retrieve")
	}
	return tags, nil

}

func (rr *MongoTagRepository) SearchByTerm(term string) ([]*model.Tag, error) {
	log.Printf("db term:" + term)
	var tags []*model.Tag
	
	//http://stackoverflow.com/questions/3305561/how-to-query-mongodb-with-like
	err := rr.db.C("tags").Find(bson.M{"name": bson.RegEx{term, ""}}).All(&tags)
	if err != nil {
		return nil, errors.New("Failed to Retrieve")
		log.Printf("Failed to Retrieve search result for :" + term)
	}
	return tags, nil
}

func (rr *MongoTagRepository) Create(tag *model.Tag) (*model.Tag, error) {
	
	log.Printf("tag going to creation:" + tag.Name)
	
	var tagToCreate *model.Tag
	err := rr.db.C("tags").Find(bson.M{"name": tag.Name}).One(&tagToCreate)
	log.Printf("is tag found  ?" )
	
	if err ==nil && tagToCreate != nil {
		log.Printf("tag creation : tag found !" )
		return tagToCreate, nil
	}
	if err != nil {
		log.Printf("tag creation : error found :(" )
		fmt.Print(err)
		// nil, errors.New("Failed to Insert")
	}
	
	log.Printf("new tag creation")
	
	tag.Id = bson.NewObjectId()
	err = rr.db.C("tags").Insert(tag)
	if err != nil {
		return nil, errors.New("Failed to Insert")
	}

	log.Printf("tag creation:" + tag.Name)
	
	var createdTag *model.Tag
	err = rr.db.C("tags").Find(bson.M{"_id": tag.Id}).One(&createdTag)
	if err != nil {
		return nil, errors.New("Failed to Insert")
	}

	return createdTag, nil
}