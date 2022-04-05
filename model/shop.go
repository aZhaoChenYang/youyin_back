package model

import "github.com/jinzhu/gorm"

type Shop struct {
	gorm.Model
	Name    string  `gorm:"not null;unique"`
	Address string  `gorm:"not null"`
	Mobile  string  `gorm:"not null"`
	Lat     float32 `gorm:"not null"`
	Long    float32 `gorm:"not null"`
}

// TableName 表名
func (u *Shop) TableName() string {
	return "yy_shop"
}
