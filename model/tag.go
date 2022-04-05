package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
}

func (u *Tag) TableName() string {
	return "yy_tag"
}
