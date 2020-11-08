package repository

import (
	core "shopping-cart/cnm-core"
	"shopping-cart/cnm-core/wrappers"
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
	ProductPrice  float64 `json:"productPrice,omitempty"`
	ProductName  string `json:"productName,omitempty"`
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

	for _,v := range cart.CartItems {
		if v.Quantity == nil {
			v.Quantity = wrappers.IntWrapper(1);
		}
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

func (*Cart) Delete(user Cart) (Cart, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Delete(&user);
	return  user, err;
}

func CartsRepositoryInstance() Service {
	return &Cart{}
}


