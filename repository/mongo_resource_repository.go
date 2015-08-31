package repository

import (
	"errors"
	"github.com/readit-tw/readit-api/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	//"fmt"
)

type MongoResourceRepository struct {
	db *mgo.Database
}

func NewMongoResourceRepository(db *mgo.Database) *MongoResourceRepository {
	return &MongoResourceRepository{db: db}
}
func (rr *MongoResourceRepository) SearchByTerm(term string) ([]*model.Resource, error) {
	var resources []*model.Resource
	
	
	// Get Tag ids matching query term
	var tags []*model.Tag
	err := rr.db.C("tags").Find(bson.M{"name": bson.RegEx{term, "i"}}).Select(bson.M{"_id": 1}).All(&tags)
	
	tagIds := make([]bson.ObjectId, 0, len(tags))
	for index := range tags{
		tagIds = append(tagIds, tags[index].Id)
	}
	
	//fmt.Println("tag objectIds is ", tagIds)
	
	//http://stackoverflow.com/questions/3305561/how-to-query-mongodb-with-like
	//err = rr.db.C("resources").Find(bson.M{"title": bson.RegEx{term, "i"}}).All(&resources)
	err = rr.db.C("resources").Find(bson.M{ "$or": 
												[]bson.M{
														bson.M{
															"title": 
															bson.RegEx{term, "i"},
															}, 
														bson.M{
															"tag_id": 
															bson.M{"$in" : tagIds},
															},
														},
												}).All(&resources)
	
	//log.Printf("after serach query ")
	
	if err != nil {
		return nil, errors.New("Failed to Retrieve")
		log.Printf("Failed to Retrieve search result for :" + term)
	}
	return resources, nil
}

func (rr *MongoResourceRepository) GetAll() ([]*model.Resource, error) {
	var resources []*model.Resource
	err := rr.db.C("resources").Find(nil).All(&resources)
	if err != nil {
		return nil, errors.New("Failed to Retrieve")
	}
	return resources, nil

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
