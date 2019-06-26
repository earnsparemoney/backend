package controllers

import (
	"github.com/earnsparemoney/backend/middleware"
	"github.com/earnsparemoney/backend/models"
	"github.com/earnsparemoney/backend/utils"
	echo "github.com/labstack/echo/v4"
	"io/ioutil"
	"strconv"
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
	uid := c.Param("account")
	Id, _ := strconv.Atoi(c.Param("taskID"))
	var tusers []models.TaskUsers
	if uc.db.Where("task_id = ? AND user_id = ?", Id, uid).Find(&tusers).RecordNotFound() {
		return c.JSON(utils.Fail("User not in task"))
	}
	if len(tusers) <= 0 {
		return c.JSON(utils.Fail("User not in task"))
	}
	uc.db.Where("task_id = ? AND user_id = ?", Id, uid).Model(models.TaskUsers{}).Updates(map[string]interface{}{"finished": 1})
	return c.JSON(utils.Success("finish successfully"))
}

func (uc *UserController) DeleteTask(c echo.Context) error {
	//Id,_ := strconv.Atoi(c.QueryParam("id"))
	Id, _ := strconv.Atoi(c.Param("taskID"))
	err, t := uc.db.GetTaskByID(uint64(Id))
	if err != nil {
		return err
	}
	uc.db.Where("task_id = ?", Id).Delete(models.TaskUsers{})
	uc.db.DeleteTask(&t)
	return c.JSON(utils.Success("delete successfully"))
}

func (uc *UserController) ParticipateTask(c echo.Context) error {
	uid := c.Param("account")
	Id, _ := strconv.Atoi(c.Param("taskID"))
	tuser := new(models.TaskUsers)
	tuser.Finished = false
	tuser.TaskID = uint64(Id)
	tuser.UserID = uid
	uc.db.CreateTaskUsers(*tuser)
	return c.JSON(utils.Success("Participate successfully"))
}

func (uc *UserController) CancelTask(c echo.Context) error {
	uid := c.Param("account")
	Id, _ := strconv.Atoi(c.Param("taskID"))
	var tusers []models.TaskUsers
	if uc.db.Where("task_id = ? AND user_id = ?", Id, uid).Find(&tusers).RecordNotFound() {
		return c.JSON(utils.Fail("User not in task"))
	}
	if len(tusers) <= 0 {
		return c.JSON(utils.Fail("User not in task"))
	}
	uc.db.Where("task_id = ? AND user_id = ?", Id, uid).Delete(models.TaskUsers{})
	return c.JSON(utils.Success("Cancel successfully"))
}

func (uc *UserController) GetallListTask(c echo.Context) error {
	err, ts := uc.db.GetAllTask()
	if err != nil {
		return err
	}
	return c.JSON(utils.Success("OK", ts))
}

func (uc *UserController) GetUnCompleteListTask(c echo.Context) error {
	var ts []models.Task
	if uc.db.Where("complete = ?", 0).Find(&ts).RecordNotFound() {
		return c.JSON(utils.Fail("User not in task"))
	}
	if len(ts) <= 0 {
		return c.JSON(utils.Fail("User not in task"))
	}
	return c.JSON(utils.Success("OK", ts))
}

func (uc *UserController) GetListTsak(c echo.Context) error {
	//token:=c.Get("claims").(map[string]interface{})
	//uid := token["account"].(string)
	AgentID := c.QueryParam("agentID")
	UserID := c.QueryParam("userID")
	HasAgentID := false
	HasUserID := false
	var ts []models.Task
	var errA error
	if AgentID != "" {
		HasAgentID = true
		errA, ts = uc.db.GetAllTask()
		if errA != nil {
			return errA
		}
	}
	if UserID != "" {
		HasUserID = true
	}
	var result []models.Task

	for _, t := range ts {
		if HasUserID {
			hasaddtask := false
			for _, tuser := range t.Users {
				if tuser.UserID == UserID {
					hasaddtask = true
					result = append(result, t)
				}
			}
			if hasaddtask {
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
