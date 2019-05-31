package models

import (
	"github.com/jinzhu/gorm"
	//"github.com/earnsparemoney/backend/utils"
	
)

type  User struct {
	Account string `json:"account" query:"account" gorm:"primary_key"`
	Name string	`json:"name" form:"name"`
	Password string `json:"password" `
	Email string	`json:"email"`
	Phone string	`json:"phone"`
	Balance float32	`json:"balance"`
}


type UserStore interface {
	CreateUser(u User) 
	GetUserByID(id string) (error, User)
	UpdateUser(u User) error
	UserModelInit()
}



//will be executed before main function
func (db *DBStore)UserModelInit() {

	db.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&User{})

}

func (db *DBStore)CreateUser(u User) error{
	db.NewRecord(u)
	db.Create(&u)
	return nil
}

func (db *DBStore)GetUserByID(id string) (error,User){
	var u User
	if db.Where("Account = ?",id).First(&u).RecordNotFound(){
		return gorm.ErrRecordNotFound,u
	}
	return nil,u
}

func (db *DBStore)UpdateUser(u *User) {
	db.Save(u)
}



