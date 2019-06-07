package controllers

import (
	"github.com/labstack/echo"
	//  "net/http"
	"github.com/earnsparemoney/backend/models"
	"github.com/earnsparemoney/backend/utils"
	//  "github.com/gorilla/sessions"
	"fmt"
	"github.com/labstack/echo-contrib/session"
	"github.com/wxnacy/wgo/arrays"
	"strconv"
	"time"
)

func (uc *UserController) PublishTask(c echo.Context) error {
	t := new(models.Task)
	if err := c.Bind(t); err != nil {
		return err
	}
	fmt.Println(t.Ddl)
	tm, err := time.Parse("2006-01-02 15:04:05", t.Ddl)
	fmt.Println(err)
	if err != nil || tm.Before(time.Now()) {
		return c.JSON(utils.Fail("time error"))
	}
	uc.db.CreateTask(*t)
	return c.JSON(utils.Success("success"))
}

func (uc *UserController) CompleteTask(c echo.Context) error {
	Id, _ := strconv.Atoi(c.QueryParam("id"))
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
	//  t.Undousers=append(t.Undousers[:index], t.Undousers[index+1:]...)
	//  t.Doneusers=append(t.Doneusers,u)
	uc.db.UpdateTask(&t)
	return c.JSON(utils.Success("update successfully"))
}

func (uc *UserController) DeleteTask(c echo.Context) error {
	Id, _ := strconv.Atoi(c.QueryParam("id"))
	err, t := uc.db.GetTaskByID(uint64(Id))
	if err != nil {
		return err
	}
	uc.db.DeleteTask(&t)
	return c.JSON(utils.Success("delete successfully"))
}

func (uc *UserController) ParticipateTask(c echo.Context) error {
	sess, _ := session.Get("session", c)
	fmt.Println(sess.Values["sessionID"])
	uid := c.Param("account")
	Id, _ := strconv.Atoi(c.QueryParam("id"))
	err, t := uc.db.GetTaskByID(uint64(Id))
	err1, u := uc.db.GetUserByID(uid)
	if err != nil || err1 != nil {
		return err
	}
	if uc.db.Model(&t).Association("Users").Count() >= t.Usernum {
		return c.JSON(utils.Fail("User num enough"))
	}
	index := arrays.Contains(t.Users, u)
	if index != -1 {
		return c.JSON(utils.Fail("User have join"))
	}
	uc.db.Model(&t).Association("Users").Append(&u)
	uc.db.Model(&t).Association("Undousers").Append(&u)
	//  t.Users=append(t.Users,u)
	//  t.Undousers=append(t.Undousers,u)
	uc.db.UpdateTask(&t)
	return c.JSON(utils.Success("delete successfully"))
}

func (uc *UserController) CancelTask(c echo.Context) error {
	sess, _ := session.Get("session", c)
	fmt.Println(sess.Values["sessionID"])
	uid := c.Param("account")
	Id, _ := strconv.Atoi(c.QueryParam("id"))
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
	sess, _ := session.Get("session", c)
	fmt.Println(sess.Values["sessionID"])
	uid := c.Param("account")
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
	sess, _ := session.Get("session", c)
	fmt.Println(sess.Values["sessionID"])
	uid := c.Param("account")
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
