package helper

import (
	"os"

	. "sidu/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	dsn := os.Getenv("DSN_CONNECTION")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error Connecting to Database")
	}
	db.AutoMigrate(&User{}, &Document{})

	return db
}
