package products

import (
	"errors"
	repository "shopping-cart/cnm-products/repository"

)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides booking methods.
type  Service interface {
	add(request productAddRequest) (productAddResponse , error)
	list() (productListResponse, error)
}

type service struct {
}

func (s *service) add(request productAddRequest) (productAddResponse, error) {
	repository:= repository.ProductRepositoryInstance()
	product, err:=repository.Add(request.Product)
	return  productAddResponse{Product: product}, err
}

func (s *service) list() (productListResponse, error) {
	repository:= repository.ProductRepositoryInstance()
	product, err:=repository.List();
	return productListResponse{Product: product , Err: err}, err
}

// NewService creates a auth service with necessary dependencies.
func NewService() Service {
	return &service{}
}

