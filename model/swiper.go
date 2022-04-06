package model

import "github.com/jinzhu/gorm"

type Swiper struct {
	gorm.Model
	Imgurl   string `gorm:"not null" json:"imgurl"`
	ShopName string `gorm:"not null" json:"shopName"`
}

// TableName 表名
func (u *Swiper) TableName() string {
	return "yy_swiper"
}

func (u *Swiper) Add() error {
	return DB.Create(u).Error
}
func (u *Swiper) Update() error {
	return DB.Save(u).Error
}
func (u *Swiper) Delete() error {
	return DB.Delete(u).Error
}
func (u *Swiper) GetList() ([]Swiper, error) {
	var swiper []Swiper
	err := DB.Find(&swiper).Error
	return swiper, err
}
