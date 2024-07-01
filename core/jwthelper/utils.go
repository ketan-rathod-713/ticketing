package jwthelper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ketan-rathod-713/ticketing/core/models"
)

func (j *jwtHelper) GenerateToken(emailId string, role string, expireTime time.Duration) (tokenstr string, err error) {

	claims := models.UserClaim{
		EmailId: emailId,
		Role:    role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "ticketing",
			Subject:   "normal data",
			// TODO: what do mean by audience and id in jwt registered claims
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstr, err = token.SignedString(j.Secret)

	return tokenstr, err
}

func (j *jwtHelper) ParseAndValidateToken(tokenstr string) (claims *models.UserClaim, err error) {
	token, err := jwt.ParseWithClaims(tokenstr, &models.UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		return j.Secret, nil
	})

	if err != nil {
		return nil, errors.New("invalid token, unable to parse it")
	} else if  foundClaims, ok := token.Claims.(*models.UserClaim); !ok {
		return nil, errors.New("invalid token with invalid claims")
	} else {
		claims = foundClaims
		return claims, nil
	}
}
