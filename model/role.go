package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Role struct {
	ID        uint                  `gorm:"primary_key" json:"id"`
	CreatedAt time.Time             `json:"-"`
	UpdatedAt time.Time             `json:"-"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:idx_deleted_at" json:"-"`
	Name      string                `gorm:"not null;size:10;uniqueIndex:idx_deleted_at" json:"name" binding:"required"`
}

func (u *Role) TableName() string {
	return "yy_role"
}

func (u *Role) Create() error {
	return DB.Create(u).Error
}
