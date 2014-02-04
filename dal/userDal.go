package dal

import (
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
 * Get an instance of models.User
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
