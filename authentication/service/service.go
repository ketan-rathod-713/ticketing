package service

import (
	"github.com/ketan-rathod-713/ticketing/authentication/dto"
	"github.com/ketan-rathod-713/ticketing/core/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service interface {
	Signup(emailId string, password string, firstName string, lastName string) (user *models.User, err error)
	GetUser(emailId string) (user *models.User, err error)
	GetCurrentUser(emailId string) (userInfo *dto.UserInfo, err error)
}

type service struct {
	DB     *mongo.Database
	Logger *zap.SugaredLogger
}

func New(DB *mongo.Database, logger *zap.SugaredLogger) Service {
	return &service{
		DB:     DB,
		Logger: logger,
	}
}
