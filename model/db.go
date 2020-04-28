package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

func New() *gorm.DB {
	// 为了处理time.Time，需要包括parseTime作为参数
	db, err := gorm.Open("mysql", "peng:0@tcp(127.0.0.1:3306)/cookbook?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.SingularTable(true) // 全局设置表名不可以为复数形式（表名默认加后缀s）。
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true) // 开启debug，打印原生sql
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
		&User{},
		&Food{},
	)
}
