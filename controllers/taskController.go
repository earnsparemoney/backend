package controllers

import (
	echo "github.com/labstack/echo/v4"
	//  "net/http"
	"github.com/earnsparemoney/backend/models"
	"github.com/earnsparemoney/backend/utils"
	//  "github.com/gorilla/sessions"
	//"fmt"
	"github.com/earnsparemoney/backend/middleware"
	"github.com/labstack/echo-contrib/session"
	"github.com/wxnacy/wgo/arrays"
	"io/ioutil"
	"strconv"
	//"time"
	//jwt "github.com/dgrijalva/jwt-go"
)

func (uc *UserController) PublishTask(c echo.Context) error {
	token := c.Request().Header.Get(echo.HeaderAuthorization)
	claims, _ := middleware.ParseToken(token, "secret")
	uid := claims["account"].(string)
	t := new(models.Task)
	s, _ := ioutil.ReadAll(c.Request().Body)
	t.Content = string(s)
	t.Agentaccount = uid
	t.Complete = false
	uc.db.CreateTask(*t)
	_, ts := uc.db.GetAllTask()
	return c.JSON(utils.Success("success", ts[len(ts)-1]))
}
func (uc *UserController) GetListTsak(c echo.Context) error {
	//token:=c.Get("claims").(map[string]interface{})
	//uid := token["account"].(string)
	AgentID := c.QueryParam("agentID")
	UserID := c.QueryParam("userID")
	HasAgentID := false
	HasUserID := false
	var ts []models.Task
	var u models.User
	var errU, errA error
	if AgentID != "" {
		HasAgentID = true
		errA, ts = uc.db.GetAllTask()
		if errA != nil {
			return errA
		}
	}
	if UserID != "" {
		HasUserID = true
		errU, u = uc.db.GetUserByID(UserID)
		if errU != nil {
			return errU
		}
	}
	var result []models.Task

	for _, t := range ts {
		if HasUserID {
			index := arrays.Contains(t.Users, u)
			if index != -1 {
				result = append(result, t)
				continue
			}
		}
		if HasAgentID {
			if t.Agentaccount == AgentID {
				result = append(result, t)
				continue
			}
		}
	}
	return c.JSON(utils.Success("OK", result))
}
func (uc *UserController) CompleteTask(c echo.Context) error {
	//Id,_ := strconv.Atoi(c.QueryParam("id"))
	Id, _ := strconv.Atoi(c.Param("taskID"))
	err, t := uc.db.GetTaskByID(uint64(Id))
	if err != nil {
		return err
	}
	t.Complete = true
	uc.db.UpdateTask(&t)
	return c.JSON(utils.Success("update successfully"))
}

func (uc *UserController) DoneuserTask(c echo.Context) error {
	Id, _ := strconv.Atoi(c.QueryParam("id"))
	account := c.QueryParam("account")
	err, t := uc.db.GetTaskByID(uint64(Id))
	err1, u := uc.db.GetUserByID(account)
	if err != nil || err1 != nil {
		return err
	}
	index := arrays.Contains(t.Undousers, u)
	if index == -1 {
		return c.JSON(utils.Fail("user error"))
	}
	uc.db.Model(&t).Association("Undousers").Delete(&u)
	uc.db.Model(&t).Association("Doneusers").Append(&u)
	//	t.Undousers=append(t.Undousers[:index], t.Undousers[index+1:]...)
	//	t.Doneusers=append(t.Doneusers,u)
	uc.db.UpdateTask(&t)
	return c.JSON(utils.Success("update successfully"))
}

func (uc *UserController) DeleteTask(c echo.Context) error {
	//Id,_ := strconv.Atoi(c.QueryParam("id"))
	Id, _ := strconv.Atoi(c.Param("taskID"))
	err, t := uc.db.GetTaskByID(uint64(Id))
	if err != nil {
		return err
	}
	uc.db.DeleteTask(&t)
	return c.JSON(utils.Success("delete successfully"))
}

