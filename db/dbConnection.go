package db

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"io/ioutil"
)

var DB gorm.DB

/**
 * Database config structure.
 */
type Config struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Dialect  string `json:"dialect"`
}

/**
 * Initialize database connection.
 */
func Init() {
	file, err := ioutil.ReadFile("./configs/database.json")
	if err != nil {
		panic(fmt.Sprintf("database.json configuration file error: %v\n", err))
	}

	var config Config
	json.Unmarshal(file, &config)

	DB, err = gorm.Open(config.Dialect, config.User+":"+config.Password+"@/"+config.Name)
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}
}
