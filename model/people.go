package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type People struct {
	ID        uint                  `gorm:"primary_key" json:"id"`
	CreatedAt time.Time             `json:"-"`
	UpdatedAt time.Time             `json:"-"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:idx_deleted_at" json:"-"`
	Number    int                   `gorm:"not null;uniqueIndex:idx_deleted_at" json:"number" binding:"required"`
}

func (u *People) TableName() string {
	return "yy_people"
}

func (u *People) Add() error {
	return DB.Create(u).Error
}

func (u *People) Update() error {
	return DB.Model(u).Updates(u).Error
}

func (u *People) Delete() error {
	return DB.Delete(u).Error
}

func (u *People) GetList() ([]People, error) {
	var list []People
	err := DB.Find(&list).Error
	return list, err
}

func (u *People) GetFromIDS(ids []uint) ([]People, error) {
	var people []People
	err := DB.Where("id in (?)", ids).Find(&people).Error
	return people, err
}

func (u *People) GetIDSFrom(peoples []People) ([]uint, error) {
	var peppleids []uint
	for _, people := range peoples {
		peppleids = append(peppleids, people.ID)
	}
	return peppleids, nil
}
