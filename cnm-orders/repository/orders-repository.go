package e

import (
	core "shopping-cart/cnm-core"
	e "shopping-cart/cnm-orders/entities"
)

type Repository interface {
	List() ([] *e.Orders,error)
	FindBy(u e.Orders) (  e.Orders,error)
	Add(u e.Orders) ( e.Orders,error)
	Delete(u e.Orders) ( e.Orders,error)

}

type repository struct {
	 
}

func (*repository) List() ([] *e.Orders, error){
	db,err :=core.GetDB()
	var found [] *e.Orders;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	for _,v := range found {
		var items [] *e.OrdersItems;
		db.Model(&items).Where("cart_id = ?", v.CartId).Find(&items)
		v.OrdersItems=  items
	}
	return  found, err;
}

func (*repository) FindBy(u e.Orders) (e.Orders, error){
	db,err :=core.GetDB()
	 var found e.Orders;
	if err != nil {
		return u, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*repository) Add(cart e.Orders) (e.Orders, error){
	db,err :=core.GetDB()

	if db.Model(&cart).Where("cart_id = ?", cart.CartId).Updates(&cart).RowsAffected == 0 {
		db.Create(&cart)
	}else {
		for _,v := range cart.OrdersItems {
			if db.Model(&v).Where("cart_id = ?", v.OrderId).Updates(&v).RowsAffected == 0 {
				db.Create(&v)
			}
		}
	}
	if err != nil {
		return cart, err;
	}

	return  cart, err;
}

func (*repository) Delete(user e.Orders) (e.Orders, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Delete(&user);
	return  user, err;
}

func CartsRepositoryInstance() Repository {
	return &repository{}
}

