package orders

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	e "shopping-cart/cnm-orders/entities"
)

type  addOrdersRequest struct {
	CartId        int         `json:"cartId,omitempty"`
	UserId        *int            `json:"userId,omitempty"`
}

type addOrdersResponse struct {
	Orders            e.Orders   `json:"orders,omitempty"`
	Err               error  `json:"error,omitempty"`
	PaymentUrl         string  `json:"paymentUrl,omitempty"`
	PaymentHost        string  `json:"paymentHost,omitempty"`
	Token            string  `json:"token,omitempty"`
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

