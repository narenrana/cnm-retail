package payments

import (
	"errors"
	repository "shopping-cart/cnm-payments/repository"

)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides booking methods.
type  Service interface {
	add(request paymentsAddRequest) (paymentsAddResponse , error)
	list() (paymentsListResponse, error)
}

type service struct {
}

func (s *service) add(request paymentsAddRequest) (paymentsAddResponse, error) {
	repository:= repository.PaymentsRepositoryInstance()
	p, err:=repository.Add(request.Payments)
	return  paymentsAddResponse{Payments: p}, err
}

func (s *service) list() (paymentsListResponse, error) {
	repository:= repository.PaymentsRepositoryInstance()
	p, err:=repository.List();
	return paymentsListResponse{Payments: p , Err: err}, err
}

// NewService creates a auth service with necessary dependencies.
func NewService() Service {
	return &service{}
}

