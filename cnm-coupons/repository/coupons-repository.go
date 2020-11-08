package coupons

import (
	core "shopping-cart/cnm-core"
	"time"
)

type Service interface {
	List() ([] DiscountCoupons,error)
	FindByDiscountCoupon(u string) (  DiscountCoupons,error)
	Add(u DiscountCoupons) ( DiscountCoupons,error)
	Delete(u DiscountCoupons) ( DiscountCoupons,error)
}

type DiscountCoupons struct {

	DiscountCouponsId     int          `gorm:"primaryKey" json:"description,omitempty"`
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

func (*DiscountCoupons) List() ([] DiscountCoupons, error){
	db,err :=core.GetDB()
	var found [] DiscountCoupons;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	for _,v := range found {
		var item [] *DiscountCouponsRules;
		db.Model(&item).Where("DiscountCoupons_id = ?", v.DiscountCouponsId).Find(&item)
		v.DiscountCouponsRules=  item
	}

	return  found, err;
}

func (*DiscountCoupons) FindBy(user DiscountCoupons) (DiscountCoupons, error){
	db,err :=core.GetDB()
	var found  DiscountCoupons;
	if err != nil {
		return user, err;
	}
	db.Find(&found);

	return  found, err;
}

func (*DiscountCoupons) FindByDiscountCoupon(coupon string) (DiscountCoupons, error){
	db,err :=core.GetDB()
	var found  DiscountCoupons;
	if err != nil {
		return DiscountCoupons{}, err;
	}
	db.Model(found).Where("discount_coupon=?",coupon).Find(&found);
	var item [] *DiscountCouponsRules;
	db.Model(&item).Where("discount_coupons_id = ?", found.DiscountCouponsId).Find(&item)
	found.DiscountCouponsRules=item;

	return  found, err;
}

func (*DiscountCoupons) Add(d DiscountCoupons) (DiscountCoupons, error){
	db,err :=core.GetDB()

	if db.Model(&d).Where("cart_id = ?", d.DiscountCouponsId).Updates(&d).RowsAffected == 0 {
		db.Create(&d)
	}else {
		for _,v := range d.DiscountCouponsRules {
			if db.Model(&v).Where("cart_id = ?", v.DiscountCouponsId).Updates(&v).RowsAffected == 0 {
				db.Create(&v)
			}
		}
	}
	if err != nil {
		return d, err;
	}

	return  d, err;
}

func (*DiscountCoupons) Delete(offer DiscountCoupons) (DiscountCoupons, error){
	db,err :=core.GetDB()
	if err != nil {
		return offer, err;
	}
	db.Delete(&offer);
	return  offer, err;
}


func DiscountCouponsRepositoryInstance() Service {
	return &DiscountCoupons{};
}

