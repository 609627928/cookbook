package model

import "github.com/jinzhu/gorm"

type Food struct {
	gorm.Model
	Name  string `gorm:"unique_index;not null"`
	Price int64  `gorm:"not null;default:0"`
	Desc  string `gorm:"null"`
}
