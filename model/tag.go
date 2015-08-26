package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Tag struct {
	Id    bson.ObjectId `bson:"_id" json:"id"`
	Name string        `json:"name"`
}

func (r *Tag) Validate() []map[string][]string {

	validationErrors := make([]map[string][]string, 0)
	if r.Name == " " || r.Name == "" {
		e := make(map[string][]string)
		e["name"] = []string{"name is required"}
		validationErrors = append(validationErrors, e)
	}

	return validationErrors
}
