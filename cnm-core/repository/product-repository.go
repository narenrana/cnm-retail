package repository

import (
	core "shopping-cart/cnm-core"
)

type ProductRepository interface {
	 get() [] Product
	 findBy(product Product)
	 add(product Product)
	 delete(product Product)
}

type Product struct {
    ProductId     uint
    ProductName   string
    ProductPrice  uint
    BaseCurrency  string
    ProductTitle  string
    ProductDesc  *string
    ImageUrl      string
}


func (*Product)  findBy() (Product, error){
    db,err :=core.GetDB()
	var product Product
	if err != nil {
      return product, err;
	}
    db.First(&product);
    return  product, err;
}


func NewDatabaseManager() Product {
	return Product{};
}

