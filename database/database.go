package database

import (
	"fmt"

	"prueba/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var url = utils.GetEnv("DB_URL")
var Connection = func() *gorm.DB {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
	}
	return db
}
