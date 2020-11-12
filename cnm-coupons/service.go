package coupons

import (
	"errors"
	"math/rand"
	"shopping-cart/cnm-coupons/entities"
	cr "shopping-cart/cnm-coupons/repository"
	"strconv"
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


	var coupons [] entities.DiscountCoupons;
	var rules [] *entities.DiscountCouponsRules;

	rules = append(rules,&entities.DiscountCouponsRules{
		DiscountCouponsId:      0,
		Key:                    "Oranges",
		Value:                  "0",
		Operator:               ">=",
		Description:            "30% Discount on Oranges",
	});

	targetTime:=time.Now().UTC();
	expireTime:=targetTime.Add(time.Second*10)
	for i := 0; i < request.Quantity; i++ {

		couponValue:="UUIDISC"
		randomTime := rand.Int63n(time.Now().Unix() - 94608000) + 94608000
		couponValue=couponValue+strconv.FormatInt(randomTime, 10)
		coupon:=entities.DiscountCoupons{
			//DiscountCouponsId:    0, //TDODO
			Description:          "30% Discount on Oranges",
			DiscountCoupon:       couponValue,
			Discount:             30,
			DiscountMode:         "PERCENTILE",
			CouponsType:          "ONE_TIME",
			ExpiryDate:           expireTime,
			Active:               true,
			DiscountCouponsRules: rules,
		}
		coupons = append(coupons,coupon )
	}
	for _, coupon := range coupons {
		 instance.Add(coupon);
	}


	return  discountCouponsResponse{DiscountCoupons: coupons,TargetTime: targetTime},nil
}

func (s *service) list() (discountCouponsListResponse, error) {
	instance:= cr.RepositoryInstance()
	product, err:=instance.List();
	return discountCouponsListResponse{Product: product , Err: err}, err
}

func (s *service) find(request findCouponRequest) (findCouponResponse,error){
	instance:= cr.RepositoryInstance()
	coupon, err:=instance.FindByDiscountCoupon(request.Coupon);
	targetTime:=time.Now().UTC();

	isValid:= coupon.Active &&  targetTime.Before(coupon.ExpiryDate.UTC())
	return findCouponResponse{ DiscountCoupons: coupon, Valid:isValid,Time: targetTime  },err

}

// NewService creates a auth service with necessary dependencies.
func NewService() Service {
	return &service{}
}


