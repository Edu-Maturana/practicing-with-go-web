package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var url = "uaeglbv6jfjd8ukb:gCyC6IQnGpjJ8JQnTIJT@tcp(bys4hq6m8jrd9fovr6hz-mysql.services.clever-cloud.com:3306)/bys4hq6m8jrd9fovr6hz"

var Connection = func() *gorm.DB {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
	}
	return db
}
