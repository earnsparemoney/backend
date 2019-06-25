package controllers

import (
	echo "github.com/labstack/echo/v4"
	"github.com/earnsparemoney/backend/utils"
	"github.com/earnsparemoney/backend/models"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"fmt"
	"github.com/wxnacy/wgo/arrays"
	"strconv"
)

func (uc *UserController) GetQuesByID(c echo.Context) error{
	qid := c.Param("qid")
	qidToint := strconv.Atoi(qid)
	if err, ques := uc.db.GetQuesByID(qidToint); err !=nil{
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

	if qid == "" && uid == "" {
		return c.JSON(utils.Success("success", uc.db.GetAllQues()))
	}
	else if qid != "" && uid == ""{
		if err, qArr := uc.db.GetQuesByAgentAccount(qid); err !=nil{
			return c.JSON(utils.Fail("question not exist"))
		}
		return c.JSON(utils.Success("success", qArr))
	}
	else if qid == "" && uid != ""{
		qArr := uc.db.GetAllQues()

		if err, user := uc.db.GetUserByID(uid); err !=nil{
			return c.JSON(utils.Fail("not such user"))
		}
		index := arrays.Contains(qArr.DoneUsers, user.Account)
		if index == -1 {
			return c.JSON(utils.Fail("not exist such questionnaire"))
		}
		return c.JSON(utils.Success("success", qArr[index]))
		
	}

	return c.JSON(utils.Fail("no support for qid and uid simultaneously"))


}

//得到用户id，如果用户已经在donuser中，则返回失败，否则将用户加进数组，返回成功
func (uc *UserController) AddUserToQues(c echo.Context) error{
	qid := c.Param("qid")
	uid := c.Param("uid")

	qArr := uc.db.GetAllQues()

	if err, user := uc.db.GetUserByID(uid); err !=nil{
		return c.JSON(utils.Fail("not such user"))
	}

	index := arrays.Contains(qArr.DoneUsers, user.Account)
	if index != -1 {

		return c.JSON(utils.Fail("user already complete the questionnaire"))
	}

	qArr.DoneUsers.append(user.Account)

	uc.db.AddUserToQues(&qArr[index])
	return c.JSON(utils.Success("success"))
}


