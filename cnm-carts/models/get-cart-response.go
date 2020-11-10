package models

import (
	"shopping-cart/cnm-carts/entities"
	e "shopping-cart/cnm-coupons/entities"
	ce "shopping-cart/cnm-offers/entities"
)

type CartResponse struct {
	CartId        		*int      	`gorm:"primaryKey" json:"cartId,omitempty"`
	CartName      		string    	`json:"cartName,omitempty"`
	UserId        		*int       	`json:"userId,omitempty"`
	CartItems     		 []*entities.CartItems `json:"cartItems,omitempty"`
	Err           		 error  	`json:"error,omitempty"`
	AppliedOffers   	[]*ce.Offers `json:"appliedOffers,omitempty"`
	OffersDiscount     float64  		`json:"offersDiscount,omitempty"`
	CouponDiscount 	     float64  		`json:"couponDiscount,omitempty"`
	DiscountCoupon      *string       `json:"discountCoupon,omitempty"`
	DiscountCouponDetails 		    *e.DiscountCoupons `json:"discountCouponDetails,omitempty"`
	TotalDiscount   	float64 	`json:"totalDiscount,omitempty"`
	TotalAmount     	float64 	`json:"totalAmount,omitempty"`
}
