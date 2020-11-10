package entities

import "time"

type DiscountCouponsRules struct {
	//TODO adopt some id generator - Time being, i am using only timestamp while creating record
	DiscountCouponsRulesId        int        `gorm:"primaryKey" json:"discountCouponsRulesId,omitempty"`
	DiscountCouponsId             int        `json:"discountCouponsId,omitempty"`
	Key                           string     `json:"key,omitempty"`
	Value                         string     `json:"value,omitempty"`
	Operator                      string     `json:"operator,omitempty"`
	Description                   string     `json:"description,omitempty"`
	DateCreated                   time.Time  `json:"dateCreated,omitempty"`
	DateUpdated                   time.Time   `json:"dateUpdated,omitempty"`
}