package offers

import (
	"context"
	e "shopping-cart/cnm-offers/entities"

	"github.com/go-kit/kit/endpoint"
)

type  addOffersRequest struct {
	Offers e.Offers;
}

type addOffersResponse struct {
	Offers          e.Offers   `json:"offers,omitempty"`
	Err               error  `json:"error,omitempty"`
}

type getOffersResponse struct {
	Offers         [] *e.Offers `json:"offers,omitempty"`
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
		response, err := s.list()
		return response, err
	}
}

