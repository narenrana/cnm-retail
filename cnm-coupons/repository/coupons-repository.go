package coupons

import (
	"github.com/jackc/pgtype"
	core "shopping-cart/cnm-core"
)

type Service interface {
	List() ([] DiscountCoupons,error)
	FindBy(u DiscountCoupons) (  DiscountCoupons,error)
	Add(u DiscountCoupons) ( DiscountCoupons,error)
	Delete(u DiscountCoupons) ( DiscountCoupons,error)
}

type DiscountCoupons struct {

	DiscountCouponsId     int          `gorm:"primaryKey" json:"description,omitempty"`
	Description           string       `json:"description,omitempty"`
	DiscountCoupon        string       `json:"discountCoupon,omitempty"`
	Discount              string       `json:"discount,omitempty"`
	DiscountMode          string       `json:"discountMode,omitempty"`
	ProductId             string       `json:"productId,omitempty"`
	CouponsExpiry         string       `json:"couponsExpiry,omitempty"`
	DateCreate            pgtype.Date  `json:"dateCreate,omitempty"`
	DateUpdated           pgtype.Date  `json:"dateUpdated,omitempty"`
	active                bool         `json:"active,omitempty"`
}

func (*DiscountCoupons) List() ([] DiscountCoupons, error){
	db,err :=core.GetDB()
	var found [] DiscountCoupons;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*DiscountCoupons) FindBy(item DiscountCoupons) (DiscountCoupons, error){
	db,err :=core.GetDB()
	 var found DiscountCoupons;
	if err != nil {
		return item, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*DiscountCoupons) Add(item DiscountCoupons) (DiscountCoupons, error){
	db,err :=core.GetDB()
	if err != nil {
		return item, err;
	}
	db.Create(&item);
	return  item, err;
}

func (*DiscountCoupons) Delete(item DiscountCoupons) (DiscountCoupons, error){
	db,err :=core.GetDB()
	if err != nil {
		return item, err;
	}
	db.Delete(&item);
	return  item, err;
}

func CouponsRepositoryInstance() Service {
	return &DiscountCoupons{};
}

