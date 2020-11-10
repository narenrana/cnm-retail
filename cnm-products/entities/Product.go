package entities


type Product struct {
	ProductId     int       `gorm:"primaryKey" json:"productId,omitempty"`
	ProductName   string    `json:"productName,omitempty"`
	ProductPrice  float64   `json:"productPrice,omitempty"`
	BaseCurrency  string    `json:"baseCurrency,omitempty"`
	ProductTitle  string    `json:"productTitle,omitempty"`
	ProductDesc   string    `json:"productDesc,omitempty"`
	ImageUrl      string    `json:"imageUrl,omitempty"`
}