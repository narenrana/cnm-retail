package services

import (
	"errors"
	"shopping-cart/cnm-auth/models"

	"shopping-cart/cnm-auth/entities"
	"shopping-cart/cnm-auth/repository"
	"shopping-cart/cnm-core/utils"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("Invalid argument")

var ErrInvalidPassword = errors.New("Invalid Password")

var ErrInvalidSignature = errors.New("Invalid Token")

// Service is the interface that provides booking methods.
type Service interface {
	Login(request models.AuthLoginRequest) (models.AuthLoginResponse, error)
	Logout(token string) (models.AuthLogoutResponse, error)
	RecoverPassword(request models.AuthRecoverPasswordRequest) (models.AuthRecoverPasswordResponse, error)
	RefreshToken(request models.AuthRefreshTokenRequest)(models.AuthRefreshTokenResponse, error)
    SignUp(request models.AuthSignUpRequest) (models.AuthSignUpResponse, error)
}

type service struct {
}



func (s *service) Login(request models.AuthLoginRequest) (models.AuthLoginResponse, error){

	instance:= repository.NewUsersRepository()
	usr, err:=instance.FindById(request.Email);
	if err != nil {
		return models.AuthLoginResponse{Err: err}, err
	}

	isValid:=utils.CheckPasswordHash(request.Password,usr.Password)

	if !isValid {
		return models.AuthLoginResponse{}, ErrInvalidPassword
	}

	token, error:=utils.CreateToken(usr.UserEmail, usr.UserId);

	if error != nil {
		return models.AuthLoginResponse{Err: err}, err
	}

	return models.AuthLoginResponse{
		Token: token,
		Email: usr.UserEmail,
	}, err
}

func (s *service) Logout(token string) (models.AuthLogoutResponse, error)   {
	response := models.AuthLogoutResponse{Email: "6712671267162"}
	return response, nil
}

func (s *service) RecoverPassword(request models.AuthRecoverPasswordRequest) (models.AuthRecoverPasswordResponse, error){
	response := models.AuthRecoverPasswordResponse{ }
	/**
	  Write login to send email to reset password
	 */
	return response, nil
}



func (s *service) RefreshToken(request models.AuthRefreshTokenRequest)(models.AuthRefreshTokenResponse, error){

	isValid,token,err:=utils.RefreshToken(request.Token);

	if err !=nil { return models.AuthRefreshTokenResponse{}, err  }

	if(!isValid) { return models.AuthRefreshTokenResponse{}, ErrInvalidSignature
	}

	return models.AuthRefreshTokenResponse{
		Token: token,
	}, nil
}

func (s *service)  SignUp(request models.AuthSignUpRequest) (models.AuthSignUpResponse, error){
	instance:=repository.NewUsersRepository()
	var users entities.Users

	users.FirstName=request.FirstName
	users.LastName=request.LastName
	users.MiddleName=request.MiddleName
	users.UserEmail=request.Email
	users.PhoneNumber=request.PhoneNumber
	users.UserEmail=request.Email
	hashKey, err :=utils.HashPassword(request.Password)
	users.Password=hashKey;
	if err != nil {
		return models.AuthSignUpResponse{},err
	}
	 usr, err:=instance.AddOrUpdate(users);

	if err!=nil   {
		return models.AuthSignUpResponse{}, err
	}

	token, error:=utils.CreateToken(usr.UserEmail,usr.UserId);

	if err!=nil   {
		return models.AuthSignUpResponse{}, error
	}

	return models.AuthSignUpResponse{
		Email: usr.UserEmail,
		Token: token,
	}, err
}


// NewService creates a auth service with necessary dependencies.
func NewService() Service {
	return &service{}
}

