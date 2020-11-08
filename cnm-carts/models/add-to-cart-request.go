package models

import "shopping-cart/cnm-carts/repository"

type  AddToCartRequest struct {
	Cart repository.Cart;
}

