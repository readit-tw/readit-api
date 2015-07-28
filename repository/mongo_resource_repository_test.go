package repository

import (
	"github.com/readit-tw/readit-api/model"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"os"
	"testing"
)

var resourceRepo *MongoResourceRepository

func TestMongoResourceRepositoryGetAll(t *testing.T) {
	res := &model.Resource{Link: "http://www.google.com"}
	res1 := &model.Resource{Link: "http://www.yahoo.com"}
	_, err := resourceRepo.Create(res)
	assert.Nil(t, err)
	_, err = resourceRepo.Create(res1)
	assert.Nil(t, err)

	resources, err := resourceRepo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(resources))
}

func TestMongoResourceRepositoryCreate(t *testing.T) {

	res := &model.Resource{Link: "http://www.google.com"}

	createdRes, err := resourceRepo.Create(res)
	assert.Nil(t, err)
	assert.Equal(t, res.Link, createdRes.Link)
}

func TestMain(m *testing.M) {
	session, err := mgo.Dial("localhost")
	db := session.DB("readit_test")
	session.SetMode(mgo.Monotonic, true)
	resourceRepo = NewMongoResourceRepository(db)
	if err != nil {
		panic(err)
	}
	testResult := m.Run()

	db.DropDatabase()
	session.Close()

	os.Exit(testResult)
}
