package  payments

import (
	"context"
	e "shopping-cart/cnm-payments/entities"
	"github.com/go-kit/kit/endpoint"
)

type paymentsAddRequest struct {
	Payments   e.Payments ;
}

type paymentsAddResponse struct {
	Payments          e.Payments   `json:"payments,omitempty"`
	Err               error  `json:"error,omitempty"`
}

type paymentsListResponse struct {
	Payments           []e.Payments `json:"products,omitempty"`
	Err              error  `json:"error,omitempty"`
}

func (r paymentsListResponse) error() error { return r.Err }

func makePaymentsAddRequestEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(paymentsAddRequest)
		response, err := s.add(req)
		return response, err
	}
}

func makePaymentsListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := s.list()
		return response, err
	}
}

