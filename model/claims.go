package model

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	UserId string
	Role   string
	Email  string `json:"email"`
	jwt.StandardClaims
}
