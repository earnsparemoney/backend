package main 

import (
	"github.com/labstack/echo"
	"github.com/earnsparemoney/backend/routers"
	"github.com/labstack/echo/middleware"
	"github.com/earnsparemoney/backend/utils"
)


func main(){
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//v1 := e.Group("/v1")

	utils.InitDBConn()
	db :=utils.GetDBConn()

	routers.RegisterUserRouters(e)

	e.Logger.Fatal(e.Start(":1323"))
}
