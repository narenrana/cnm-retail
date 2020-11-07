package repository

import (
	"github.com/shopspring/decimal"
	core "shopping-cart/cnm-core"
)

type Service interface {
	List() ([] *Cart,error)
	FindBy(u Cart) (  Cart,error)
	Add(u Cart) ( Cart,error)
	Delete(u Cart) ( Cart,error)

}

type Cart struct {
	//TODO adopt some id generator - Time being i am using only timestamp while creating record
	CartId        *int      `gorm:"primaryKey" json:"cartId,omitempty"`
	CartName      string    `json:"cartName,omitempty"`
	UserId        int       `json:"userId,omitempty"`
	CartItems     []*CartItems `gorm:"foreignKey:CartId;references:CartId" json:"cartItems,omitempty"`
	OffersId      int      `json:"offersId,omitempty"`
	DiscountCoupon string   `json:"discountCoupon,omitempty"`
}

type CartItems struct {
	//TODO adopt some id generator - Time being i am using only timestamp while creating record
	CartItemsId   *int `gorm:"primaryKey" json:"cartItemsId,omitempty"`
	CartId        *int `json:"cartId,omitempty"`
	Quantity      *int  `json:"quantity,omitempty"`
	ProductId     int `json:"productId,omitempty"`
	ProductPrice  decimal.Decimal `json:"productPrice,omitempty"`
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

func (*Cart) FindBy(user Cart) (Cart, error){
	db,err :=core.GetDB()
	 var found Cart;
	if err != nil {
		return user, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*Cart) Add(cart Cart) (Cart, error){
	db,err :=core.GetDB()

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

