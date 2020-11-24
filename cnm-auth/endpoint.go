package auth

import (
	"context"
	"shopping-cart/cnm-auth/models"
	"shopping-cart/cnm-auth/services"

	"github.com/go-kit/kit/endpoint"
)


func makeAuthLoginEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.AuthLoginRequest)
		response, err := s.Login(req)
		return response, err
	}
}

func makeAuthLogoutEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.AuthLogoutRequest)
		response, err := s.Logout(req.Token)
		return response, err
	}
}

func makeRefreshTokenEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.AuthRefreshTokenRequest)
		response, err := s.RefreshToken(req)
		return response, err
	}
}

func makeSignUpEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.AuthSignUpRequest)
		response, err := s.SignUp(req)
		return response, err
	}
}

func makeRecoverPasswordEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.AuthSignUpRequest)
		response, err := s.SignUp(req)
		return response, err
	}
}

