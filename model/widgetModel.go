package model

import (
	"gopkg.in/mgo.v2/bson"
)

//Widget application model
type Widget struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Color     string        `bson:"color" json:"color"`
	Price     string        `bson:"price" json:"price"`
	Inventory uint64        `bson:"inventory" json:"inventory"`
	Melts     bool          `bson:"melts" json:"melts"`
}
