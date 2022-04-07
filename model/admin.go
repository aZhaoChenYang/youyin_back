package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	gorm.Model
	Username string `gorm:"not null;unique;size:20" json:"username" binding:"required"`
	Password string `gorm:"not null;size:128" json:"password" binding:"required"`
	Nickname string `gorm:"not null;size:20" json:"nickname"`
}

// TableName 表名
func (u *Admin) TableName() string {
	return "yy_admin"
}

// Create 添加一条记录
func (u *Admin) Create() error {
	return GetDB().Create(u).Error
}

// GetAll 获取全部信息
func (u *Admin) GetAll() ([]Admin, error) {
	var admins []Admin
	err := GetDB().Select("id, username, nickname").Find(&admins).Error
	return admins, err
}

// Update 更新一条记录
func (u *Admin) Update() error {
	return GetDB().Model(u).Updates(u).Error
}

// Delete 删除一条记录
func (u *Admin) Delete() error {
	return GetDB().Delete(u).Error
}

func (u *Admin) Login() error {
	password := u.Password
	if err := GetDB().Where("username = ?", u.Username).First(u).Error; err != nil {
		return err
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}
