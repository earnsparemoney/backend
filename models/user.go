package models

import (
	"github.com/jinzhu/gorm"
)

type  User struct {
	Account string
	Name string
	Password string
	Email string
	Phone string
	Money float32
}


func CreateUser(u User) error{
	db.NewRecord(u)
	db.Create(&u)
	return nil
}

func GetUserByID(id string) 


