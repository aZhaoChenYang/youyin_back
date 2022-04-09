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
