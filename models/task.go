package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	ID           uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key:true"`
	Content      string `json:"content"`
	Agentaccount string `json:"agentaccount"`
	Users        []User `json:"users" gorm:"many2many:TaskUsers;"`
	Undousers    []User `json:"undousers" gorm:"many2many:TaskundoUsers;"`
	Doneusers    []User `json:"doneusers" gorm:"many2many:TaskdoneUsers;"`
	Complete     bool   `json:"complete"`
}

type ContentData struct {
	Title       string `json:"title"`
	Description string `json:"description" `
	Pay         string `json:"pay" `
	StartDate   string `json:"startDate" `
	EndDate     string `json:"endDate" `
	Type        string `json:"type" `
}

type TaskStore interface {
	CreateTask(t Task)
	GetTaskByID(id uint64) (error, Task)
	UpdateTask(t Task) error
	TaskModelInit()
}

//will be executed before main function
func (db *DBStore) TaskModelInit() {
	db.Set("gorm:table_options", "ENGINE=InnoDB AUTO_INCREMENT=1;").AutoMigrate(&Task{})
}

func (db *DBStore) CreateTask(t Task) error {
	db.NewRecord(t)
	db.Create(&t)
	return nil
}

func (db *DBStore) GetTaskByID(id uint64) (error, Task) {
	var t Task
	if db.Preload("Users").Preload("Undousers").Preload("Doneusers").Where("ID = ?", id).First(&t).RecordNotFound() {
		return gorm.ErrRecordNotFound, t
	}
	return nil, t
}

func (db *DBStore) GetAllTask() (error, []Task) {
	var t []Task
	if db.Preload("Users").Preload("Undousers").Preload("Doneusers").Find(&t).RecordNotFound() {
		return gorm.ErrRecordNotFound, t
	}
	return nil, t
}

func (db *DBStore) UpdateTask(t *Task) {
	db.Save(t)
}

func (db *DBStore) DeleteTask(t *Task) {
	db.Delete(t)
}
