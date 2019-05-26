package routers

import (
	"github.com/labstack/echo"
	"github.com/earnsparemoney/backend/controllers"
	"github.com/earnsparemoney/backend/models"
)

func RegisterUserRouters(e *echo.Echo,ds *models.DBStore){

	uc := controllers.GetUserController(ds)

	e.POST("/user",uc.AddUser)
	e.GET("/user/login",uc.LoginUser)
	e.GET("/user/:account",uc.GetUserByAccount)
	e.PUT("/user/:account",uc.UpdateUser)
	e.GET("/user/logout",uc.LogoutUser)

} 

