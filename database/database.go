package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	conn, err := gorm.Open(sqlite.Open("database.sqlite"))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
