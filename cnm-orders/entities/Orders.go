package entities

import "github.com/jackc/pgtype"

type Orders struct {
	//TODO adopt some id generator - Time being i am using only timestamp while creating record
	OrderId       *int           `gorm:"primaryKey" json:"order_id,omitempty"`
	CartId        string         `json:"cart_id,omitempty"`
	UserId        int            `json:"userId,omitempty"`
	Amount        float32        `json:"amount,omitempty"`
	Status        string         `json:"status,omitempty"`
	active        bool           `json:"active,omitempty"`
	DateCreated   pgtype.Date    `json:"dateCreated,omitempty"`
	DateUpdated   pgtype.Date    `json:"dateUpdated,omitempty"`
	OrdersItems     []*OrdersItems  `gorm:"foreignKey:OrderId;references:OrderId" json:"ordersItems,omitempty"`
}