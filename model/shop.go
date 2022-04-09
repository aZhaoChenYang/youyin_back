package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Shop struct {
	gorm.Model
	Name    string   `gorm:"not null;unique;size:20" json:"name" binding:"required"`
	Address string   `gorm:"not null;size:255" json:"address" binding:"required"`
	Mobile  string   `gorm:"not null;size:11" json:"mobile" binding:"required"`
	Lat     float32  `gorm:"not null" json:"lat" binding:"required"`
	Long    float32  `gorm:"not null" json:"long" binding:"required"`
	Swipers []Swiper `json:"-"`
	Scripts []Script `json:"-"`
}

var ForeignkeyError = errors.New("该店铺下还有轮播图，不能删除")

// TableName 表名
func (u *Shop) TableName() string {
	return "yy_shop"
}

func (u *Shop) Create() error {
	return GetDB().Create(u).Error
}

func (u *Shop) Update() error {
	return GetDB().Model(u).Updates(u).Error
}

func (u *Shop) Delete() error {
	err := GetDB().First(u).Error
	if err != nil {
		return err
	}
	count := GetDB().Model(u).Association("Swipers").Count()
	if count > 0 {
		return ForeignkeyError
	}
	return GetDB().Delete(u).Error
}

func (u *Shop) GetList() ([]Shop, error) {
	var shops []Shop
	err := GetDB().Find(&shops).Error
	return shops, err
}
