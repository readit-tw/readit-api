package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Resource struct {
	Id    bson.ObjectId `bson:"_id" json:"id"`
	Title string        `json:"title"`
	Link  string        `json:"link"`
}

func (r *Resource) Validate() []map[string][]string {

	validationErrors := make([]map[string][]string, 0)
	if r.Title == "" {
		e := make(map[string][]string)
		e["title"] = []string{"title is required"}
		validationErrors = append(validationErrors, e)
	}
	if r.Link == "" {
		e := make(map[string][]string)
		e["link"] = []string{"link is required"}
		validationErrors = append(validationErrors, e)
	}

	return validationErrors

}
