package services

import (
	"errors"
	"fmt"
	"log"
	"shopping-cart/cnm-carts/repository"
	"shopping-cart/cnm-core/constants"
	oe "shopping-cart/cnm-offers/entities"
	offersRepository "shopping-cart/cnm-offers/repository"
	"sort"
	"strconv"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides booking methods.
type  OffersService interface {
	getOffers(items [] *repository.CartItems) ([] *oe.Offers,float64, error)
	comboOfferDiscount(items [] *repository.CartItems,offer *oe.Offers)  ( *oe.Offers ,float64, error)
	individualOfferDiscount(items [] *repository.CartItems,offer *oe.Offers )  ( *oe.Offers, float64, error)

}

type offersService struct {
}

func (s *offersService) getOffers(items [] *repository.CartItems) ([] *oe.Offers,float64, error){

	offersRepositoryInstance:=offersRepository.OffersRepositoryInstance()
	offers, error :=offersRepositoryInstance.List();

	if error!=nil {
		return offers, 0, error

	}
	var offersFound [] *oe.Offers;
	var totalDiscount float64;

	for _, offer := range offers {
		fmt.Print(offer)
		if offer.OffersType ==constants.COMBO_OFFER  {
			comboOfferDiscount,discount, err:=s.comboOfferDiscount(items,offer)
			totalDiscount=totalDiscount+discount;
			offersFound = append(offersFound, comboOfferDiscount);
			error=err

		}

		if offer.OffersType ==constants.INDIVIDUAL_ITEM_OFFER {
			comboOfferDiscount,discount, err:=s.individualOfferDiscount(items,offer)
			totalDiscount = totalDiscount+discount;
			error=err
			offersFound = append(offersFound, comboOfferDiscount);
		}
	}

	return offersFound, totalDiscount, error
}


func (s *offersService) comboOfferDiscount(items [] *repository.CartItems, offer *oe.Offers )  ( *oe.Offers,float64, error) {
	log.Println("-------------------Product combo Offer Discount Service--------------------------")
	//Prepare Key Value of Product and quantity from rules
	productComboMap := make(map[string]int)
	log.Println("-------------------Product Combo Rules Mapping--------------------------")
	s.productRuleMap(offer, productComboMap)
	//Prepare Key Value of Product and quantity from cart

	productAndQuantityMap := make(map[string]int)
	//Prepare map of product and price
	productPriceMapping := make(map[string]float64)
	s.productAndQuantityMap(items, productAndQuantityMap, productPriceMapping)

	log.Println("-------------------Product Quantities Mapping--------------------------")
	//TODO - remove after testing
	for key, value := range productAndQuantityMap { log.Println("Key-", key, " , Value-", value) }
	log.Println("-------------------Product Price Mapping--------------------------")
	//TODO remove after testing
	for key, value := range productPriceMapping { log.Println("Key-", key, ",Value-", value)}

	//Prepare Key Value of Product and possible individual pairs
	individualCombination := make(map[string]int)
	s.individualPossiblePairs(productComboMap, individualCombination, productAndQuantityMap)

	log.Println("-------------------Print individual Combinations--------------------------")
	//TODO remove after testing
	for key, value := range individualCombination { log.Println("Key-", key, " , Value-", value) }
	log.Println("-------------------Lets Find Min Possible Set--------------------------")
	//Find Min of all pairs
	minPossibleComboCount := s.maxPossibleCombo(individualCombination)

	log.Println("Possible Combo-", minPossibleComboCount)
	log.Println("-------------------Lets find total discount--------------------------")
	if minPossibleComboCount > 0 {
		ruleItemsPricePerCombo := 0.0
		for key, value := range productComboMap {
			itemPrice := productPriceMapping[key]
			ruleItemsPricePerCombo = ruleItemsPricePerCombo + float64(value)*itemPrice
		}

		totalEligibleAmountForDiscount := ruleItemsPricePerCombo * float64(minPossibleComboCount)
		log.Println("Total eligible amount for discount- ", totalEligibleAmountForDiscount)
		//TODO more rules can be added for other discount modes
		discount := 0.0
		if offer.DiscountMode == constants.PERCENTILE {
			discount = totalEligibleAmountForDiscount * offer.Discount / 100
		}

		log.Println("Total Combo discount : ", discount)

		return offer, discount, nil
	}

	return nil, 0.0, nil
}

func (s *offersService) maxPossibleCombo(individualCombination map[string]int) int {

	if len(individualCombination) <=0 {
		return 0
	}

	var valueArray []int
	for key, value:= range individualCombination {
		log.Println("Key:", key, "Value:", value)
		valueArray = append(valueArray,  value)
	}
	sort.Slice(valueArray, func(i, j int) bool {
		return valueArray[i]< valueArray[j]
	})

	minPossibleComboCount := valueArray[0]

	return minPossibleComboCount
}

func (s *offersService) individualPossiblePairs(productComboMap map[string]int, individualCombination map[string]int, productAndQuantityMap map[string]int) {
	for key, value := range productComboMap {
		individualCombination[key] = productAndQuantityMap[key] / value
	}
}

func (s *offersService) productAndQuantityMap(items []*repository.CartItems, productAndQuantityMap map[string]int, productPriceMapping map[string]float64) {
	for _, item := range items {
		productAndQuantityMap[item.Product.ProductName] = productAndQuantityMap[item.Product.ProductName] + *item.Quantity
		productPriceMapping[item.Product.ProductName] = item.Product.ProductPrice

	}
}

func (s *offersService) productRuleMap(offer *oe.Offers, itemMap map[string]int) {
	for _, rule := range offer.OffersRules {
		ruleValue, _ := strconv.ParseInt(rule.Value, 10, 64)
		itemMap[rule.Key] = int(ruleValue)
		log.Println("Key-", rule.Key, "Value-", itemMap[rule.Key])
	}
}

func (s *offersService)  individualOfferDiscount(items []*repository.CartItems, offer *oe.Offers) (*oe.Offers, float64, error){

	log.Println("-------------------Product Individual Offer Discount Service--------------------------")
	//Prepare Key Value of Product and quantity from rules
	productRuleMap := make(map[string]int)
	log.Println("-------------------Product Rule and Operator --------------------------")
	s.productRuleMap(offer, productRuleMap)
	if len(productRuleMap) !=1 {
		log.Println("Empty or Multiple Products found for INDIVIDUAL OFFER.")
		return  nil,0.0, nil;
	}
	productAndQuantityMap := make(map[string]int)
	//Prepare map of product and price
	productPriceMapping := make(map[string]float64)
	s.productAndQuantityMap(items, productAndQuantityMap, productPriceMapping)

	log.Println("-------------------Product Quantities Mapping--------------------------")
	//TODO - remove after testing
	for key, value := range productAndQuantityMap { log.Println("Key-", key, " , Value-", value) }
	log.Println("-------------------Product Price Mapping--------------------------")
	//TODO remove after testing
	for key, value := range productPriceMapping { log.Println("Key-", key, ",Value-", value)}

	offersRule:=offer.OffersRules[0]
	discount := 0.0
	log.Println("offersRule.Operator->",(offersRule.Operator==constants.GREATER_THAN_OR_EQUALS_TO))
	log.Println(productAndQuantityMap[offersRule.Key],">=",productAndQuantityMap[offersRule.Key]>= productRuleMap[offersRule.Key],  productRuleMap[offersRule.Key] )
	if offersRule.Operator==constants.GREATER_THAN_OR_EQUALS_TO && productAndQuantityMap[offersRule.Key]>= productRuleMap[offersRule.Key] {
		log.Println("-------------------Lets find total discount of type - ",constants.GREATER_THAN_OR_EQUALS_TO)
		totalEligibleAmountForDiscount:=float64(productAndQuantityMap[offersRule.Key]) * productPriceMapping[offersRule.Key];
		discount=totalEligibleAmountForDiscount * offer.Discount / 100
	}
	log.Println("Total Individual discount : ", discount)
	return offer,discount, nil;
}



// NewService creates a  service with necessary dependencies.
func NewOffersServiceService() OffersService {
	return &offersService{}
}

