package middleware


import(
	"github.com/dgrijalva/jwt-go"
	"time"
	"errors"
	
)

type MyToken struct{
	Account string
}


func (t *MyToken)GenerateToken(secret string) (string,error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["account"] = t.Account
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	
	encodeToken, err := token.SignedString([]byte(secret))
	if err !=nil{
		return "",err
	}
	return encodeToken,nil

}



func ParseToken(token string, secret string) (jwt.MapClaims,error){
	if token == ""{
		return nil, errors.New("not found token")
	}
	tokenParse,err := jwt.Parse(token, func(token *jwt.Token)(interface{},error){
		return []byte(secret), nil
	})
	switch  err.(type){
	case nil:
		if !tokenParse.Valid{
			return nil, errors.New("token invalid")
		}

		if claims, ok :=tokenParse.Claims.(jwt.MapClaims); ok{
			return claims,nil
		}
		return nil,errors.New("err parse claims")
		//return tokenParse.Claims.(jwt.MapClaims),nil

	case *jwt.ValidationError:
		return nil,errors.New("token validation error")
	
	default:
		return nil,errors.New("token invalid")
	}


}


/*
func ValidateToken(token string, secret string) bool{


}
*/

