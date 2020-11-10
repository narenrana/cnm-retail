package entities



type Cart struct {
	//TODO adopt some id generator - Time being i am using only timestamp while creating record
	CartId        *int      `gorm:"primaryKey" json:"cartId,omitempty"`
	CartName      string    `json:"cartName,omitempty"`
	UserId        *int       `json:"userId,omitempty"`
	CartItems     []*CartItems `gorm:"foreignKey:CartId;references:CartId" json:"cartItems,omitempty"`
	//OffersId      int      `json:"offersId,omitempty"`
	DiscountCoupon *string   `json:"discountCoupon,omitempty"`
}