package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"objectId" bson:"id"`
	EmailId   string             `json:"emailId" bson:"emailId"`
	Password  string             `json:"password" bson:"password"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
}
