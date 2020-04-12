package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func DbConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "/test.db")

	if err != nil {
		panic("Unable to connect to db")
	}

	db.AutoMigrate(&Book{})

	return db
}
