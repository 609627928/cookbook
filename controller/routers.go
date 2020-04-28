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
			//TokenLookup: "query:token", // 规定了token从哪取，默认从header里，且字段名默认为Authorization
			//AuthScheme:  "",            // 规定了token value前的字段，默认为"Bearer"，与实际token用空格分开
		},
	))
	foods.POST("", ctrl.CreateFood)
	foods.GET("", ctrl.Foods)
	//foods.PUT("/:fid", ctrl.UpdateFood)
	//foods.DELETE("/:fid", ctrl.DeleteFood)

}
