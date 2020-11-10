package repository

import (
	e "shopping-cart/cnm-auth/entities"
	core "shopping-cart/cnm-core"
)

type UsersAuthRepository interface {
	FindById(u string) (e.Users, error)
	AddOrUpdate(u e.Users)  (e.Users, error)
	Delete(u e.Users)  (e.Users, error)
}

type usersAuthRepository struct {
	
}


func (*usersAuthRepository) FindById(email string) (e.Users, error) {
	db,err :=core.GetDB()
	var found e.Users

	db.Model(&found).Where("user_email=?",email).First(&found);
	return  found, err;
}

func (*usersAuthRepository) AddOrUpdate(user e.Users) (e.Users, error) {
	db,err :=core.GetDB()
	if db.Model(&user).Where("user_id = ?", user.UserId).Updates(&user).RowsAffected == 0 {
		db.Create(&user)
	}
	return user, err
}

func (*usersAuthRepository) Delete(user e.Users) (e.Users, error) {
	db,err :=core.GetDB()
    db.Delete(&user)
	return user , err
}

func NewUsersRepository() UsersAuthRepository {
	return &usersAuthRepository{}
}

