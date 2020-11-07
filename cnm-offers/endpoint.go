package offers

import (
	"context"
	"shopping-cart/cnm-offers/repository"

	"github.com/go-kit/kit/endpoint"
)

type  addOffersRequest struct {
	Offers repository.Offers;
}

type addOffersResponse struct {
	Offers          repository.Offers   `json:"offers,omitempty"`
	Err               error  `json:"error,omitempty"`
}

type getOffersResponse struct {
	Offers         [] *repository.Offers `json:"offers,omitempty"`
	Err            error  `json:"error,omitempty"`
}

func makeOffersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOffersRequest)
		response, err := s.add(req)
		return response, err
	}
}

func makeGetOffersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := s.get()
		return response, err
	}
}

