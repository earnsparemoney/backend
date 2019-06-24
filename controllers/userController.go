package controllers

import (
	//"github.com/earnsparemoney/backend/models"
	echo "github.com/labstack/echo/v4"
	"net/http"
	"github.com/earnsparemoney/backend/utils"
	"github.com/earnsparemoney/backend/models"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"fmt"
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

//get parameters from body 
//find if account and password is same as db's values
//if right create session return user info
//if wrong, return wrong message 
func (uc *UserController)LoginUser(c echo.Context) error{
	account:=c.QueryParam("account")
	password:=c.QueryParam("password")
	err, u := uc.db.GetUserByID(account)
	if err !=nil{
		return c.JSON(utils.Fail("not found user"))
	}
	if u.Password != password{
		return c.JSON(utils.Fail("password incorrect"))
	}
	sess,_:=session.Get("session",c)
	sess.Options = &sessions.Options{
		Path: "/",
		MaxAge: 86400 * 7,
		HttpOnly: true,
		Secure: false,
	}
	sess.Values["sessionID"] = account
	sess.Save(c.Request(), c.Response())
	fmt.Println(sess)
	return c.JSON(utils.Success("log in success",u))
	//return c.String(http.StatusOK, "hello world")	
}

func (uc *UserController)GetUserByAccount(c echo.Context) error{
	sess,_:=session.Get("session",c)
	fmt.Println(sess.Values["sessionID"])
	uid := c.Param("account")
	err,u := uc.db.GetUserByID(uid)
	if err ==nil {
		return c.JSON(utils.Success("success",u))
	}
	return c.JSON(utils.Fail("not found user"))
	//return c.String(http.StatusOK, "hello world")	
}

//find user if 
func (uc *UserController)UpdateUser(c echo.Context) error{
	u :=new(models.User)
	if err :=c.Bind(u); err!=nil{
		return err
	}
	uc.db.UpdateUser(u)
	return c.JSON(utils.Success("update successfully"))
	//return c.String(http.StatusOK, "hello world")	
}

func (uc *UserController)LogoutUser(c echo.Context) error{
	return c.String(http.StatusOK, "hello world")	
}