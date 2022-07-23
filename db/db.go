package db

import (
	"fmt"

	"domain/train"

	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "./muscle.db")
	if err != nil {
		fmt.Println("strorage err: ", err)
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&train.Train)
}
