package models

type DeleteCartItemRequest struct {
	CartItemIds   []int `json:"cartItemIds,omitempty"`
	CartId        *int `json:"cartId,omitempty"`
	UserId        *int `json:"userId,omitempty"`
}