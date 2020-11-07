package orders

import (
	"github.com/jackc/pgtype"
	core "shopping-cart/cnm-core"
)

type Service interface {
	List() ([] *Orders,error)
	FindBy(u Orders) (  Orders,error)
	Add(u Orders) ( Orders,error)
	Delete(u Orders) ( Orders,error)

}

type Orders struct {
	//TODO adopt some id generator - Time being i am using only timestamp while creating record
	OrderId       *int           `gorm:"primaryKey" json:"order_id,omitempty"`
	CartId        string         `json:"cart_id,omitempty"`
	UserId        int            `json:"userId,omitempty"`
	Amount        float32        `json:"amount,omitempty"`
	Status        string         `json:"status,omitempty"`
	active        bool           `json:"active,omitempty"`
	DateCreated   pgtype.Date    `json:"dateCreated,omitempty"`
	DateUpdated   pgtype.Date    `json:"dateUpdated,omitempty"`
	OrdersItems     []*OrdersItems  `gorm:"foreignKey:OrderId;references:OrderId" json:"ordersItems,omitempty"`
}

type OrdersItems struct {
	//TODO adopt some id generator - Time being i am using only timestamp while creating record
	OrderItemsId   *int    `gorm:"primaryKey" json:"orderItemsId,omitempty"`
	OrderId        *int    `json:"orderId,omitempty"`
	ProductId      *int    `json:"productId,omitempty"`
	ProductName     int    `json:"productName,omitempty"`
	ProductPrice    int    `json:"productPrice,omitempty"`
	DiscountId     int     `json:"discountId,omitempty"`
	DateCreated   pgtype.Date    `json:"dateCreated,omitempty"`
	DateUpdated   pgtype.Date    `json:"dateUpdated,omitempty"`
}

func (*Orders) List() ([] *Orders, error){
	db,err :=core.GetDB()
	var found [] *Orders;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	for _,v := range found {
		var items [] *OrdersItems;
		db.Model(&items).Where("cart_id = ?", v.CartId).Find(&items)
		v.OrdersItems=  items
	}
	return  found, err;
}

func (*Orders) FindBy(u Orders) (Orders, error){
	db,err :=core.GetDB()
	 var found Orders;
	if err != nil {
		return u, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*Orders) Add(cart Orders) (Orders, error){
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

func (*Orders) Delete(user Orders) (Orders, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Delete(&user);
	return  user, err;
}

func CartsRepositoryInstance() Service {
	return &Orders{}
}

