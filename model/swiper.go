package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Swiper struct {
	ID        uint                  `gorm:"primary_key" json:"id"`
	CreatedAt time.Time             `json:"-"`
	UpdatedAt time.Time             `json:"-"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:idx_deleted_at" json:"-"`
	Imgurl    string                `gorm:"size:128;not null;uniqueIndex:idx_deleted_at" json:"imgurl" binding:"required"`
	ShopID    uint                  `gorm:"not null;uniqueIndex:idx_deleted_at" json:"shop_id" binding:"required"`
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
	var list []Swiper
	err := DB.Find(&list).Error
	return list, err
}
