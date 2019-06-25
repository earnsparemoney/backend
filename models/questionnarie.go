package models

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
	DoneUsers	 []string `json:"doneUsers"`
	//DoneUsers    []User   `json:"doneUsers" gorm:"foreignkey:Account"`
}

type QuesStore interface{
	GetQuesByAgentAccount(agentAccount string) (error, []Questionnaire)
	GetQuesByAttendUser(uid string) (error, []Questionnaire) 
	GetQuesByUserAndAgent(uid string, agentAccount string) (error, []Questionnaire)
	AddQues(q *Questionnaire) error 
	AddUserToQues(qid uint64, uid string) error
	GetAllQues() []Questionnaire

}

func (db *DBStore) GetQuesByID (qid int) (error, Questionnaire){
	var q Questionnaire
	if db.Where("ID=?", qid).First(&q).RecordNotFound(){
		return gorm.ErrRecordNotFound, q
	}

	return nil, q
}


func (db *DBStore)GetQuesByAgentAccount(agentAccount string) (error, []Questionnaire) {
	var qArr []Questionnaire
	if db.Where("AgentAccount = ?", agentAccount).Find(&qArr).RecordNotFound(){
		return gorm.ErrRecordNotFound, qArr
	}

	return nil, qArr
}

func (db *DBStore)GetQuesByAttendUser(uid string) (error, []Questionnaire) {
	var qArr []Questionnaire
	if db.Where("DoneUsers ", ).Find(&qArr).RecordNotFound(){
		return gorm.ErrRecordNotFound, qArr
	}

	return nil, qArr
}

/*
func (db *DBStore)GetQuesByUserAndAgent(uid string, agentAccount string) (error, []Questionnaire) {

}
*/

func (db *DBStore) GetAllQues() []Questionnaire {
	var allQ []Questionnaire
	return db.Find(&allQ)
}

func (db *DBStore)AddQues(q *Questionnaire) error {
	db.Create(q)
	return nil
}

func (db *DBStore)AddUserToQues(q * Questionnaire) error {
	db.Save(q)
	return nil
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
