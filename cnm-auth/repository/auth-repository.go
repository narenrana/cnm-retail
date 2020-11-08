package repository

import (
	core "shopping-cart/cnm-core"
)

type UsersAuthRepository interface {

	FindById(u string) (Users, error)
	AddOrUpdate(u Users)  (Users, error)
	Delete(u Users)  (Users, error)
}

type Users struct {
	UserId         *int        `gorm:"primaryKey" json:"userId,omitempty"`
	FirstName      string     `json:"firstName,omitempty"`
	MiddleName     string    `json:"middleName,omitempty"`
	LastName       string      `json:"lastName,omitempty"`
	UserEmail      string      `json:"userEmail,omitempty"`
	Password       string      `json:"password,omitempty"`
	PhoneNumber    string      `json:"phoneNumber,omitempty"`

}


func (u *Users) FindById(email string) (Users, error) {
	db,err :=core.GetDB()
	var found Users

	db.Model(&found).Where("user_email=?",email).First(&found);
	return  found, err;
}

func (u *Users) AddOrUpdate(user Users) (Users, error) {
	db,err :=core.GetDB()
	if db.Model(&user).Where("user_id = ?", user.UserId).Updates(&user).RowsAffected == 0 {
		db.Create(&user)
	}
	return user, err
}

func (u *Users) Delete(user Users) (Users, error) {
	db,err :=core.GetDB()
    db.Delete(&user)
	return user , err
}

func NewUsersRepository() UsersAuthRepository {
	return &Users{}
}

