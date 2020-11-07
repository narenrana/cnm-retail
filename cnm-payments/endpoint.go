package  payments

import (
	"context"
	repository "shopping-cart/cnm-payments/repository"
	"github.com/go-kit/kit/endpoint"
)

type paymentsAddRequest struct {
	Payments   repository.Payments ;
}

type paymentsAddResponse struct {
	Payments          repository.Payments   `json:"payments,omitempty"`
	Err               error  `json:"error,omitempty"`
}

type paymentsListResponse struct {
	Payments           []repository.Payments `json:"products,omitempty"`
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

