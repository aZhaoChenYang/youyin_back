package model

import "github.com/jinzhu/gorm"

type People struct {
	gorm.Model
	Number int `gorm:"not null" json:"number" binding:"required"`
}

func (u *People) TableName() string {
	return "yy_people"
}

func (u *People) Add() error {
	return DB.Create(u).Error
}

func (u *People) Update() error {
	return DB.Save(u).Error
}

func (u *People) Delete() error {
	return DB.Delete(u).Error
}

func (u *People) GetList() (interface{}, error) {
	var list []People
	err := DB.Find(&list).Error
	return list, err
}
