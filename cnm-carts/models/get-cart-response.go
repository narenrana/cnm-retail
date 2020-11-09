package models

import (
	"shopping-cart/cnm-carts/repository"
	coupoenRepository "shopping-cart/cnm-coupons/repository"
	offersRepository "shopping-cart/cnm-offers/repository"
)

type CartResponse struct {
	CartId        		*int      	`gorm:"primaryKey" json:"cartId,omitempty"`
	CartName      		string    	`json:"cartName,omitempty"`
	UserId        		int       	`json:"userId,omitempty"`
	CartItems     		 []*repository.CartItems `json:"cartItems,omitempty"`
	Err           		 error  	`json:"error,omitempty"`
	AppliedOffers   	[]*offersRepository.Offers `json:"appliedOffers,omitempty"`
	OffersDiscount     float64  		`json:"offersDiscount,omitempty"`
	CouponDiscount 	     float64  		`json:"couponDiscount,omitempty"`
	DiscountCoupon      *string       `json:"discountCoupon,omitempty"`
	DiscountCouponDetails 		*coupoenRepository.DiscountCoupons `json:"discountCouponDetails,omitempty"`
	TotalDiscount   	float64 	`json:"totalDiscount,omitempty"`
	TotalAmount     	float64 	`json:"totalAmount,omitempty"`
}
