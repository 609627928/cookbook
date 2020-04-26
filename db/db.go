package db

import (
	"cookbook/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

func New() *gorm.DB {
	db, err := gorm.Open("mysql", "peng:0@tcp(127.0.0.1:3306)/cookbook?charset=utf8")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func TestDB() *gorm.DB {
	db, err := gorm.Open("mysql", "./../realworld_test.db")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	if err := os.Remove("./../realworld_test.db"); err != nil {
		return err
	}
	return nil
}

//TODO: err check
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}
