package carts

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"shopping-cart/cnm-carts/models"

	"shopping-cart/cnm-carts/services"

)



func makeAddToCartEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.AddToCartRequest)
		response, err := s.Add(req)
		return response, err
	}
}

func makeGetCartEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.GetCartRequest)
		response, err := s.Get(req)
		return response, err
	}
}

