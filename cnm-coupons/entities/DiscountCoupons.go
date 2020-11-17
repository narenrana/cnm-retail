package entities

import "time"

type DiscountCoupons struct {
	DiscountCouponsId     int          `gorm:"primaryKey" json:"discountCouponsId,omitempty"`
	Description           string       `json:"description,omitempty"`
	DiscountCoupon        string       `json:"discountCoupon,omitempty"`
	Discount              float64       `json:"discount,omitempty"`
	DiscountMode          string       `json:"discountMode,omitempty"`
	CouponsType          string        `json:"couponsType,omitempty"`
	ExpiryDate            time.Time    `json:"expiryDate,omitempty"`
	DateCreated           time.Time   `json:"dateCreated,omitempty"`
	DateUpdated           time.Time    `json:"dateUpdated,omitempty"`
	Active                bool         `json:"active,omitempty"`
	DiscountCouponsRules  [] *DiscountCouponsRules   `gorm:"foreignKey:DiscountCouponsId;references:DiscountCouponsId" json:"discountCouponsRules,omitempty"`
}
