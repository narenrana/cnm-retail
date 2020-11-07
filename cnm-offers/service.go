package offers

import (
	"errors"
	repository "shopping-cart/cnm-offers/repository"

)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides booking methods.
type  Service interface {
	add(request addOffersRequest) (addOffersResponse , error)
	get() (getOffersResponse, error)
}

type service struct {
}

func (s *service) add(request addOffersRequest) (addOffersResponse, error) {
	repository:= repository.CartsRepositoryInstance()
	cart, err:=repository.Add(request.Offers)
	return  addOffersResponse{Offers: cart}, err
}

func (s *service) get() (getOffersResponse, error) {
	repository:= repository.CartsRepositoryInstance()
	offer, err:=repository.List();
	//TODO- cart empty check missing
	return getOffersResponse{Offers: offer}, err
}

// NewService creates a  service with necessary dependencies.
func NewService() Service {
	return &service{}

}

