package services

import (
	models "shopping-cart/cnm-carts/models"
	repository "shopping-cart/cnm-carts/repository"
)


// Service is the interface that provides booking methods.
type  Service interface {
	Add(request   models.AddToCartRequest) (models.AddToCartResponse, error)
	Get(request models.GetCartRequest) (models.GetCartResponse, error)

}

type service struct {
}


func (s *service) calculateTotalAmount(items []*repository.CartItems) float64{
	total:=0.00;
	for _, item := range items {
		total=total+ item.ProductPrice * float64(*item.Quantity)
	}
	return total;
}



func (s *service) Add(request models.AddToCartRequest) (models.AddToCartResponse, error) {
	repository:= repository.CartsRepositoryInstance()
	cart, err:=repository.Add(request.Cart)
	return  models.AddToCartResponse{Cart: cart}, err
}

func (s *service) Get(request models.GetCartRequest) (models.GetCartResponse, error) {
	repository:= repository.CartsRepositoryInstance()
	cart, err:=repository.FirstOrCreate(request.UseId);
	offersService:=NewOffersServiceService()
	 couponService:=NewCouponServiceService()
	 coupon,couponDiscount, err:=couponService.couponDiscount(cart)
	offers, totalDiscount, err:=offersService.getOffers(cart.CartItems)
	return models.GetCartResponse{
		Cart: cart,
		AppliedOffers: offers,
		TotalDiscount: totalDiscount+couponDiscount,
		IsCouponApplied: couponDiscount>0,
		DiscountCoupon: coupon,
		TotalAmount: s.calculateTotalAmount(cart.CartItems),
	}, err
}

 
// NewService creates a  service with necessary dependencies.
func NewService() Service {
	return &service{}
}

