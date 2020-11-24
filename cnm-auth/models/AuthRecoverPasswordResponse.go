package models

type AuthRecoverPasswordResponse struct {
	UserName       string
	Email          string
	Password       string
}