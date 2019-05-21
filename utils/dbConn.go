package utils

import (
	"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db*gorm.DB

func InitDBConn() {
	var err error
	db, err = gorm.Open("mysql","root:sactestdatabase/test1?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	} 
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func GetDBConn() *gorm.DB{
	return db
}