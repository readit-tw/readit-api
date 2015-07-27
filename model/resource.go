package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Resource struct {
	Id   bson.ObjectId `bson:"_id" json:"id"`
	Link string        `json:"link"`
}
