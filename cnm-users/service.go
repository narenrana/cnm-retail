package users

import (
	"errors"
	"fmt"
	userRepository "shopping-cart/cnm-users/repository"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides booking methods.
type  Service interface {
	add(request userRequest) (userResponse , error)
	list() (userListResponse, error)
}

type service struct {
}



func (s *service) add(request userRequest) (userResponse, error) {
	repository:= userRepository.UsersRepositoryInstance()
	usr, err:=repository.Add(userRepository.Users{
		FirstName: request.FirstName,
		MiddleName: request.MiddleName,
		LastName: request.LastName,
		UserEmail: request.UserEmail,
		Password: request.Password,
		PhoneNumber: request.PhoneNumber,
			})
	fmt.Print(usr,err);
	return userResponse{UserId: usr.UserId}, nil
}

func (s *service) list() (userListResponse, error) {
	repository:= userRepository.UsersRepositoryInstance()
	users, err:=repository.List();
	return userListResponse{Users:users , Err: err}, err
}

// NewService creates a auth service with necessary dependencies.
func NewService() Service {
	return &service{}
}

