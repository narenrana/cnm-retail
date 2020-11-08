package auth

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type authLoginRequest struct {
	Email        string `json:"email,omitempty"`
	Password     string   `json:"password,omitempty"`
	RememberMe   bool   `json:"rememberMe,omitempty"`
}

type authLoginResponse struct {
	Email        string `json:"email,omitempty"`
	Token      string `json:"token,omitempty"`
	Err        error  `json:"error,omitempty"`
}

type authSignUpRequest struct {
	FirstName      string   `json:"firstName,omitempty"`
	MiddleName     string   `json:"middleName,omitempty"`
	LastName       string     `json:"lastName,omitempty"`
	Email          string     `json:"email,omitempty"`
	Password       string      `json:"password,omitempty"`
	PhoneNumber    string   `json:"phoneNumber,omitempty"`
}

type authSignUpResponse struct {
	Email          string `json:"email,omitempty"`
	Token          string  `json:"token,omitempty"`
	Err        error  `json:"error,omitempty"`

}

type authRecoverPasswordRequest struct {
	UserName       string
	Email          string
	Password       string
}

type authRecoverPasswordResponse struct {
	UserName       string
	Email          string
	Password       string
}


type authLogoutRequest struct {
	Token      string `json:"token,omitempty"`
}

type authLogoutResponse struct {
	Email       string `json:"name"`
	err        error  `json:"error"`
}

type authRefreshTokenRequest struct {
	Token      string `json:"token,omitempty"`
}

type authRefreshTokenResponse struct {
	Token      string `json:"token,omitempty"`
}

func (r authLoginResponse) error() error { return r.Err }

func makeAuthLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(authLoginRequest)
		response, err := s.login(req)
		return response, err
	}
}

func makeAuthLogoutEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(authLogoutRequest)
		response, err := s.logout(req.Token)
		return response, err
	}
}

func makeRefreshTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(authRefreshTokenRequest)
		response, err := s.refreshToken(req)
		return response, err
	}
}

func makeSignUpEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(authSignUpRequest)
		response, err := s.signUp(req)
		return response, err
	}
}

func makeRecoverPasswordEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(authSignUpRequest)
		response, err := s.signUp(req)
		return response, err
	}
}

