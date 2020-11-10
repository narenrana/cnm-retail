package  coupons

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	e "shopping-cart/cnm-coupons/entities"
)

type discountCouponsAddRequest struct {
	DiscountCoupons e.DiscountCoupons ;
}

type discountCouponsResponse struct {
	DiscountCoupons e.DiscountCoupons `json:"discountCoupons,omitempty"`
	Err             error             `json:"error,omitempty"`
}

type discountCouponsListResponse struct {
	Product           []e.DiscountCoupons `json:"products,omitempty"`
	Err              error                `json:"error,omitempty"`
}

type findCouponRequest struct {
	Coupon   string ;
}

type findCouponResponse struct {
	DiscountCoupons e.DiscountCoupons `json:"discountCoupons,omitempty"`
	Valid           bool              `json:"valid"`
	Message         string            `json:"message"`
	Err             error             `json:"error,omitempty"`
}


func (r discountCouponsListResponse) error() error { return r.Err }

func makeDiscountCouponsAddRequestEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(discountCouponsAddRequest)
		response, err := s.add(req)
		return response, err
	}
}

func makeDiscountCouponsListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := s.list()
		return response, err
	}
}

func makeCouponFindEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(findCouponRequest)
		response, err := s.find(req)
		return response, err
	}
}

