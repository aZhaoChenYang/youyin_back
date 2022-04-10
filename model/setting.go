package model

import (
	"github.com/jinzhu/gorm"
)

type Setting struct {
	gorm.Model
	AppKey string `gorm:"unique;nut null" json:"key"`
	Value  string `gorm:"nut null" json:"value"`
	Value1 string `json:"value1"`
	Desc   string `gorm:"nut null" json:"desc"`
	Type   int    `gorm:"nut null"`
	Switch int    `gorm:"nut null"`
}

func (u *Setting) TableName() string {
	return "yy_settings"
}

// GetSettingByType 根据type获取设置
func (u *Setting) GetList(typeid int) ([]Setting, error) {
	var settings []Setting
	err := GetDB().Where("type = ?", typeid).Find(&settings).Error
	return settings, err
}

// 修改设置
func (u *Setting) Update() error {
	return GetDB().Model(u).Updates(u).Error
}

// 根据key获取value，value1
func (u *Setting) GetValue(key string) (string, string, error) {
	err := GetDB().Where("app_key = ?", key).First(u).Error
	return u.Value, u.Value1, err
}

// 根据key获取value
func (u *Setting) GetValueByKey(key string) (string, error) {
	err := GetDB().Where("app_key = ?", key).First(u).Error
	return u.Value, err
}
