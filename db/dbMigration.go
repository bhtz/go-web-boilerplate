package db

import (
	"../app/models"
	"fmt"
)

/**
 * create models tables
 */
func CreateTables() {
	DB.CreateTable(models.User{})
	fmt.Printf("database tables created\n")
}

/**
 * drop models tables
 */
func DropTables() {
	DB.DropTable(models.User{})
	fmt.Printf("database tables dropped\n")
}

/**
 * migrate tables.
 */
func Migrate() {
	DB.AutoMigrate(models.User{})
	fmt.Printf("migration end\n")
}
