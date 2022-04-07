package model

import "github.com/jinzhu/gorm"

type Swiper struct {
	gorm.Model
	Imgurl   string `gorm:"size:128;not null" json:"imgurl" binding:"required"`
	ShopName string `gorm:"size:20;not null" json:"shop_name" binding:"required"`
}

// TableName 表名
func (u *Swiper) TableName() string {
	return "yy_swiper"
}

func (u *Swiper) Add() error {
	return DB.Create(u).Error
}
func (u *Swiper) Update() error {
	return DB.Model(u).Updates(u).Error
}
func (u *Swiper) Delete() error {
	return DB.Delete(u).Error
}
func (u *Swiper) GetList() ([]Swiper, error) {
	var swiper []Swiper
	err := DB.Find(&swiper).Error
	return swiper, err
}
