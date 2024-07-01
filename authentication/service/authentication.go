package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ketan-rathod-713/ticketing/authentication/dto"
	"github.com/ketan-rathod-713/ticketing/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) Signup(emailId string, password string, firstName string, lastName string) (*models.User, error) {
	log.Println("signup started")

	// check if user already exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := s.DB.Collection("users").FindOne(ctx, bson.M{"emailId": emailId})

	var user models.User
	err := result.Decode(&user)

	// if there are no documents in collection then move forward and do signup
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		s.Logger.Infof("error decoding model user got from database %v", err.Error())
		return nil, err
	}

	if user.EmailId != "" {
		s.Logger.Info("emailId already exists")
		return nil, errors.New("emailId already exists")
	}

	//TODO: save user to database and return it

	var signupReq dto.SignupReq = dto.SignupReq{
		EmailId:   emailId,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}

	insertResult, err := s.DB.Collection("users").InsertOne(ctx, signupReq)
	if err != nil {
		s.Logger.Infof("error inserting signupReq into database %v", err.Error())
		return nil, err
	}

	objId, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		s.Logger.Info("failed to get inserted object id")
		return nil, errors.New("failed to get inserted object id")
	}

	// update id of user to inserted object id in mongodb
	user.Id = objId
	user.EmailId = emailId
	user.FirstName = firstName
	user.LastName = lastName

	return &user, nil
}

func (s *service) GetUser(emailId string) (user *models.User, err error) {

	// TODO: get user from database

	user = &models.User{
		EmailId:   emailId,
		FirstName: "ketan",
		LastName:  "rathod",
		Password:  "1234",
	}

	return user, nil
}
