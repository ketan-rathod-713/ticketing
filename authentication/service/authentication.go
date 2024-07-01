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
)

func (s *service) Signup(emailId string, password string, firstName string, lastName string) (*models.User, error) {
	log.Println("signup started")

	// check if user already exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := s.DB.Collection("users").FindOne(ctx, bson.M{"emailId": emailId})

	var user models.User
	err := result.Decode(&user)

	if err != nil {
		return nil, err
	}

	if user.EmailId != "" {
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
		return nil, err
	}

	objId, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to get inserted object id")
	}

	// update id of user to inserted object id in mongodb
	user.Id = objId

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
