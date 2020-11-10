package repository

import (
	core "shopping-cart/cnm-core"
	 e "shopping-cart/cnm-products/entities"
)

type Repository interface {
	List() ([] e.Product,error)
	FindByIds(ids [] int) ([] e.Product,error)
	FindBy(u e.Product) (  e.Product,error)
	Add(u e.Product) ( e.Product,error)
	Delete(u e.Product) ( e.Product,error)
}

type repository struct {

}


func (p *repository) FindByIds(ids []int) ([]e.Product, error) {
	db,err :=core.GetDB()
	var found [] e.Product;
	if err != nil {
		return nil, err;
	}
	db.Model(&found).Where("product_id in (?)",ids).Find(&found);
	return  found, err;
}

func (p *repository) List() ([] e.Product, error){
	db,err :=core.GetDB()
	var found [] e.Product;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	return  found, err;
}

func (p *repository) FindBy(user e.Product) (e.Product, error){
	db,err :=core.GetDB()
	 var found e.Product;
	if err != nil {
		return user, err;
	}
	db.Find(&found);
	return  found, err;
}

func (p *repository) Add(user e.Product) (e.Product, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Create(&user);
	return  user, err;
}

func (p *repository) Delete(user e.Product) (e.Product, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Delete(&user);
	return  user, err;
}

func ProductRepositoryInstance() Repository {
	return &repository{};
}

