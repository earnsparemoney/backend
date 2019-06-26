package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	ID           uint64      `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key:true"`
	Content      string      `json:"content"`
	Agentaccount string      `json:"agentaccount"`
	Complete     bool        `json:"complete"`
	Users        []TaskUsers `json:"users" gorm:"ForeignKey:TaskID"`
	//Users []User  `json:"users" gorm:"many2many:TaskUsers;"`
	//Undousers []User  `json:"undousers" gorm:"many2many:TaskundoUsers;"`
	//Doneusers []User  `json:"doneusers" gorm:"many2many:TaskdoneUsers;"`
}

type ContentData struct {
	Title       string `json:"title"`
	Description string `json:"description" `
	Pay         string `json:"pay" `
	StartDate   string `json:"startDate" `
	EndDate     string `json:"endDate" `
	Type        string `json:"type" `
}

type TaskUsers struct {
	TaskID   uint64 `json:"taskid"`
	UserID   string `json:"userid"`
	Finished bool   `json:"finished"`
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
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&TaskUsers{})
}

func (db *DBStore) CreateTask(t Task) error {
	db.NewRecord(t)
	db.Create(&t)
	return nil
}

func (db *DBStore) CreateTaskUsers(t TaskUsers) error {
	db.NewRecord(t)
	db.Create(&t)
	return nil
}

func (db *DBStore) GetTaskByID(id uint64) (error, Task) {
	var t Task
	if db.Preload("Users").Where("ID = ?", id).First(&t).RecordNotFound() {
		return gorm.ErrRecordNotFound, t
	}
	return nil, t
}

func (db *DBStore) GetAllTask() (error, []Task) {
	var t []Task
	if db.Preload("Users").Find(&t).RecordNotFound() {
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
