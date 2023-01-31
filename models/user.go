package models

import "gopkg.in/mgo.v2/bson"

// "json:" is the value we return as request response to client,  bson:"_id" is the format its saved as in mongoDB
// sturct type e.g Gender is how we access tha values when workig with mongoDB
type User struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age int `json:"age" bson:"age"`
} 