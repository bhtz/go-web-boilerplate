package db

import (
	"../app/models"
	"fmt"
)

/**
 * Insert database fixture here.
 */
func RunFixtures() {

	bhtz := models.User{
		Email:     "heintz.benjamin@gmail.com",
		Firstname: "Benjamin",
		Lastname:  "Heintz",
		Github:    "www.github.com/bhtz",
	}

	foo := models.User{
		Email:     "foo.bar@gmail.com",
		Firstname: "Foo",
		Lastname:  "Bar",
		Github:    "www.github.com/foobar",
	}

	DB.Save(&bhtz)
	DB.Save(&foo)

	fmt.Printf("Fixtures saved !\n")
}
