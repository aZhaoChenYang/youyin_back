package model

import (
	"gorm.io/gorm"
	"time"
)

type Script struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string   `gorm:"not null;unique;size:255" json:"title" binding:"required"`
	ImgUrl    string   `gorm:"not null;size:255" json:"imgurl" binding:"required"`
	Describe  string   `gorm:"not null;size:10240" json:"describe" binding:"required"`
	Time      int      `gorm:"not null" json:"time" binding:"required"`
	Boys      uint     `gorm:"not null" json:"boys" binding:"required"`
	Girls     uint     `gorm:"not null" json:"girls" binding:"required"`
	Price1    float32  `gorm:"not null" json:"price1" binding:"required"`
	Price2    float32  `gorm:"not null" json:"price2" binding:"required"`
	ShopID    uint     `gorm:"not null" json:"shop_id" binding:"required"`
	Peoples   []People `gorm:"many2many:script_people"`
	Tags      []Tag    `gorm:"many2many:script_tag"`
}

type Jsonscript struct {
	ID       uint    `json:"id"`
	Name     string  `json:"title" binding:"required"`
	ImgUrl   string  `json:"imgurl" binding:"required"`
	Describe string  `json:"describe" binding:"required"`
	Time     int     `json:"time" binding:"required"`
	Boys     uint    `json:"boys"`
	Girls    uint    `json:"girls"`
	Price1   float32 `json:"price1" binding:"required"`
	Price2   float32 `json:"price2" binding:"required"`
	ShopID   uint    `json:"shop_id" binding:"required"`
	Peoples  []uint  `json:"peoples" binding:"required"`
	Tags     []uint  `json:"tags" binding:"required"`
}

type JsonBasescript struct {
	ID       uint   `json:"id"`
	Name     string `json:"title"`
	ImgUrl   string `json:"imgurl"`
	Time     int    `json:"time"`
	Boys     uint   `json:"boys"`
	Girls    uint   `json:"girls"`
	ShopName string `json:"shop_name"`
}

// TableName 表名
func (u *Script) TableName() string {
	return "yy_script"
}

func (u *Script) Create() error {
	return DB.Create(u).Error
}

func (u *Script) GetList() ([]JsonBasescript, error) {
	var jsonscripts []JsonBasescript
	err := GetDB().Table("yy_script").Select("yy_script.id,yy_script.name,yy_script.img_url,yy_script.time,yy_script.boys,yy_script.girls,yy_shop.name as shop_name").Joins("left join yy_shop on yy_script.shop_id = yy_shop.id").Scan(&jsonscripts).Error
	return jsonscripts, err
}

func (u *Script) Get(id uint) (interface{}, error) {
	var script Script
	err := GetDB().Find(&script, id).Error
	if err != nil {
		return nil, err
	}
	err = GetDB().Model(&script).Association("Tags").Find(&script.Tags)
	if err != nil {
		return nil, err
	}
	tags, err := (&Tag{}).GetIDSFrom(script.Tags)
	if err != nil {
		return nil, err
	}
	err = GetDB().Model(&script).Association("Peoples").Find(&script.Peoples)
	if err != nil {
		return nil, err
	}
	peoples, err := (&People{}).GetIDSFrom(script.Peoples)
	if err != nil {
		return nil, err
	}
	jsonscript := Jsonscript{
		ID:       script.ID,
		Name:     script.Name,
		ImgUrl:   script.ImgUrl,
		Describe: script.Describe,
		Time:     script.Time,
		Boys:     script.Boys,
		Girls:    script.Girls,
		Price1:   script.Price1,
		Price2:   script.Price2,
		ShopID:   script.ShopID,
		Peoples:  peoples,
		Tags:     tags,
	}

	return jsonscript, err
}

func (u *Script) Delete() error {
	return GetDB().Transaction(func(tx *gorm.DB) error {

		err := tx.Model(&u).Association("Tags").Clear()
		if err != nil {
			return err
		}
		err = tx.Model(&u).Association("Peoples").Clear()
		if err != nil {
			return err
		}
		err = tx.Delete(u).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (u *Script) Update() error {
	return GetDB().Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&u).Association("Tags").Replace(u.Tags)
		if err != nil {
			return err
		}
		err = tx.Model(&u).Association("Peoples").Replace(u.Peoples)
		if err != nil {
			return err
		}
		err = tx.Model(&u).Updates(u).Error
		if err != nil {
			return err
		}
		return nil
	})
}
