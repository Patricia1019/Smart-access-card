package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func AutoMigrate(args ...interface{}) {
	err := db.AutoMigrate(args...).Error
	if err != nil {
		panic(err)
	}
}

func InitDB(dialect string, args ...interface{}) {
	var err error
	// db, err = gorm.Open("mysql", "root@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local")
	db, err = gorm.Open(dialect, args...)
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
}
