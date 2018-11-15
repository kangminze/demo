package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sevenNt/wzap"
	"demo/src/model"
)

var (
	db *gorm.DB
)

func CreateTable() {
	if !db.HasTable(&model.User{}) {
		if err :=db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.User{}).Error; err != nil {
			wzap.Err(err)
			panic(err)
		}
	}
}

func Init()  {

	var err error
	db, err = gorm.Open("mysql", "root:123456@(10.20.1.182:3306)/kang_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		wzap.Err(err)
		panic("数据库连接失败")
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)

	wzap.Info("db init success")

	CreateTable()
}