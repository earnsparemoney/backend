package main

import (
	"github.com/earnsparemoney/backend/routers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	//"github.com/earnsparemoney/backend/utils"
	"github.com/earnsparemoney/backend/controllers"
	"github.com/earnsparemoney/backend/models"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	//"golang.org/x/crypto/acme/autocert"
)

type Env struct {
	db   *models.DBStore
	echo *echo.Echo
	uc   *controllers.UserController
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//v1 := e.Group("/v1")

	dbConfig := "root:sactestdatabase@tcp(222.200.190.31:33336)/earnsparemoney?charset=utf8&parseTime=True&loc=Local"

	db, err := models.NewDB(dbConfig)
	if err != nil {
		panic(err)
	}

	//e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	env := &Env{db, e, controllers.GetUserController(db)}
	env.db.UserModelInit()
	env.db.TaskModelInit()

	routers.RegisterUserRouters(env.echo, env.uc)

	e.Logger.Fatal(e.StartTLS(":443", "server.crt", "server.key"))
}
