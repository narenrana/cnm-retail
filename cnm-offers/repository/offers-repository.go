package repository

import (
	core "shopping-cart/cnm-core"
	e "shopping-cart/cnm-offers/entities"
	
)

type Repository interface {
	List() ([] *e.Offers,error)
	FindBy(u e.Offers) (  e.Offers,error)
	Add(u e.Offers) ( e.Offers,error)
	Delete(u e.Offers) ( e.Offers,error)

}

type repository struct {

}

func (*repository) List() ([] *e.Offers, error){
	db,err :=core.GetDB()
	var found [] *e.Offers;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	for _,v := range found {
		var item [] *e.OffersRules;
		db.Model(&item).Where("offers_id = ?", v.OffersId).Find(&item)
		v.OffersRules=  item
	}

	return  found, err;
}

func (*repository) FindBy(user e.Offers) (e.Offers, error){
	db,err :=core.GetDB()
	var found  e.Offers;
	if err != nil {
		return user, err;
	}
	db.Find(&found);

	return  found, err;
}

func (*repository) Add(cart e.Offers) (e.Offers, error){
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

func (*repository) Delete(offer e.Offers) (e.Offers, error){
	db,err :=core.GetDB()
	if err != nil {
		return offer, err;
	}
	db.Delete(&offer);
	return  offer, err;
}

func OffersRepositoryInstance() Repository {
	return &repository{};
}

