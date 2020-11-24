package entities

import "time"

type OffersRules struct {
	OffersRulesId        int        `gorm:"primaryKey" json:"offersRulesId,omitempty"`
	OffersId             int        `json:"offersId,omitempty"`
	Key                  string     `json:"key,omitempty"`
	Value                string     `json:"value,omitempty"`
	Operator             string     `json:"operator,omitempty"`
	Description          string     `json:"description,omitempty"`
	DateCreated          time.Time  `json:"dateCreated,omitempty"`
	DateUpdated         time.Time   `json:"dateUpdated,omitempty"`
}
