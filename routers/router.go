package routers

import (
	"github.com/labstack/echo"
	"github.com/earnsparemoney/backend/controllers"
)

func RegisterUserRouters(e *echo.Echo){
	e.POST("/user",controllers.AddUser)
	e.GET("/user/login",controllers.LoginUser)
	e.GET("/user/:account",controllers.GetUserByAccount)
	e.PUT("/user/:account",controllers.UpdateUser)
	e.GET("/user/logout",controllers.LogoutUser)

}