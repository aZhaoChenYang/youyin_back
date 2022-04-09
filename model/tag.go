package model

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name string `gorm:"not null;size:10" json:"name" binding:"required"`
}

func (u *Tag) TableName() string {
	return "yy_tag"
}

func (u *Tag) Add() error {
	return GetDB().Create(u).Error
}

func (u *Tag) Update() error {
	return GetDB().Model(u).Updates(u).Error
}

func (u *Tag) Delete() error {
	return GetDB().Delete(u).Error
}

func (u *Tag) GetList() (interface{}, error) {
	var list []Tag
	err := GetDB().Find(&list).Error
	return list, err
}

func (u *Tag) GetFromIDS(ids []uint) ([]Tag, error) {
	var tag []Tag
	err := DB.Where("id in (?)", ids).Find(&tag).Error
	return tag, err
}

func (u *Tag) GetIDSFrom(ids []Tag) ([]uint, error) {
	var tag []uint
	for _, id := range ids {
		tag = append(tag, id.ID)
	}
	return tag, nil
}
