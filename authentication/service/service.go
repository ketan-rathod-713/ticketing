package service

import (
	"github.com/ketan-rathod-713/ticketing/core/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	Signup(emailId string, password string) (user *models.User, err error)
}

type service struct {
	DB *mongo.Database
}

func New(DB *mongo.Database) Service {
	return &service{
		DB: DB,
	}
}
