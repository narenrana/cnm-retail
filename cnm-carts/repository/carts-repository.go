package repository

import (
	ce "shopping-cart/cnm-carts/entities"
	core "shopping-cart/cnm-core"
	"shopping-cart/cnm-core/wrappers"
)

type Repository interface {
	List() ([] *ce.Cart,error)
	FirstOrCreate(userId *int) (  ce.Cart,error)
	Add(u ce.Cart) ( ce.Cart,error)
	DeleteCartItem(cartItemIds []int) ([]*ce.CartItems, error)

}

type repository struct {

}



func (*repository) List() ([] *ce.Cart, error){
	db,err :=core.GetDB()
	var found [] *ce.Cart;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	for _,v := range found {
		var cartItem [] *ce.CartItems;
		db.Model(&cartItem).Where("cart_id = ?", v.CartId).Find(&cartItem)
		v.CartItems=  cartItem
	}
	return  found, err;
}

func (*repository) FirstOrCreate(userId *int) (ce.Cart, error){
	db,err :=core.GetDB()
	 var found ce.Cart;
	found.UserId= userId;
	if err != nil {
		return ce.Cart{}, err;
	}
	db.Model(&found).Where("user_id= ?",userId).FirstOrCreate(&found);
	var cartItem [] *ce.CartItems;
	db.Model(&cartItem).Where("cart_id = ?", found.CartId).Find(&cartItem)
	found.CartItems=cartItem;
	return  found, err;
}


func (*repository) Add(cart ce.Cart) (ce.Cart, error){
	db,err :=core.GetDB()
	var cartItemIds[] int ;
	for _,v := range cart.CartItems {
		if v.Quantity == nil {
			v.Quantity = wrappers.IntWrapper(1);
		}

		cartItemIds = append(cartItemIds,v.ProductId )
	}

	if db.Model(&cart).Where("cart_id = ?", cart.CartId).Updates(&cart).RowsAffected == 0 {
		db.Create(&cart)
	}else {
		for _,v := range cart.CartItems {
			if db.Model(&v).Where("cart_id = ?", v.CartId).Updates(&v).RowsAffected == 0 {
				db.Create(&v)
			}
		}
	}

	if err != nil {
		return cart, err;
	}

	return  cart, err;
}



func (*repository) DeleteCartItem(cartItemIds []int) ([]*ce.CartItems, error){
	db,err :=core.GetDB()
	if err != nil {
		return nil, err;
	}
	var cartItem   [] *ce.CartItems


	db.Model(cartItem).Where("cart_items_id in (?)",cartItemIds).Delete(cartItem)
	return  cartItem, err;
}


func CartsRepositoryInstance() Repository {
	return &repository{}
}


