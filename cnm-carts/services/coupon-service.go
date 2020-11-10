package services


import (
	"log"
	repository "shopping-cart/cnm-carts/repository"
	"shopping-cart/cnm-core/constants"
	e "shopping-cart/cnm-coupons/entities"
	ec "shopping-cart/cnm-coupons/repository"

	"strconv"
)

// Service is the interface that provides booking methods.
type  CouponService interface {
    couponDiscount(cart repository.Cart ) (*e.DiscountCoupons,float64, error)
    findDiscount(items []*repository.CartItems, discountCoupon e.DiscountCoupons) (*e.DiscountCoupons, float64, error)
	productsAndQuantityMap(items []*repository.CartItems, productAndQuantityMap map[string]int, productPriceMapping map[string]float64)
    productDiscountRuleMap( coupon e.DiscountCoupons, itemMap map[string]int)
}

type couponService struct {
}




func (s *couponService) couponDiscount(cart repository.Cart ) (*e.DiscountCoupons,float64, error) {
	couponsRepositoryInstance:= ec.RepositoryInstance();

	if cart.DiscountCoupon==nil {
		return nil,0.0, nil

	}

	discountCoupon, error :=couponsRepositoryInstance.FindByDiscountCoupon(*cart.DiscountCoupon);
	if error!=nil {
		return nil,0.0, nil

	}

	if discountCoupon.CouponsType  ==constants.ONE_TIME  {
			couponDiscount,discount, err:=s.findDiscount(cart.CartItems,discountCoupon)
			return couponDiscount, discount,err
			}

    if discountCoupon.CouponsType  ==constants.MULTI_USE  {
			//TODO
		}


	return &discountCoupon,0.0, nil;
}


func (s *couponService)  findDiscount(items []*repository.CartItems, discountCoupon e.DiscountCoupons) (*e.DiscountCoupons, float64, error){

	log.Println("-------------------Product Coupon Discount Service--------------------------")
	//Prepare Key Value of Product and quantity from rules
	productRuleMap := make(map[string]int)
	log.Println("-------------------Product Rule and Operator --------------------------")
	s.productDiscountRuleMap(discountCoupon, productRuleMap)
	if len(productRuleMap) !=1 {
		log.Println("Empty or Multiple Products found for INDIVIDUAL OFFER.")
		return  &e.DiscountCoupons{},0.0, nil;
	}
	productAndQuantityMap := make(map[string]int)
	//Prepare map of product and price
	productPriceMapping := make(map[string]float64)
	s.productsAndQuantityMap(items, productAndQuantityMap, productPriceMapping)


	log.Println("-------------------Product Quantities Mapping--------------------------")
	//TODO - remove after testing
	for key, value := range productAndQuantityMap { log.Println("Key-", key, " , Value-", value) }
	log.Println("-------------------Product Price Mapping--------------------------")
	//TODO remove after testing
	for key, value := range productPriceMapping { log.Println("Key-", key, ",Value-", value)}

	discountRule:=discountCoupon.DiscountCouponsRules[0]
	discount := 0.0
	log.Println("couponsRule.Operator->",(discountRule.Operator==constants.GREATER_THAN_OR_EQUALS_TO))
	log.Println(productAndQuantityMap[discountRule.Key],">=",productAndQuantityMap[discountRule.Key]>= productRuleMap[discountRule.Key],  productRuleMap[discountRule.Key] )
	if discountRule.Operator==constants.GREATER_THAN_OR_EQUALS_TO && productAndQuantityMap[discountRule.Key]>= productRuleMap[discountRule.Key] {
		log.Println("-------------------Lets find total discount of type - ",constants.GREATER_THAN_OR_EQUALS_TO)
		totalEligibleAmountForDiscount:=float64(productAndQuantityMap[discountRule.Key]) * productPriceMapping[discountRule.Key];
		discount=totalEligibleAmountForDiscount * discountCoupon.Discount / 100
	}
	log.Println("Total coupon discount:", discount)
	return &discountCoupon ,discount, nil;
}



func (s *couponService) productsAndQuantityMap(items []*repository.CartItems, productAndQuantityMap map[string]int, productPriceMapping map[string]float64) {
	for _, item := range items {
		productAndQuantityMap[item.Product.ProductName] = productAndQuantityMap[item.Product.ProductName] + *item.Quantity
		productPriceMapping[item.Product.ProductName] = item.Product.ProductPrice

	}
}

func (s *couponService) productDiscountRuleMap( coupon e.DiscountCoupons, itemMap map[string]int) {
	for _, rule := range coupon.DiscountCouponsRules {
		ruleValue, _ := strconv.ParseInt(rule.Value, 10, 64)
		itemMap[rule.Key] = int(ruleValue)
		log.Println("Key-", rule.Key, "Value-", itemMap[rule.Key])
	}
}


// NewService creates a  service with necessary dependencies.
func NewCouponServiceService() CouponService {
	return &couponService{}
}

