package repository

import (
	"github.com/readit-tw/readit-api/model"
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestResourceRepositoryCreate(t *testing.T) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Session Failed %v", err)
		return
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	db := session.DB("readit_test")
	resourceRepo := &ResourceRepository{db}
	res := &model.Resource{Link: "http://www.google.com"}

	createdRes, err := resourceRepo.Create(res)
	if err != nil {
		t.Errorf("Failed %v", err)
		return
	}
	expected, actual := res.Link, createdRes.Link
	if expected != actual {
		t.Errorf("Failed asserting %s to be %s", actual, expected)
	}
}

func TestMain(m *testing.M) {
	testResult := m.Run()

	session, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Session Failed %v", err)
		return
	}
	session.DB("readit_test").DropDatabase()
	session.Close()

	os.Exit(testResult)
}
