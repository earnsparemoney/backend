package routers

import (
	"github.com/earnsparemoney/backend/controllers"
	echo "github.com/labstack/echo/v4"
	//"github.com/earnsparemoney/backend/models"
)

func TaskRouters(e *echo.Echo, uc *controllers.UserController) {
	e.POST("/tasks/publish", uc.PublishTask)
	e.GET("/tasks/complete", uc.CompleteTask)
	e.GET("/tasks/done", uc.DoneuserTask)
	e.GET("/tasks/delete", uc.DeleteTask)
	e.GET("/tasks/participate", uc.ParticipateTask)
	e.GET("/tasks/cancel", uc.CancelTask)
	e.GET("/tasks/details", uc.GetTaskById)
	e.GET("/tasks/list_user", uc.GetuserListTask)
	e.GET("/tasks​/list_agent", uc.GetagentListTask)
	e.GET("/tasks/list_all", uc.GetallListTask)
}
