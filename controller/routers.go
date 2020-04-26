package controller

import (
	"github.com/labstack/echo"
)

func (ctrl *Controller) Register(routerApi *echo.Group) {
	routerUser := routerApi.Group("/user")
	routerUser.POST("/signup", ctrl.SignUp)
	routerUser.POST("/login", ctrl.Login)

}
