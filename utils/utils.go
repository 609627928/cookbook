package utils

import (
	"fmt"
	"github.com/labstack/echo"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["body"] = v.Message
	default:
		e.Errors["body"] = v.Error()
	}
	return e
}

var JWTSecret = []byte("!!SECRET!!")

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}

func AccessForbidden() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "access forbidden"
	return e
}

func TypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
