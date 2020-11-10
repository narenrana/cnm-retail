package users

import (
	"context"
	"shopping-cart/cnm-users/entities"

	"github.com/go-kit/kit/endpoint"
)

type userRequest struct {
	FirstName      string
	MiddleName     string
	LastName       string
	UserEmail      string
	Password       string
	PhoneNumber    string
}

type userResponse struct {
	UserId            uint `json:"userId,omitempty"`
	FirstName         string `json:"firstName,omitempty"`
	MiddleName        string `json:"middleName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	UserEmail         string `json:"userEmail,omitempty"`
	PhoneNumber       string `json:"phoneNumber,omitempty"`
	Err               error  `json:"error,omitempty"`
}

type userListResponse struct {
	UserDetails            [] entities.UserDetails `json:"users,omitempty"`
	Err              error  `json:"error,omitempty"`
}

func (r userResponse) error() error { return r.Err }

func makeAddUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(userRequest)
		response, err := s.add(req)
		return response, err
	}
}



func makeUserListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := s.list()
		return response, err
	}
}

