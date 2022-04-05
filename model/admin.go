package model

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	Username string `gorm:"not null;unique" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Nickname string `gorm:"not null" json:"nickname"`
}

// TableName 表名
func (u *Admin) TableName() string {
	return "yy_admin"
}
