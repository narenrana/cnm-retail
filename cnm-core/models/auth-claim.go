package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	UserId  *int `json:"userId"`
	jwt.StandardClaims
}
