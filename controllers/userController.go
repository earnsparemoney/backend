package controllers

import (
	//"github.com/earnsparemoney/backend/models"
	"github.com/labstack/echo"
	"net/http"
	"github.com/earnsparemoney/backend/utils"
)

//get parameters from request body
//check if some parameters are invalid
// call model's function to add a new user
func AddUser(c echo.Context) error{
	
}

func LoginUser(c echo.Context) error{
	return c.String(http.StatusOK, "hello world")	
}

func GetUserByAccount(c echo.Context) error{
	return c.String(http.StatusOK, "hello world")	
}

func UpdateUser(c echo.Context) error{
	return c.String(http.StatusOK, "hello world")	
}

func LogoutUser(c echo.Context) error{
	return c.String(http.StatusOK, "hello world")	
}