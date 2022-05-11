package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Tag struct {
	ID        uint                  `gorm:"primary_key" json:"id"`
	CreatedAt time.Time             `json:"-"`
	UpdatedAt time.Time             `json:"-"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:idx_deleted_at" json:"-"`
	Name      string                `gorm:"not null;size:10;uniqueIndex:idx_deleted_at" json:"name" binding:"required"`
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
