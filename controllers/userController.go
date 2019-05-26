package controllers

import (
	//"github.com/earnsparemoney/backend/models"
	"github.com/labstack/echo"
	"net/http"
	"github.com/earnsparemoney/backend/utils"
	"github.com/earnsparemoney/backend/models"
)

//get parameters from request body
//check if some parameters are invalid
// call model's function to add a new user
func AddUser(c echo.Context) error{
	u := new(models.User)
	if err:=c.Bind(u); err!=nil{
		return err
	}

	if errGet,_:=models.GetUserByID(u.Account);errGet == nil{
		return c.JSON(utils.Fail("already has user with same account"))
	}
	models.CreateUser(*u)
	return c.JSON(utils.Success("success"))
	
}

func LoginUser(c echo.Context) error{
	return c.String(http.StatusOK, "hello world")	
}

func GetUserByAccount(c echo.Context) error{
	uid := c.Param("account")
	err,u := models.GetUserByID(uid)
	if err ==nil {
		return c.JSON(utils.Success("success",u))
	}
	return c.JSON(utils.Fail("not found user"))
	//return c.String(http.StatusOK, "hello world")	
}

func UpdateUser(c echo.Context) error{
	return c.String(http.StatusOK, "hello world")	
}

func LogoutUser(c echo.Context) error{
	return c.String(http.StatusOK, "hello world")	
}