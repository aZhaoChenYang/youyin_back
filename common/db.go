package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"youyin/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := Conf.MYSQL.DriverName
	host := Conf.MYSQL.Host
	port := Conf.MYSQL.Port
	username := Conf.MYSQL.Username
	password := Conf.MYSQL.Password
	database := Conf.MYSQL.Database
	charset := Conf.MYSQL.Charset
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	db.LogMode(true)
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Shop{})
	db.AutoMigrate(&model.Tag{})
	db.AutoMigrate(&model.People{})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
