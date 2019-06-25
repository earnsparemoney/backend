package routers

import (
	"github.com/earnsparemoney/backend/controllers"
	echo "github.com/labstack/echo/v4"
	//"github.com/earnsparemoney/backend/models"
	"github.com/earnsparemoney/backend/middleware"
)

func RegisterQuesRouters(e *echo.Echo, uc *controllers.UserController) {

	// this should be parameters of REgisterUserRouters too
	//uc := controllers.GetUserController(ds)
	e.POST("/questionnaires", uc.AddQues)
	e.GET("/questionnaires", uc.GetQues)
	e.POST("/questionnaires/:qid/participant/:uid", uc.AddUserToQues)
	e.GET("/questionnaires/:qid", uc.GetQuesByID)
}


/*

	e.POST("/user", uc.AddUser)
	e.GET("/user/login", uc.LoginUser)
	e.GET("/user/:account", uc.GetUserByAccount)
	e.PUT("/user/:account", uc.UpdateUser, middleware.JWTAuth())
	e.GET("/user/logout", uc.LogoutUser)

	TaskRouters(e, uc)
*/