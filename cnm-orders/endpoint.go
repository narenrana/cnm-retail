package orders

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	e "shopping-cart/cnm-orders/entities"
)

type  addOrdersRequest struct {
	Orders e.Orders;
}

type addOrdersResponse struct {
	Orders            e.Orders   `json:"product,omitempty"`
	Err               error  `json:"error,omitempty"`
}

type getOrderListResponse struct {
	Orders           [] *e.Orders `json:"cart,omitempty"`
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

