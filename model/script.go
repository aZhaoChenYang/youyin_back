package model

import "github.com/jinzhu/gorm"

type Script struct {
	gorm.Model
	Name     string  `gorm:"not null;unique"`
	imgUrl   string  `gorm:"not null"`
	describe string  `gorm:"not null;size:10240"`
	time     int     `gorm:"not null"`
	boys     uint    `gorm:"not null"`
	girls    uint    `gorm:"not null"`
	price1   float32 `gorm:"not null"`
	price2   float32 `gorm:"not null"`
}

// TableName 表名
func (u *Script) TableName() string {
	return "yy_script"
}
