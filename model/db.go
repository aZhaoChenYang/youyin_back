package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"youyin/common"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := common.Conf.MYSQL.DriverName
	host := common.Conf.MYSQL.Host
	port := common.Conf.MYSQL.Port
	username := common.Conf.MYSQL.Username
	password := common.Conf.MYSQL.Password
	database := common.Conf.MYSQL.Database
	charset := common.Conf.MYSQL.Charset
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	db.LogMode(true)
	db.AutoMigrate(&Setting{})
	db.AutoMigrate(&Admin{})
	db.AutoMigrate(&Shop{})
	db.AutoMigrate(&Swiper{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&People{})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
