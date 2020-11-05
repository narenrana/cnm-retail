package repository

import (
	core "shopping-cart/cnm-core"
)

type UsersRepository interface {
	Get() [] users
	FindBy(u users)
	Add(u users)
	Delete(u users)
}

type users struct {
	UserId         uint
	UserName       string
	UserEmail      string
	Password       string
	PhoneNumber    string
	DateCreate     string
	DateUpdated    string

}


func (*users) FindBy() (users, error){
	db,err :=core.GetDB()
	var usr users
	if err != nil {
		return usr, err;
	}
	db.First(&usr);
	return  usr, err;
}


func NewUsersRepository() users {
	return users{};
}

