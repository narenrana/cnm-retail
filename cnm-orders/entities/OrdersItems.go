package entities

type OrdersItems struct {
	OrderItemsId   *int    `gorm:"primaryKey" json:"orderItemsId,omitempty"`
	OrderId        *int    `json:"orderId,omitempty"`
	ProductId      int    `json:"productId,omitempty"`
	ProductName    string    `json:"productName,omitempty"`
	ProductPrice    float64    `json:"productPrice,omitempty"`
	Quantity        *int    `json:"quantity,omitempty"`
}
