package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// user struct for storing user info
type User struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id"`
	EmailId         string             `json:"emailId" bson:"emailId"`
	Password        string             `json:"password" bson:"password"`
	Role            string             `json:"role" bson:"role"`
	IsEmailVerified bool               `json:"isEmailVerified" bson:"isEmailVerified"`
	CreatedAt       string             `json:"createdAt" bson:"createdAt"`
	UpdatedAt       string             `json:"updatedAt" bson:"updatedAt"`
}
