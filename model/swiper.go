package model

import "github.com/jinzhu/gorm"

type Swiper struct {
	gorm.Model
	Imgurl   string `gorm:"not null"`
	ShopName string `gorm:"not null"`
}

// TableName 表名
func (u *Swiper) TableName() string {
	return "yy_swiper"
}
