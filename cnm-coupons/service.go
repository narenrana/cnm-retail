package coupons

import (
	"errors"
	cr "shopping-cart/cnm-coupons/repository"
	"time"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides booking methods.
type  Service interface {
	add(request discountCouponsAddRequest) (discountCouponsResponse , error)
	list() (discountCouponsListResponse, error)
    find(request findCouponRequest) (findCouponResponse,error)
}

type service struct {
}


func (s *service) add(request discountCouponsAddRequest) (discountCouponsResponse, error) {
	instance:= cr.RepositoryInstance()
	product, err:=instance.Add(request.DiscountCoupons)
	return  discountCouponsResponse{DiscountCoupons: product}, err
}

func (s *service) list() (discountCouponsListResponse, error) {
	instance:= cr.RepositoryInstance()
	product, err:=instance.List();
	return discountCouponsListResponse{Product: product , Err: err}, err
}

func (s *service) find(request findCouponRequest) (findCouponResponse,error){
	instance:= cr.RepositoryInstance()
	coupon, err:=instance.FindByDiscountCoupon(request.Coupon);
	isValid:= coupon.Active &&  coupon.ExpiryDate.After(time.Now())
	return findCouponResponse{ DiscountCoupons: coupon, Valid:isValid  },err

}

// NewService creates a auth service with necessary dependencies.
func NewService() Service {
	return &service{}
}

