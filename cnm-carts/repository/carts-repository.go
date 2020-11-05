package repository

import (
	core "shopping-cart/cnm-core"
)

type Service interface {
	List() ([] Cart,error)
	FindBy(u Cart) (  Cart,error)
	Add(u Cart) ( Cart,error)
	Delete(u Cart) ( Cart,error)

}

type Cart struct {
	CartId        uint
	CartName      string
	UserId        int
	CartItems     CartItems
}

type CartItems struct {
	ItemId       uint
	CartId        int
	ProductId     int
	ProductPrice  string
	ProductName  string
}

func (*Cart) List() ([] Cart, error){
	db,err :=core.GetDB()
	var found [] Cart;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*Cart) FindBy(user Cart) (Cart, error){
	db,err :=core.GetDB()
	 var found Cart;
	if err != nil {
		return user, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*Cart) Add(user Cart) (Cart, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Create(&user);
	return  user, err;
}

func (*Cart) Delete(user Cart) (Cart, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Delete(&user);
	return  user, err;
}

func CartsRepositoryInstance() Service {
	return &Cart{};
}

