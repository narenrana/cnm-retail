package services

import (
	ce "shopping-cart/cnm-carts/entities"
	"shopping-cart/cnm-products/entities"
	repositoryDao "shopping-cart/cnm-products/repository"
)


// Service is the interface that provides booking methods.
type  ProductPriceService interface {
	findCartItemProduct(items []*ce.CartItems) ([]*ce.CartItems, error)
}

type productPriceService struct {
}

func (p productPriceService) findCartItemProduct(items []*ce.CartItems) (  []*ce.CartItems, error) {
	instance:=repositoryDao.ProductRepositoryInstance();
    var productIds []int ;
    for _,item:= range  items {
		productIds = append(productIds,item.ProductId )
	}
	productPriceMap := make(map[int] entities.Product)
	products, err := instance.FindByIds(productIds);
	for _,product:= range  products {
		productPriceMap[ product.ProductId]=product;
	}

	for _, cartItem:= range  items {
		cartItem.Product=productPriceMap[ cartItem.ProductId];
	}


	return items,err
}

// NewService creates a  service with necessary dependencies.
func NewProductPriceService() ProductPriceService {
	return &productPriceService{ }
}


