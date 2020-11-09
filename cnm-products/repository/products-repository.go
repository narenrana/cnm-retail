package repository

import (
	core "shopping-cart/cnm-core"
)

type Service interface {
	List() ([] Product,error)
	FindBy(u Product) (  Product,error)
	Add(u Product) ( Product,error)
	Delete(u Product) ( Product,error)
}

type Product struct {
	ProductId     int       `gorm:"primaryKey" json:"productId,omitempty"`
	ProductName   string    `json:"productName,omitempty"`
	ProductPrice  float64   `json:"productPrice,omitempty"`
	BaseCurrency  string    `json:"baseCurrency,omitempty"`
	ProductTitle  string    `json:"productTitle,omitempty"`
	ProductDesc   string    `json:"productDesc,omitempty"`
	ImageUrl      string    `json:"imageUrl,omitempty"`
}

func (*Product) List() ([] Product, error){
	db,err :=core.GetDB()
	var found [] Product;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*Product) FindBy(user Product) (Product, error){
	db,err :=core.GetDB()
	 var found Product;
	if err != nil {
		return user, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*Product) Add(user Product) (Product, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Create(&user);
	return  user, err;
}

func (*Product) Delete(user Product) (Product, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Delete(&user);
	return  user, err;
}

func ProductRepositoryInstance() Service {
	return &Product{};
}

