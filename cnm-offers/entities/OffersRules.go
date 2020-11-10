package entities

import "time"

type OffersRules struct {
	//TODO adopt some id generator - Time being, i am using only timestamp while creating record
	OffersRulesId        int        `gorm:"primaryKey" json:"offersRulesId,omitempty"`
	OffersId             int        `json:"offersId,omitempty"`
	Key                  string     `json:"key,omitempty"`
	Value                string     `json:"value,omitempty"`
	Operator             string     `json:"operator,omitempty"`
	Description          string     `json:"description,omitempty"`
	DateCreated          time.Time  `json:"dateCreated,omitempty"`
	DateUpdated         time.Time   `json:"dateUpdated,omitempty"`
}
