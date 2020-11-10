package carts

import (
	"context"
	//"shopping-cart/cnm-core/repository"
	"shopping-cart/cnm-products/entities"

	"github.com/go-kit/kit/endpoint"
)

type productAddRequest struct {
	Product entities.Product;
}

type productAddResponse struct {
	Product          entities.Product   `json:"product,omitempty"`
	Err               error  `json:"error,omitempty"`
}

type productListResponse struct {
	Product           []entities.Product `json:"products,omitempty"`
	Err              error  `json:"error,omitempty"`
}

func (r productListResponse) error() error { return r.Err }

func makeProductAddRequestEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(productAddRequest)
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

