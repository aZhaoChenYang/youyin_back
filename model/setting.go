package model

import (
	"github.com/jinzhu/gorm"
)

type Setting struct {
	gorm.Model
	Key    string `gorm:"unique;nut null"`
	Value  string `gorm:"nut null"`
	Value1 string
	Desc   string `gorm:"nut null"`
	Type   int    `gorm:"nut null"`
	Switch int    `gorm:"nut null"`
}

func (u *Setting) TableName() string {
	return "yy_setting"
}
