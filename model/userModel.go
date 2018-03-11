package model

import (
	"gopkg.in/mgo.v2/bson"
)

//User application model
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Login    string        `bson:"login" json:"name"`
	Password string        `bson:"password" json:"-"`
	Session  string        `bson:"session" json:"-"`
	Gravatar string        `bson:"gravatar" json:"gravatar"`
}
