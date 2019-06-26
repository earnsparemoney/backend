package controllers

import (
	echo "github.com/labstack/echo/v4"
	"github.com/earnsparemoney/backend/utils"
	"github.com/earnsparemoney/backend/models"
	//"github.com/gorilla/sessions"
	//"github.com/labstack/echo-contrib/session"
	//"fmt"
	"github.com/wxnacy/wgo/arrays"
	"strconv"
)

func (uc *UserController) GetQuesByID(c echo.Context) error{
	qid := c.Param("qid")
	qidToint, err := strconv.Atoi(qid)
	var ques models.Questionnaire
	if err, ques = uc.db.GetQuesByID(qidToint); err !=nil{
		return c.JSON(utils.Fail("not such questionnaire"))
	}
	return c.JSON(utils.Success("success", ques))
}


func (uc *UserController) AddQues(c echo.Context) error{
	qu := new(models.Questionnaire)
	if err := c.Bind(qu); err != nil{
		return err
	}

	uc.db.AddQues(qu)
	return c.JSON(utils.Success("success"))

}

//判断uid 和qid是不是为空， 4种情况，4中model处理方式放回
func (uc *UserController)  GetQues(c echo.Context) error{
	qid := c.QueryParam("agentID")
	uid := c.QueryParam("userID")

	var err error 
	var result []models.Questionnaire

	if qid == "" && uid == "" {
		return c.JSON(utils.Success("success", uc.db.GetAllQues()))
	} else if qid != "" && uid == ""{
		var qArr []models.Questionnaire
		
		if err, qArr = uc.db.GetQuesByAgentAccount(qid); err !=nil{
			return c.JSON(utils.Fail("question not exist"))
		}
		return c.JSON(utils.Success("success", qArr))
	} else if qid == "" && uid != ""{
		qArr := uc.db.GetAllQues()
		var user models.User
		var err error
		if err, user = uc.db.GetUserByID(uid); err !=nil{
			return c.JSON(utils.Fail("not such user"))
		}

		for _, q :=range(qArr){

			index := arrays.Contains(q.DoneUsers, user)
			if index !=-1{
				result = append(result, q)
				continue
			}
		}

		/*index := arrays.Contains(qArr.DoneUsers, user.Account)
		if index == -1 {
			return c.JSON(utils.Fail("not exist such questionnaire"))
		}
		*/
		if len(result) ==0{
			return c.JSON(utils.Fail("not such questionnaire exist"))
		}
		return c.JSON(utils.Success("success", result))
		
	}

	return c.JSON(utils.Fail("no support for qid and uid simultaneously"))


}

//得到用户id，如果用户已经在donuser中，则返回失败，否则将用户加进数组，返回成功
// fix: need to add UserQues to db 
func (uc *UserController) AddUserToQues(c echo.Context) error{

	userQ := new(models.UserQues)
	
	qArr := uc.db.GetAllQues()
	var user models.User
	var err error
	if err = c.Bind(userQ); err!=nil{
		return c.JSON(utils.Fail("body info is wrong"))
	}


	uid  := userQ.UserAccount
	qid := userQ.ID

	if err, user = uc.db.GetUserByID(uid); err !=nil{
		return c.JSON(utils.Fail("not such user"))
	}

	for _, q :=range(qArr){
		if q.ID == qid{
			index := arrays.Contains(q.DoneUsers, user)
			if index != -1 {
				return c.JSON(utils.Fail("user already complete the questionnaire"))	//这个逻辑是不是有问题啊，卧槽，凭什么不给别人多填
			} 
			q.DoneUsers = append(q.DoneUsers,user)
			uc.db.AddUserToQues(&q)
			uc.db.AddUserResult(userQ)
			return c.JSON(utils.Success("success"))
		}

	}

	return c.JSON(utils.Fail("not found questionnaire, try to check your qid"))


}

	/*
	index := arrays.Contains(qArr.DoneUsers, user.Account)
	if index != -1 {

		return c.JSON(utils.Fail("user already complete the questionnaire"))
	}
	*/

	//qArr.DoneUsers.append(user.Account)

	//uc.db.AddUserToQues(&qArr[index])
	//return c.JSON(utils.Success("success"))

//get all UserQues with same ques id

func (uc *UserController) GetUserQuesByID(c echo.Context) error{
	qid := c.Param("qid")
	qidToInt, err := strconv.Atoi(qid)

	if err != nil{
		return c.JSON(utils.Fail("error:should be number "))
	}

	return c.JSON(utils.Success("success",uc.db.GetUserQuesByQuesID(qidToInt)))
}
