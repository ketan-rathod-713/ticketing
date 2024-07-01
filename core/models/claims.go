package models

import "github.com/golang-jwt/jwt/v5"

type UserClaim struct {
	EmailId string
	Role    string
	jwt.RegisteredClaims
}
