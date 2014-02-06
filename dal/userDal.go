package dal

import (
	"../db"
	"../models"
)

/**
 * UserDal structure.
 */
type UserDal struct {
}

/**
 * UserDal constructor.
 */
func NewUserDal() UserDal {
	u := UserDal{}
	return u
}

/**
 * Get an instance of models.User for test
 */
func (u *UserDal) GetUser() models.User {
	user := models.User{
		Id:        1,
		Email:     "heintz.benjamin@gmail.com",
		Firstname: "benjamin",
		Lastname:  "heintz",
		Github:    "https://github.com/bhtz",
	}

	return user
}

/**
 * get user by id.
 */
func (u *UserDal) Get(id int64) models.User {
	var user models.User
	db.DB.First(&user, id)
	return user
}

/**
 * get all users.
 */
func (u *UserDal) GetAll() []models.User {
	var users []models.User
	db.DB.Find(&users)
	return users
}
