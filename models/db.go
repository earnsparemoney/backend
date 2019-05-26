package models

import (
	"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
)


type DBStore struct{
	*gorm.DB //inherite the method of gorm.DB
} 



//"root:sactestdatabase@tcp(222.200.190.31:33336)/earnsparemoney?charset=utf8&parseTime=True&loc=Local"
func NewDB(dbConfig string) (*DBStore, error){
	var err error
	db, err := gorm.Open("mysql",dbConfig)
	if err != nil{
		panic(err)
	} 
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return &DBStore{db},err
}



