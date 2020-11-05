package auth

import (
	"errors"
	"fmt"
	"shopping-cart/cnm-auth/repository"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides booking methods.
type Service interface {
	login(user string, password string, rememberMe string) (authLoginResponse, error)
	logout(token string) (authLogoutResponse, error)
}

type service struct {
}

func (s *service) login(user string, password string, rememberMe string) (authLoginResponse, error) {

	response := authLoginResponse{Name: "narender", Token: "6712671267162"}
	repository:=repository.NewUsersRepository()
	usr, err:=repository.FindBy();
	fmt.Print(usr,err);
	return response, nil
}

func (s *service) logout(token string) (authLogoutResponse, error)   {

	response := authLogoutResponse{name: "6712671267162"}

	return response, nil
}

// NewService creates a auth service with necessary dependencies.
func NewService() Service {
	return &service{}
}