func (uc *UserController) ParticipateTask(c echo.Context) error {
	//token:=c.Get("claims").(map[string]interface{})
	//uid := token["account"].(string)
	uid := c.Param("account")
	Id, _ := strconv.Atoi(c.Param("taskID"))
	err, t := uc.db.GetTaskByID(uint64(Id))
	err1, u := uc.db.GetUserByID(uid)
	if err != nil || err1 != nil {
		return err
	}
	//	if(uc.db.Model(&t).Association("Users").Count()>=t.Usernum){return c.JSON(utils.Fail("User num enough"))}
	index := arrays.Contains(t.Users, u)
	if index != -1 {
		return c.JSON(utils.Fail("User have join"))
	}
	uc.db.Model(&t).Association("Users").Append(&u)
	uc.db.Model(&t).Association("Undousers").Append(&u)
	//	t.Users=append(t.Users,u)
	//	t.Undousers=append(t.Undousers,u)
	uc.db.UpdateTask(&t)
	return c.JSON(utils.Success("delete successfully"))
}

func (uc *UserController) CancelTask(c echo.Context) error {
	//token:=c.Get("claims").(map[string]interface{})
	//uid := token["account"].(string)
	uid := c.Param("account")
	Id, _ := strconv.Atoi(c.Param("taskID"))
	err, t := uc.db.GetTaskByID(uint64(Id))
	err1, u := uc.db.GetUserByID(uid)
	if err != nil || err1 != nil {
		return err
	}
	index := arrays.Contains(t.Users, u)
	if index == -1 {
		return c.JSON(utils.Fail("User not in task"))
	}
	uc.db.Model(&t).Association("Users").Delete(&u)
	index = arrays.Contains(t.Undousers, u)
	if index != -1 {
		uc.db.Model(&t).Association("Undousers").Delete(&u)
	}
	index = arrays.Contains(t.Doneusers, u)
	if index != -1 {
		uc.db.Model(&t).Association("Doneusers").Delete(&u)
	}
	uc.db.UpdateTask(&t)
	return c.JSON(utils.Success("delete successfully"))
}

func (uc *UserController) GetuserListTask(c echo.Context) error {
	token := c.Request().Header.Get(echo.HeaderAuthorization)
	claims, _ := middleware.ParseToken(token, "secret")
	uid := claims["account"].(string)
	err, ts := uc.db.GetAllTask()
	err1, u := uc.db.GetUserByID(uid)
	if err != nil || err1 != nil {
		return err
	}
	var result []models.Task
	for _, t := range ts {
		index := arrays.Contains(t.Users, u)
		if index != -1 {
			result = append(result, t)
		}
	}
	return c.JSON(utils.Success("OK", result))
}

func (uc *UserController) GetagentListTask(c echo.Context) error {
	token := c.Request().Header.Get(echo.HeaderAuthorization)
	claims, _ := middleware.ParseToken(token, "secret")
	uid := claims["account"].(string)
	err, ts := uc.db.GetAllTask()
	err1, u := uc.db.GetUserByID(uid)
	if err != nil || err1 != nil {
		return err
	}
	var result []models.Task
	for _, t := range ts {
		if t.Agentaccount == u.Account {
			result = append(result, t)
		}
	}
	return c.JSON(utils.Success("OK", result))
}

func (uc *UserController) GetallListTask(c echo.Context) error {
	err, ts := uc.db.GetAllTask()
	if err != nil {
		return err
	}
	return c.JSON(utils.Success("OK", ts))
}

func (uc *UserController) GetTaskById(c echo.Context) error {
	Id, _ := strconv.Atoi(c.QueryParam("id"))
	err, t := uc.db.GetTaskByID(uint64(Id))
	if err == nil {
		return c.JSON(utils.Success("success", t))
	}
	return c.JSON(utils.Fail("not found task"))
}
