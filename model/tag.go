package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func (u *Tag) TableName() string {
	return "yy_tag"
}
