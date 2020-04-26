package main

import (
	"cookbook/controller"
	"cookbook/db"
	"cookbook/router"
	"cookbook/store"
)

func main() {

	r := router.New()
	routerApi := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)
	ctrl := controller.NewController(us)
	ctrl.Register(routerApi)

	r.Logger.Fatal(r.Start("127.0.0.1:8080"))

	////实例化echo对象。
	//e := echo.New()
	////注册一个Get请求, 路由地址为: / 并且绑定一个控制器函数, 这里使用的是闭包函数。
	//e.GET("/", func(c echo.Context) error {
	//	//控制器函数直接返回一个hello world，http响应状态为http.StatusOK，200状态。
	//	return c.String(http.StatusOK, "hello world")
	//})
	////启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，相当于0.0.0.0。
	//e.Logger.Fatal(e.Start(":8080"))
}
