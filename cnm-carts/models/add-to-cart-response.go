package models

import "shopping-cart/cnm-carts/entities"

type AddToCartResponse struct {
	Cart          entities.Cart   `json:"product,omitempty"`
	Err            error  `json:"error,omitempty"`
}
