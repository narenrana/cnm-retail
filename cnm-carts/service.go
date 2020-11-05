package products

import (
	"errors"
	repository "shopping-cart/cnm-carts/repository"

)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides booking methods.
type  Service interface {
	add(request addToCartRequest) (addToCartResponse , error)
	get() (getCartResponse, error)
}

type service struct {
}

func (s *service) add(request addToCartRequest) (addToCartResponse, error) {
	repository:= repository.CartsRepositoryInstance()
	cart, err:=repository.Add(request.Cart)
	return  addToCartResponse{Cart: cart}, err
}

func (s *service) get() (getCartResponse, error) {
	repository:= repository.CartsRepositoryInstance()
	cart, err:=repository.List();
	//TODO- cart empty check missing
	return getCartResponse{Cart: cart[0]}, err
}

// NewService creates a auth service with necessary dependencies.
func NewService() Service {
	return &service{}
}

