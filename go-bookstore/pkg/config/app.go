package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// here we connect to the database
func Connect() {
	d, err := gorm.Open("mysql", "kullop:Fuckyou123*@/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d
}

// here we return the database to use it in other places of our program
func GetDB() *gorm.DB {
	return db
}
