package coupons

import (
	"errors"
	repository "shopping-cart/cnm-coupons/repository"

)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides booking methods.
type  Service interface {
	add(request discountCouponsAddRequest) (discountCouponsResponse , error)
	list() (discountCouponsListResponse, error)
}

type service struct {
}

func (s *service) add(request discountCouponsAddRequest) (discountCouponsResponse, error) {
	repository:= repository.DiscountCouponsRepositoryInstance()
	product, err:=repository.Add(request.DiscountCoupons)
	return  discountCouponsResponse{DiscountCoupons: product}, err
}

func (s *service) list() (discountCouponsListResponse, error) {
	repository:= repository.DiscountCouponsRepositoryInstance()
	product, err:=repository.List();
	return discountCouponsListResponse{Product: product , Err: err}, err
}

// NewService creates a auth service with necessary dependencies.
func NewService() Service {
	return &service{}
}

