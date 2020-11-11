package models

import "github.com/dgrijalva/jwt-go"

type PaymentClaims struct {
	OrderId *int `json:"username"`
	UserId  *int `json:"userId"`
	Amount float64
	jwt.StandardClaims
}

