package repository

import (
	core "shopping-cart/cnm-core"
	e "shopping-cart/cnm-users/entities"
	
)

type UsersRepository interface {
	List() ([]e.UserDetails,error)
	FindBy(u e.UserDetails)  (e.UserDetails, error)
	Add(u e.UserDetails) (e.UserDetails, error)
	Delete(u e.UserDetails) (e.UserDetails, error)
}

type usersRepository struct {
	
}


func (*usersRepository) List() ([] e.UserDetails, error){
	db,err :=core.GetDB()
	var found [] e.UserDetails;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*usersRepository) FindBy(user e.UserDetails) (e.UserDetails, error){
	db,err :=core.GetDB()
	 var found e.UserDetails;
	if err != nil {
		return user, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*usersRepository) Add(user e.UserDetails) (e.UserDetails, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Create(&user);
	return  user, err;
}

func (*usersRepository) Delete(user e.UserDetails) (e.UserDetails, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Delete(&user);
	return  user, err;
}

func UsersRepositoryInstance() UsersRepository {
	return &usersRepository{};
}

