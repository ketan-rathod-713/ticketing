package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"objectId"`
	EmailId   string             `json:"emailId"`
	Password  string             `json:"password"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
}
