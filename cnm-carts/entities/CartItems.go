package entities

import e "shopping-cart/cnm-products/entities"

type CartItems struct {
	//TODO adopt some id generator - Time being i am using only timestamp while creating record
	CartItemsId   *int `gorm:"primaryKey" json:"cartItemsId,omitempty"`
	CartId        *int `json:"cartId,omitempty"`
	Quantity      *int  `json:"quantity,omitempty"`
	ProductId     int `json:"productId,omitempty"`
	Currency    string `json:currency,omitempty"`
	Product     e.Product `gorm:"foreignKey:ProductId;references:ProductId" json:"products,omitempty"`
}