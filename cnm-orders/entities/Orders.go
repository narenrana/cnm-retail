package entities

type Orders struct {
	//TODO adopt some id generator - Time being i am using only timestamp while creating record
	OrderId       *int           `gorm:"primaryKey" json:"order_id,omitempty"`
	CartId        *int         `json:"cart_id,omitempty"`
	UserId        *int            `json:"userId,omitempty"`
	Amount        float64        `json:"amount,omitempty"`
	Status        string         `json:"status,omitempty"`
	active        bool           `json:"active,omitempty"`
	OrderData     string         `json:"orderData,omitempty"`
	OrdersItems   []OrdersItems  `gorm:"foreignKey:OrderId;references:OrderId" json:"ordersItems,omitempty"`
}