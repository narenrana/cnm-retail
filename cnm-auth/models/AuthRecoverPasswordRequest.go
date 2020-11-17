package models


type AuthRecoverPasswordRequest struct {
	UserName       string
	Email          string
	Password       string
}

