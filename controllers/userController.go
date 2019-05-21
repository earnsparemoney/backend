package controllers

import (
	//"github.com/earnsparemoney/backend/models"
	"github.com/labstack/echo"
	"net/http"
)

func AddUser(c echo.Context) error{
	return c.String(http.StatusOK, "hello world")	
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