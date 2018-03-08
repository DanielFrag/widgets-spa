package model

import (
	"gopkg.in/mgo.v2/bson"
)

//User application model
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Login    string        `bson:"login" json:"name"`
	Password string        `bson:"password"`
	Session  string        `bson:"session"`
	Gravatar string        `bson:"gravatar" json:"gravatar"`
}
