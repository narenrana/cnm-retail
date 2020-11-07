package orders

import (
	"errors"
	repository "shopping-cart/cnm-orders/repository"

)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides booking methods.
type  Service interface {
	add(request addOrdersRequest) (addOrdersResponse, error)
	get() (getOrderListResponse, error)
}

type service struct {
}

func (s *service) add(request addOrdersRequest) (addOrdersResponse, error) {
	repository:= repository.CartsRepositoryInstance()
	item, err:=repository.Add(request.Orders)
	return  addOrdersResponse{Orders: item}, err
}

func (s *service) get() (getOrderListResponse, error) {
	repository:= repository.CartsRepositoryInstance()
	cart, err:=repository.List();
	//TODO- cart empty check missing
	return getOrderListResponse{Orders: cart}, err
}

// NewService creates a  service with necessary dependencies.
func NewService() Service {
	return &service{}
}

