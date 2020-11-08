package repository

import (
	core "shopping-cart/cnm-core"
	"time"
)

type Service interface {
	List() ([] *Offers,error)
	FindBy(u Offers) (  Offers,error)
	Add(u Offers) ( Offers,error)
	Delete(u Offers) ( Offers,error)

}


type Offers struct {
	//TODO adopt some id generator - Time being, i am using only timestamp while creating record
	//json:"offer_id,omitempty",
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
}

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

func (*Offers) List() ([] *Offers, error){
	db,err :=core.GetDB()
	var found [] *Offers;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	for _,v := range found {
		var item [] *OffersRules;
		db.Model(&item).Where("offers_id = ?", v.OffersId).Find(&item)
		v.OffersRules=  item
	}

	return  found, err;
}

func (*Offers) FindBy(user Offers) (Offers, error){
	db,err :=core.GetDB()
	var found  Offers;
	if err != nil {
		return user, err;
	}
	db.Find(&found);

	return  found, err;
}

func (*Offers) Add(cart Offers) (Offers, error){
	db,err :=core.GetDB()

	if db.Model(&cart).Where("cart_id = ?", cart.OffersId).Updates(&cart).RowsAffected == 0 {
		db.Create(&cart)
	}else {
		for _,v := range cart.OffersRules {
			if db.Model(&v).Where("cart_id = ?", v.OffersId).Updates(&v).RowsAffected == 0 {
				db.Create(&v)
			}
		}
	}
	if err != nil {
		return cart, err;
	}

	return  cart, err;
}

func (*Offers) Delete(offer Offers) (Offers, error){
	db,err :=core.GetDB()
	if err != nil {
		return offer, err;
	}
	db.Delete(&offer);
	return  offer, err;
}

func OffersRepositoryInstance() Service {
	return &Offers{};
}

