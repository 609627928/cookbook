package controller

import (
	"cookbook/model"
	"cookbook/utils"
	"github.com/labstack/echo"
	"net/http"
)

func (ctrl *Controller) SignUp(c echo.Context) error {
	var u model.User
	req := &userRegisterRequest{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := ctrl.userStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	//return c.JSON(http.StatusCreated, "SignUp success")
	return c.JSON(http.StatusCreated, newUserResponse(&u))
}

func (ctrl *Controller) Login(c echo.Context) error {
	req := &userLoginRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u, err := ctrl.userStore.GetByEmail(req.User.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	if !u.CheckPassword(req.User.Password) {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	//return c.JSON(http.StatusOK, "Login success")
	return c.JSON(http.StatusOK, newUserResponse(u))
}

//import "net/http"
//
//// 路由定义post请求, url路径为：/users, 绑定saveUser控制器函数
//e.POST("/users", saveUser)
//
//// 控制器函数
//func saveUser(c echo.Context) error {
//	username := c.FormValue("username")
//	username := c.FormValue("username")
//
//	//调用model保存数据
//
//	html := 调用模板引擎渲染html页面
//	//以Html页面的形式响应请求
//	return c.HTML(http.StatusOK, html)
//}
