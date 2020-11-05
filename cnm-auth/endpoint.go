package auth

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type authLoginRequest struct {
	name       string
	password   string
	rememberMe string
}

type authLoginResponse struct {
	Name       string `json:"name,omitempty"`
	Password   string `json:"password,omitempty"`
	RememberMe string `json:"rememberMe,omitempty"`
	Token      string `json:"token,omitempty"`
	Date       string `json:"date,omitempty"`
	Err        error  `json:"error,omitempty"`
}

func (r authLoginResponse) error() error { return r.Err }

func makeAuthLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(authLoginRequest)
		response, err := s.login(req.name, req.password, req.rememberMe)
		return authLoginResponse{Name: response.Name, Token: response.Token, Date: response.Date, Err: err}, nil
	}
}

type authLogoutRequest struct {
	token      string `json:"token,omitempty"`
}

type authLogoutResponse struct {
	name       string `json:"name"`
	err        error  `json:"error"`
}

func makeAuthLogoutEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(authLogoutRequest)
		response, err := s.logout(req.token)
		return authLogoutResponse{ name: response.name,err: err}, nil
	}
}

