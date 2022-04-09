package model

import (
	"github.com/jinzhu/gorm"
)

type Swiper struct {
	gorm.Model
	Imgurl string `gorm:"size:128;not null" json:"imgurl" binding:"required"`
	ShopID uint   `gorm:"not null" json:"shop_id" binding:"required"`
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
func (u *Swiper) GetList() (interface{}, error) {
	type typeswiper struct {
		ID       uint   `json:"ID"`
		Imgurl   string `json:"imgurl"`
		ShopName string `json:"shop_name"`
	}
	var swiper []typeswiper
	err := GetDB().Table("yy_swiper").Select("yy_swiper.id,yy_swiper.imgurl,yy_shop.name as shop_name").Joins("left join yy_shop on yy_swiper.shop_id = yy_shop.id").Scan(&swiper).Error
	return swiper, err
}
