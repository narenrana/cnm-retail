package models

import "time"

type OrderDetails struct {
	OrderId  *int  `json:"orderId"`
	Price   float64 `json:"price"`
	Date    time.Time `json:"price"`
	UserId     *int  `json:"userId"`
}
