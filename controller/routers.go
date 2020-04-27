package controller

import (
	"cookbook/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Validator = NewValidator()
	return e
}

func (ctrl *Controller) Register(routerApi *echo.Group) {
	users := routerApi.Group("/user")
	users.POST("/signup", ctrl.SignUp)
	users.POST("/login", ctrl.Login)

	foods := routerApi.Group("/food", middleware.JWTWithConfig(
		middleware.JWTConfig{
			Skipper: func(c echo.Context) bool {
				if c.Request().Method == "GET" && c.Path() != "/api/food/feed" {
					return true
				}
				return false
			},
			SigningKey: utils.JWTSecret,
		},
	))
	foods.POST("", ctrl.CreateFood)
	foods.GET("", ctrl.Foods)
	//foods.PUT("/:fid", ctrl.UpdateFood)
	//foods.DELETE("/:fid", ctrl.DeleteFood)

}
