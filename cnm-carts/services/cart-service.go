package services

import (
	models "shopping-cart/cnm-carts/models"
	cartRepo "shopping-cart/cnm-carts/repository"
	coupons "shopping-cart/cnm-coupons/repository"
	offersRepo "shopping-cart/cnm-offers/repository"
)


// Service is the interface that provides booking methods.
type  Service interface {
	Add(request   models.AddToCartRequest) (models.CartResponse, error)
	Get(request models.GetCartRequest) (models.CartResponse, error)

}

type service struct {
}


func (s *service) calculateTotalAmount(items []*cartRepo.CartItems) float64{
	total:=0.00;
	for _, item := range items {
		total=total+ item.ProductPrice * float64(*item.Quantity)
	}
	return total;
}



func (s *service) Add(request models.AddToCartRequest) (models.CartResponse, error) {
	repository:= cartRepo.CartsRepositoryInstance()
	cart, err:=repository.Add(request.Cart)
	//return  models.AddToCartResponse{Cart: cart}, err
	 return s.parePareCartResponse(cart, err);
}

func (s *service) Get(request models.GetCartRequest) (models.CartResponse, error) {
	repository:= cartRepo.CartsRepositoryInstance()
	cart, err:=repository.FirstOrCreate(request.UseId);
	return s.parePareCartResponse(cart, err)
}

func (s *service) parePareCartResponse(cart cartRepo.Cart, err error) (models.CartResponse, error) {
	offersService := NewOffersServiceService()
	var offers []*offersRepo.Offers
	var coupon *coupons.DiscountCoupons
	totalDiscount := 0.0
	couponDiscount := 0.0
	//totalDiscount:=0.0
	if len(cart.CartItems) > 0 {
		couponService := NewCouponServiceService()
		coupon, couponDiscount, _ = couponService.couponDiscount(cart)
		offers, totalDiscount, _ = offersService.getOffers(cart.CartItems)
	}
	return models.CartResponse{
		CartId:                cart.CartId,
		CartName:              cart.CartName,
		CartItems:             cart.CartItems,
		UserId:                cart.UserId,
		AppliedOffers:         offers,
		TotalDiscount:         totalDiscount + couponDiscount,
		IsCouponApplied:       couponDiscount > 0,
		DiscountCoupon:        cart.DiscountCoupon,
		DiscountCouponDetails: coupon,
		TotalAmount:           s.calculateTotalAmount(cart.CartItems),
	}, err
}

 
// NewService creates a  service with necessary dependencies.
func NewService() Service {
	return &service{}
}

