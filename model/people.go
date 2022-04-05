package model

import "github.com/jinzhu/gorm"

type People struct {
	gorm.Model
	Number int `gorm:"not null"`
}

func (u *People) TableName() string {
	return "yy_people"
}
