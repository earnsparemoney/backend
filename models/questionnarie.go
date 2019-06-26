package models

import (
	"github.com/jinzhu/gorm"
)

type publish struct {
	qid uint64
	uid string
}

type receive struct {
	qid uint64
	uid string
}

type Questionnaire struct {
	ID           int   `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key:true"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Startdate    string   `json:"startdate"`
	EndDate      string   `json:"enddate"`
	Adward       string   `json:"adward"`
	Body         string   `json:"body"`
	Usernum      int      `json:"usernum"`
	AgentAccount string   `json:"agentAccount"`
	DoneUsers	 []User `json:"doneUsers" `
	//DoneUsers    []User   `json:"doneUsers" gorm:"foreignkey:Account"`
}


type UserQues struct {
	ID int `json:"id" gorm:"primary_key"` //ques id
	Result string `json:"result"`
	UserAccount string `json:"userAccount" gorm:"primary_key"`
}

type QuesStore interface{
	GetQuesByAgentAccount(agentAccount string) (error, []Questionnaire)
	GetQuesByAttendUser(uid string) (error, []Questionnaire) 
	GetQuesByUserAndAgent(uid string, agentAccount string) (error, []Questionnaire)
	AddQues(q *Questionnaire) error 
	AddUserToQues(qid uint64, uid string) error
	GetAllQues() []Questionnaire

}

func (db *DBStore) QuesModelInit() {
	db.Set("gorm:table_options", "ENGINE=InnoDB AUTO_INCREMENT=1;").AutoMigrate(&Questionnaire{})
	db.Set("gorm:table_options", "ENGINE=InnoDB AUTO_INCREMENT=1;").AutoMigrate(&UserQues{})
	
}


// modify to UserQues
func (db *DBStore) GetQuesByID (qid int) (error, Questionnaire){
	var q Questionnaire
	if db.Preload("DoneUsers").Where("ID=?", qid).First(&q).RecordNotFound(){
		return gorm.ErrRecordNotFound, q
	}

	return nil, q
}

// right
func (db *DBStore)GetQuesByAgentAccount(agentAccount string) (error, []Questionnaire) {
	var qArr []Questionnaire
	if db.Preload("DoneUsers").Where("AgentAccount = ?", agentAccount).Find(&qArr).RecordNotFound(){
		return gorm.ErrRecordNotFound, qArr
	}

	return nil, qArr
}

//right

/*
func (db *DBStore)GetQuesByAttendUser(uid string) (error, []Questionnaire) {
	var qArr []Questionnaire
	if db.Where("DoneUsers ", ).Find(&qArr).RecordNotFound(){
		return gorm.ErrRecordNotFound, qArr
	}

	return nil, qArr
}
*/

/*
func (db *DBStore)GetQuesByUserAndAgent(uid string, agentAccount string) (error, []Questionnaire) {

}
*/

//right
func (db *DBStore) GetAllQues() []Questionnaire {
	var allQ []Questionnaire
	db.Find(&allQ)
	return	allQ
}

func (db *DBStore)AddQues(q *Questionnaire) error {
	db.Create(q)
	return nil
}

func (db *DBStore)AddUserToQues(q * Questionnaire) error {
	db.Save(q)
	return nil
}

func (db *DBStore) AddUserResult(userQ * UserQues ) error{
	db.Save(userQ)
	return nil
}

func (db *DBStore) GetUserQuesByQuesID(id int) []UserQues{
	var uqs []UserQues
	db.Preload("DoneUsers").Where("ID=?", id).Find(&uqs)
	return uqs
}



/*
type questModel struct{
	ID           uint64   `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key:true"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Startdate    string   `json:"startdate"`
	EndDate      string   `json:"enddate"`
	Adward       string   `json:"adward"`
	Body         string   `json:"body"`
	Usernum      int      `json:"usernum"`
	Agent		User	`gorm:""`
	DoneUsers	[]User	`gorm:""`


}
*/
