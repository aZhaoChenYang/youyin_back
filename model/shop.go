package model

import "github.com/jinzhu/gorm"

type Shop struct {
	gorm.Model
	Name    string   `gorm:"not null;unique" json:"name"`
	Address string   `gorm:"not null" json:"address"`
	Mobile  string   `gorm:"not null" json:"mobile"`
	Lat     float32  `gorm:"not null" json:"lat"`
	Long    float32  `gorm:"not null" json:"long"`
	Swipers []Swiper `gorm:"foreignKey:ShopName;references:Name"`
}

// TableName 表名
func (u *Shop) TableName() string {
	return "yy_shop"
}
