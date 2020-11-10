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
	DeleteCartItems(request models.DeleteCartItemRequest) (models.CartResponse, error)

}

type service struct {
}

func (s *service) calculateTotalAmount(items []*cartRepo.CartItems) float64{
	total:=0.00;
	for _, item := range items {
		total=total+ item.Product.ProductPrice * float64(*item.Quantity)
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
	cart, err:=repository.FirstOrCreate(request.UserId);
	return s.parePareCartResponse(cart, err)
}

func (s *service) DeleteCartItems(request models.DeleteCartItemRequest) (models.CartResponse, error) {
	repository:= cartRepo.CartsRepositoryInstance()
	_, _ = repository.DeleteCartItem(request.CartItemIds)
	var getRequest models.GetCartRequest
	getRequest.UserId=request.UserId
	return s.Get(getRequest)
}

func (s *service) parePareCartResponse(cart cartRepo.Cart, err error) (models.CartResponse, error) {
	offersService := NewOffersServiceService()
	productService := NewProductPriceService()
	var offers []*offersRepo.Offers
	var coupon *coupons.DiscountCoupons
	offersDiscount := 0.0
	couponDiscount := 0.0
	//totalDiscount:=0.0
	cartItems, _ :=productService.findCartItemProduct(cart.CartItems);
	cart.CartItems=cartItems;
	if len(cart.CartItems) > 0 {
		couponService := NewCouponServiceService()
		coupon, couponDiscount, _ = couponService.couponDiscount(cart)
		offers, offersDiscount, _ = offersService.getOffers(cartItems)
	}
	return models.CartResponse{
		CartId:                cart.CartId,
		CartName:              cart.CartName,
		CartItems:             cart.CartItems,
		UserId:                cart.UserId,
		AppliedOffers:         offers,
		TotalDiscount:         offersDiscount + couponDiscount,
		//IsCouponApplied:       couponDiscount > 0,
		OffersDiscount:         offersDiscount,
		CouponDiscount:        couponDiscount,
		DiscountCoupon:        cart.DiscountCoupon,
		DiscountCouponDetails: coupon,
		TotalAmount:           s.calculateTotalAmount(cartItems),
	}, err
}

 
// NewService creates a  service with necessary dependencies.
func NewService() Service {
	return &service{}
}

