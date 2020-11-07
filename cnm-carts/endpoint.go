package products

import (
	"context"
	"shopping-cart/cnm-carts/repository"

	"github.com/go-kit/kit/endpoint"
)

type  addToCartRequest struct {
	Cart repository.Cart;
}

type addToCartResponse struct {
	Cart          repository.Cart   `json:"product,omitempty"`
	Err               error  `json:"error,omitempty"`
}

type getCartResponse struct {
	Cart           [] *repository.Cart `json:"cart,omitempty"`
	Err            error  `json:"error,omitempty"`
}


func makeAddToCartEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addToCartRequest)
		response, err := s.add(req)
		return response, err
	}
}

func makeGetCartEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := s.get()
		return response, err
	}
}

