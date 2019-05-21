package main 

import (
	"github.com/labstack/echo"
	"github.com/earnsparemoney/backend/routers"
	"github.com/labstack/echo/middleware"
)


func main(){
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//v1 := e.Group("/v1")

	routers.RegisterUserRouters(e)

	e.Logger.Fatal(e.Start(":1323"))
}
