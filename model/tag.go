package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"not null" json:"name" binding:"required"`
}

func (u *Tag) TableName() string {
	return "yy_tag"
}

func (u *Tag) Add() error {
	return GetDB().Create(u).Error
}

func (u *Tag) Update() error {
	return GetDB().Save(u).Error
}

func (u *Tag) Delete() error {
	return GetDB().Delete(u).Error
}

func (u *Tag) GetList() (interface{}, error) {
	var list []Tag
	err := GetDB().Find(&list).Error
	return list, err
}
