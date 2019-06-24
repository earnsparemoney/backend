package middleware

import (
	echo "github.com/labstack/echo/v4"
	"github.com/earnsparemoney/backend/utils"
	//"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"fmt"
)


var secret =  "secret"

//check if token is valid 
// if token is valid pass 
//	if token is invalid or there is no token in the header 
// then check if there is corresponding session in the request cookie 
// if there is no seesionId, then the request is invalid 
// if there is sessionID, then check if sessionID is valid 
// if sessionId is valid, then create token and return, this request API will not response
//	if sessionID is invalid, then return invlaid request. 
func JWTAuth() echo.MiddlewareFunc{
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error{
			token :=c.Request().Header.Get(echo.HeaderAuthorization)
			fmt.Println(token)
			claims,err := ParseToken(token,secret)
			fmt.Println(claims,err)
			if err !=nil{
				sess,_ :=session.Get("session",c)
				if account, ok := sess.Values["sessionID"].(string) ;ok {
					if account == ""{
						c.JSON(utils.Fail("not session, please log in"))
					} else {
						token :=&MyToken{account}
						tokenString, err := token.GenerateToken(secret)
						if err !=nil{
							return c.JSON(utils.Error(err.Error()))
						}
						return c.JSON(utils.Success("return a new token",tokenString))
					}

				}
				
				return c.JSON(utils.Fail(err.Error()))

			}
			c.Set("claims",claims)
			return next(c)
		}
	}
}


