package main 

import (
	"github.com/labstack/echo"
	"github.com/earnsparemoney/backend/routers"
	"github.com/labstack/echo/middleware"
	//"github.com/earnsparemoney/backend/utils"
	"github.com/earnsparemoney/backend/models"
)

type Env struct {
	db *models.DBStore
	echo *echo.Echo
}


func main(){
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//v1 := e.Group("/v1")

	dbConfig := "root:sactestdatabase@tcp(222.200.190.31:33336)/earnsparemoney?charset=utf8&parseTime=True&loc=Local"

	db, err := models.NewDB(dbConfig)
	if err !=nil{
		panic(err)
	}

	env :=&Env{db, e}
	env.db.UserModelInit()


	routers.RegisterUserRouters(env.echo, env.db)

	e.Logger.Fatal(e.Start(":1323"))
}