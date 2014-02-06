package db

import (
	"../models"
)

/**
 * create models tables
 */
func CreateTables() {
	DB.CreateTable(models.User{})
}

/**
 * drop models tables
 */
func DropTables() {
	DB.DropTable(models.User{})
}

/**
 * migrate tables.
 */
func Migrate() {
	DB.AutoMigrate(models.User{})
}
