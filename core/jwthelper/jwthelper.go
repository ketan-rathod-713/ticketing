package jwthelper

import (
	"time"

	"github.com/ketan-rathod-713/ticketing/core/models"
)

type JWTHelper interface {
	GenerateToken(emailId string, role string, expireTime time.Duration) (tokenstr string, err error)
	ParseAndValidateToken(tokenstr string) (claims *models.UserClaim, err error)
}

type jwtHelper struct {
	Secret []byte
}

func New(secret string) JWTHelper {
	return &jwtHelper{
		Secret: []byte(secret),
	}
}
