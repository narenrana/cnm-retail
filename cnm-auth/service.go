package auth

import (
	"errors"
	repositoryDao "shopping-cart/cnm-auth/repository"
	"shopping-cart/cnm-core/utils"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("Invalid argument")

var ErrInvalidPassword = errors.New("Invalid Password")

var ErrInvalidSignature = errors.New("Invalid Token")

// Service is the interface that provides booking methods.
type Service interface {
	login(request authLoginRequest) (authLoginResponse, error)
	logout(token string) (authLogoutResponse, error)
	recoverPassword(request authRecoverPasswordRequest) (authRecoverPasswordResponse, error)
	refreshToken(request authRefreshTokenRequest)(authRefreshTokenResponse, error)
    signUp(request authSignUpRequest ) (authSignUpResponse, error)
}

type service struct {
}



func (s *service) login(request authLoginRequest) (authLoginResponse, error){

	repository:= repositoryDao.NewUsersRepository()
	usr, err:=repository.FindById(request.Email);
	if err != nil {
		return authLoginResponse{Err: err}, err
	}

	isValid:=utils.CheckPasswordHash(request.Password,usr.Password)

	if !isValid {
		return authLoginResponse{}, ErrInvalidPassword
	}

	token, error:=utils.CreateToken(usr.UserEmail);

	if error != nil {
		return authLoginResponse{Err: err}, err
	}

	return authLoginResponse{
		Token: token,
		Email: usr.UserEmail,
	}, err
}

func (s *service) logout(token string) (authLogoutResponse, error)   {
	response := authLogoutResponse{Email: "6712671267162"}
	return response, nil
}

func (s *service) recoverPassword(request authRecoverPasswordRequest) (authRecoverPasswordResponse, error){
	response := authRecoverPasswordResponse{ }
	/**
	  Write login to send email to reset password
	 */
	return response, nil
}



func (s *service) refreshToken(request authRefreshTokenRequest)(authRefreshTokenResponse, error){

	isValid,token,err:=utils.RefreshToken(request.Token);

	if err !=nil { return authRefreshTokenResponse{}, err  }

	if(!isValid) { return authRefreshTokenResponse{}, ErrInvalidSignature }

	return authRefreshTokenResponse{
		Token: token,
	}, nil
}

func (s *service)  signUp(request authSignUpRequest ) (authSignUpResponse, error){
	repository:=repositoryDao.NewUsersRepository()
	var users repositoryDao.Users

	users.FirstName=request.FirstName
	users.LastName=request.LastName
	users.MiddleName=request.MiddleName
	users.UserEmail=request.Email
	users.PhoneNumber=request.PhoneNumber
	users.UserEmail=request.Email
	hashKey, err :=utils.HashPassword(request.Password)
	users.Password=hashKey;
	if err != nil {
		return authSignUpResponse{},err
	}
	 usr, err:=repository.AddOrUpdate(users);

	if err!=nil   {
		return authSignUpResponse{}, err
	}

	token, error:=utils.CreateToken(usr.UserEmail);

	if err!=nil   {
		return authSignUpResponse{}, error
	}

	return authSignUpResponse{
		Email: usr.UserEmail,
		Token: token,
	}, err
}


// NewService creates a auth service with necessary dependencies.
func NewService() Service {
	return &service{}
}

