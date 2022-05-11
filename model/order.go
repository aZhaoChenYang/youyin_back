package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Order struct {
	ID        uint                  `gorm:"primary_key" json:"id"`
	CreatedAt time.Time             `json:"-"`
	UpdatedAt time.Time             `json:"-"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:idx_deleted_at" json:"-"`
	ScriptId  uint                  `gorm:"foreignkey:ScriptId;uniqueIndex:idx_deleted_at" json:"script_id"`
	DateTime  time.Time             `gorm:"type:datetime;uniqueIndex:idx_deleted_at" json:"dateTime"`
	ShopId    uint                  `gorm:"foreignkey:ShopId;uniqueIndex:idx_deleted_at" json:"shop_id"`
	Status    uint                  `gorm:"type:varchar(255)" json:"status"`
	Users     []User                `gorm:"many2many:order_users"`
}

func (u *Order) TableName() string {
	return "yy_orders"
}

func (u *Order) Create() error {
	return DB.Create(u).Error
}

func (u *Order) GetList() ([]Order, error) {
	var orderList []Order
	err := DB.Find(&orderList).Error
	return orderList, err
}
