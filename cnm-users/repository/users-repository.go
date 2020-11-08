package repository

import (
	core "shopping-cart/cnm-core"
)

type UsersRepository interface {
	List() ([]UserDetails,error)
	FindBy(u UserDetails)  (UserDetails, error)
	Add(u UserDetails) (UserDetails, error)
	Delete(u UserDetails) (UserDetails, error)
}

type UserDetails struct {
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

func (*UserDetails) List() ([] UserDetails, error){
	db,err :=core.GetDB()
	var found [] UserDetails;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*UserDetails) FindBy(user UserDetails) (UserDetails, error){
	db,err :=core.GetDB()
	 var found UserDetails;
	if err != nil {
		return user, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*UserDetails) Add(user UserDetails) (UserDetails, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Create(&user);
	return  user, err;
}

func (*UserDetails) Delete(user UserDetails) (UserDetails, error){
	db,err :=core.GetDB()
	if err != nil {
		return user, err;
	}
	db.Delete(&user);
	return  user, err;
}

func UsersRepositoryInstance() UsersRepository {
	return &UserDetails{};
}

