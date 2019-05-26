package controllers

import (
	//"github.com/earnsparemoney/backend/models"
	"github.com/labstack/echo"
	"net/http"
	"github.com/earnsparemoney/backend/utils"
	"github.com/earnsparemoney/backend/models"
)

type UserController struct{
	db *models.DBStore
}

func GetUserController(db * models.DBStore) *UserController{
	uc := &UserController{db}
	return uc
}

//get parameters from request body
//check if some parameters are invalid
// call model's function to add a new user
func (uc *UserController)AddUser(c echo.Context) error{
	u := new(models.User)
	if err:=c.Bind(u); err!=nil{
		return err
	}

	if errGet,_:=uc.db.GetUserByID(u.Account);errGet == nil{
		return c.JSON(utils.Fail("already has user with same account"))
	}
	uc.db.CreateUser(*u)
	return c.JSON(utils.Success("success"))
	
}

func (uc *UserController)LoginUser(c echo.Context) error{
	return c.String(http.StatusOK, "hello world")	
}

func (uc *UserController)GetUserByAccount(c echo.Context) error{
	uid := c.Param("account")
	err,u := uc.db.GetUserByID(uid)
	if err ==nil {
		return c.JSON(utils.Success("success",u))
	}
	return c.JSON(utils.Fail("not found user"))
	//return c.String(http.StatusOK, "hello world")	
}

func (uc *UserController)UpdateUser(c echo.Context) error{
	return c.String(http.StatusOK, "hello world")	
}

func (uc *UserController)LogoutUser(c echo.Context) error{
	return c.String(http.StatusOK, "hello world")	
}