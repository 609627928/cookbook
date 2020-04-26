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
}
