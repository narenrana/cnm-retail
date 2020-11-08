package orders

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	repository "shopping-cart/cnm-orders/repository"
)

type  addOrdersRequest struct {
	Orders repository.Orders;
}

type addOrdersResponse struct {
	Orders            repository.Orders   `json:"product,omitempty"`
	Err               error  `json:"error,omitempty"`
}

type getOrderListResponse struct {
	Orders           [] *repository.Orders `json:"cart,omitempty"`
	Err            error  `json:"error,omitempty"`
}


func makeOrdersAddEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrdersRequest)
		response, err := s.add(req)
		return response, err
	}
}

func makeGetOrdersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := s.get()
		return response, err
	}
}

