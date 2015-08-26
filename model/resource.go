package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Resource struct {
	Id    bson.ObjectId `bson:"_id" json:"id"`
	Title string        `json:"title"`
	Link  string        `json:"link"`
	Type  string 		`json:"type"`
	Tags  []string		`bson:"-" json:"tags"`// Only for mapping UI/Client tags -> It will always omitted while db commit/persist
	Tag	[]bson.ObjectId `bson:"tag_id,omitempty"`
}

func (r *Resource) Validate() []map[string][]string {

	validationErrors := make([]map[string][]string, 0)
	if r.Title == " " || r.Title == "" {
		e := make(map[string][]string)
		e["title"] = []string{"title is required"}
		validationErrors = append(validationErrors, e)
	}
	if r.Link == " " || r.Link == "" {
		e := make(map[string][]string)
		e["link"] = []string{"link is required"}
		validationErrors = append(validationErrors, e)
	}

	return validationErrors

}
