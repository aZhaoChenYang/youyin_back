package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname   string `gorm:"not null;size:20" json:"nickname"`
	OpenID     string `gorm:"not null;size:50;unique" json:"openid"`
	SessionKey string `gorm:"not null;size:50" json:"session_key"`
	Avatar     string `gorm:"not null;size:255" json:"avatar"`
	Phone      string `gorm:"not null;size:11" json:"phone"`
	Vip        int    `gorm:"not null;default:0" json:"vip"`
}

func (u *User) TableName() string {
	return "yy_user"
}

func (u *User) Create() error {
	return DB.Create(u).Error
}

func (u *User) IsExist() bool {
	err := DB.Where("open_id = ?", u.OpenID).First(u).Error
	fmt.Println(err)
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func (u *User) Update() error {
	return DB.Model(u).Updates(u).Error
}
