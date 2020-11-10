package repository

import (
	core "shopping-cart/cnm-core"
	"shopping-cart/cnm-core/wrappers"
	"shopping-cart/cnm-products/repository"
)

type Service interface {
	List() ([] *Cart,error)
	FirstOrCreate(userId int) (  Cart,error)
	Add(u Cart) ( Cart,error)
	Delete(u Cart) ( Cart,error)

}

type Cart struct {
	//TODO adopt some id generator - Time being i am using only timestamp while creating record
	CartId        *int      `gorm:"primaryKey" json:"cartId,omitempty"`
	CartName      string    `json:"cartName,omitempty"`
	UserId        int       `json:"userId,omitempty"`
	CartItems     []*CartItems `gorm:"foreignKey:CartId;references:CartId" json:"cartItems,omitempty"`
	//OffersId      int      `json:"offersId,omitempty"`
	DiscountCoupon *string   `json:"discountCoupon,omitempty"`
}

type CartItems struct {
	//TODO adopt some id generator - Time being i am using only timestamp while creating record
	CartItemsId   *int `gorm:"primaryKey" json:"cartItemsId,omitempty"`
	CartId        *int `json:"cartId,omitempty"`
	Quantity      *int  `json:"quantity,omitempty"`
	ProductId     int `json:"productId,omitempty"`
	Currency    string `json:currency,omitempty"`
	Product     repository.Product `gorm:"foreignKey:ProductId;references:ProductId" json:"products,omitempty"`
}

func (*Cart) List() ([] *Cart, error){
	db,err :=core.GetDB()
	var found [] *Cart;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	for _,v := range found {
		var cartItem [] *CartItems;
		db.Model(&cartItem).Where("cart_id = ?", v.CartId).Find(&cartItem)
		v.CartItems=  cartItem
	}
	return  found, err;
}

func (*Cart) FirstOrCreate(userId int) (Cart, error){
	db,err :=core.GetDB()
	 var found Cart;
	found.UserId= userId
	if err != nil {
		return Cart{}, err;
	}
	db.Model(&found).Where("user_id= ?",userId).FirstOrCreate(&found);
	var cartItem [] *CartItems;
	db.Model(&cartItem).Where("cart_id = ?", found.CartId).Find(&cartItem)
	found.CartItems=cartItem;
	return  found, err;
}


func (*Cart) Add(cart Cart) (Cart, error){
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
    //clean extra records
    var deletedItems CartItems
	db.Model(&deletedItems).Where("cart_items_id not in (?)", cartItemIds).Delete(&deletedItems)
	//db.Delete("cart_items_id not in (?)",cartItemIds, "cart_id=?",cart.CartId)

	if err != nil {
		return cart, err;
	}

	return  cart, err;
}

func (*Cart) Delete(user Cart) (Cart, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Delete(&user);
	return  user, err;
}

func (*Cart) DeleteByIds(user Cart) (Cart, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Where("cart_items_id=?")
	return  user, err;
}


func CartsRepositoryInstance() Service {
	return &Cart{}
}


