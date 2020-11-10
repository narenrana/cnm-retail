package coupons

import (
	 
	core "shopping-cart/cnm-core"
	e "shopping-cart/cnm-coupons/entities"
)


type Repository interface {
	List() ([]  e.DiscountCoupons,error)
	FindByDiscountCoupon(u string) (  e.DiscountCoupons,error)
	Add(u e.DiscountCoupons) ( e.DiscountCoupons,error)
	Delete(u e.DiscountCoupons) ( e.DiscountCoupons,error)
}


type repository struct {

}


func (*repository) List() ([] e.DiscountCoupons, error){
	db,err :=core.GetDB()
	var found [] e.DiscountCoupons;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	for _,v := range found {
		var item [] *e.DiscountCouponsRules;
		db.Model(&item).Where("DiscountCoupons_id = ?", v.DiscountCouponsId).Find(&item)
		v.DiscountCouponsRules=  item
	}

	return  found, err;
}

func (*repository) FindBy(user e.DiscountCoupons) (e.DiscountCoupons, error){
	db,err :=core.GetDB()
	var found  e.DiscountCoupons;
	if err != nil {
		return user, err;
	}
	db.Find(&found);

	return  found, err;
}

func (*repository) FindByDiscountCoupon(coupon string) (e.DiscountCoupons, error){
	db,err :=core.GetDB()
	var found  e.DiscountCoupons;
	if err != nil {
		return e.DiscountCoupons{}, err;
	}
	db.Model(found).Where("discount_coupon=?",coupon).Find(&found);

	if &found==nil {
		return found,nil
	}
	var item [] *e.DiscountCouponsRules;
	db.Model(&item).Where("discount_coupons_id = ?", found.DiscountCouponsId).Find(&item)
	found.DiscountCouponsRules=item;

	return  found, err;
}

func (*repository) Add(d e.DiscountCoupons) (e.DiscountCoupons, error){
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

func (*repository) Delete(offer e.DiscountCoupons) (e.DiscountCoupons, error){
	db,err :=core.GetDB()
	if err != nil {
		return offer, err;
	}
	db.Delete(&offer);
	return  offer, err;
}


func  RepositoryInstance() Repository {
	return &repository{};
}

