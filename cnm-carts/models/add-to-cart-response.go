package models

import "shopping-cart/cnm-carts/repository"

type AddToCartResponse struct {
	Cart          repository.Cart   `json:"product,omitempty"`
	Err            error  `json:"error,omitempty"`
}
