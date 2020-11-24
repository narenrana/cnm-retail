package entities

import "time"

type Offers struct {
	OffersId           int             `gorm:"primaryKey" json:"offersId,omitempty"`
	Description        string          `json:"description,description"`
	Discount           float64          `json:"discount,omitempty"`
	DiscountMode       string          `json:"discountMode,omitempty"`
	OffersType         string          `json:"offersType,omitempty"`
	active             string          `json:"active,omitempty"`
	expireDate         time.Time       `json:"expireDate,omitempty"`
	DateCreated        time.Time     `json:"dateCreated,omitempty"`
	DateUpdated        time.Time    `json:"dateUpdated,omitempty"`
	OffersRules        [] *OffersRules `gorm:"foreignKey:OffersId;references:OffersId" json:"offersRules,omitempty"`
	offerTotalAmount   *float64          `gorm:"-" json:"discount,omitempty"`
}
