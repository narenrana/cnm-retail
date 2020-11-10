package entities

import "github.com/jackc/pgtype"

type Payments struct {

	PaymentId         int          `gorm:"primaryKey" json:"description,omitempty"`
	UserId            string       `json:"userId,omitempty"`
	PaymentMode       string       `json:"paymentMode,omitempty"`
	OrderId           string       `json:"orderId,omitempty"`
	OrderDate         string       `json:"orderDate,omitempty"`
	OrderAmount       string       `json:"orderAmount,omitempty"`
	PaymentStatus     string       `json:"PaymentStatus,omitempty"`
	NameOnCard        pgtype.Date  `json:"nameOnCard,omitempty"`
	CardType          pgtype.Date  `json:"cardType,omitempty"`
	CardLastDigit     string       `json:"cardLastDigit,omitempty"`
	active            bool         `json:"active,omitempty"`
	DateCreated   pgtype.Date    `json:"dateCreated,omitempty"`
	DateUpdated   pgtype.Date    `json:"dateUpdated,omitempty"`
}
