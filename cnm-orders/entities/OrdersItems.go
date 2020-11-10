package entities

import "github.com/jackc/pgtype"

type OrdersItems struct {
	//TODO adopt some id generator - Time being i am using only timestamp while creating record
	OrderItemsId   *int    `gorm:"primaryKey" json:"orderItemsId,omitempty"`
	OrderId        *int    `json:"orderId,omitempty"`
	ProductId      *int    `json:"productId,omitempty"`
	ProductName     int    `json:"productName,omitempty"`
	ProductPrice    int    `json:"productPrice,omitempty"`
	DiscountId     int     `json:"discountId,omitempty"`
	DateCreated   pgtype.Date    `json:"dateCreated,omitempty"`
	DateUpdated   pgtype.Date    `json:"dateUpdated,omitempty"`
}
