package models

import (
	"shopping-cart/cnm-carts/repository"
	coupoenRepository "shopping-cart/cnm-coupons/repository"
	offersRepository "shopping-cart/cnm-offers/repository"
)

type GetCartResponse struct {
	Cart           repository.Cart `json:"cart,omitempty"`
	Err            error  `json:"error,omitempty"`
	AppliedOffers   []*offersRepository.Offers `json:"appliedOffers,omitempty"`
	IsCouponApplied bool  `json:"isCouponApplied,omitempty"`
	DiscountCoupon *coupoenRepository.DiscountCoupons `json:"discountCoupon,omitempty"`
	TotalDiscount   float64 `json:"totalDiscount,omitempty"`
	TotalAmount     float64 `json:"totalAmount,omitempty"`
}
