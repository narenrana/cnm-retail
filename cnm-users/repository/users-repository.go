package repository

import (
	core "shopping-cart/cnm-core"
)

type UsersRepository interface {
	List() ([]Users,error)
	FindBy(u Users)  (Users, error)
	Add(u Users) (Users, error)
	Delete(u Users) (Users, error)
}

type Users struct {
	UserId         uint
	FirstName      string
	MiddleName     string
	LastName       string
	UserEmail      string
	Password       string
	PhoneNumber    string
	DateCreate     string
	DateUpdated    string

}

func (*Users) List() ([] Users, error){
	db,err :=core.GetDB()
	var found [] Users;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*Users) FindBy(user Users) (Users, error){
	db,err :=core.GetDB()
	 var found Users;
	if err != nil {
		return user, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*Users) Add(user Users) (Users, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Create(&user);
	return  user, err;
}

func (*Users) Delete(user Users) (Users, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Delete(&user);
	return  user, err;
}

func UsersRepositoryInstance() UsersRepository {
	return &Users{};
}

