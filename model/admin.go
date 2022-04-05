package model

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	Username string `gorm:"not null;unique" binding:"required"`
	Password string `gorm:"not null" binding:"required"`
	Nickname string `gorm:"not null"`
}

// TableName 表名
func (u *Admin) TableName() string {
	return "yy_admin"
}
