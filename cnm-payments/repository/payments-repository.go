package payments

import (
	core "shopping-cart/cnm-core"
	e "shopping-cart/cnm-payments/entities"
)

type Repository interface {
	List() ([] e.Payments,error)
	FindBy(u e.Payments) (  e.Payments,error)
	Add(u e.Payments) ( e.Payments,error)
	Delete(u e.Payments) ( e.Payments,error)
}

type repository struct {
}


func (repository) List() ([] e.Payments, error){
	db,err :=core.GetDB()
	var found [] e.Payments;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	return  found, err;
}

func (repository) FindBy(item e.Payments) (e.Payments, error){
	db,err :=core.GetDB()
	 var found e.Payments;
	if err != nil {
		return item, err;
	}
	db.Find(&found);
	return  found, err;
}

func (repository) Add(item e.Payments) (e.Payments, error){
	db,err :=core.GetDB()
	if err != nil {
		return item, err;
	}
	db.Create(&item);
	return  item, err;
}

func (repository) Delete(item e.Payments) (e.Payments, error){
	db,err :=core.GetDB()
	if err != nil {
		return item, err;
	}
	db.Delete(&item);
	return  item, err;
}

func PaymentsRepositoryInstance() Repository {
	return &repository{};
}

