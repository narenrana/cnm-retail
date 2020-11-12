package e

import (
	core "shopping-cart/cnm-core"
	e "shopping-cart/cnm-orders/entities"
)

type Repository interface {
	List(userId *int) ([] *e.Orders,error)
	FindBy(u e.Orders) (  e.Orders,error)
	Add(u e.Orders) ( e.Orders,error)
	Delete(u e.Orders) ( e.Orders,error)

}

type repository struct {
	 
}

func (*repository) List(userId *int) ([] *e.Orders, error){
	db,err :=core.GetDB()
	var found [] *e.Orders;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	for _,v := range found {
		var items [] e.OrdersItems;
		db.Model(&items).Where("order_id = ?", userId).Find(&items)
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

func (*repository) Add(order e.Orders) (e.Orders, error){
	db,err :=core.GetDB()

	db.Create(&order)

	if err != nil {
		return order, err;
	}

	return order, err;
}

func (*repository) Delete(user e.Orders) (e.Orders, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Delete(&user);
	return  user, err;
}

func OrderRepositoryInstance() Repository {
	return &repository{}
}

