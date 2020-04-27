package main

import (
	"cookbook/controller"
	"cookbook/model"
	"cookbook/store"
)

func main() {
	r := controller.New()
	routerApi := r.Group("/api")

	d := model.New()
	model.AutoMigrate(d)

	us := store.NewUserStore(d)
	fs := store.NewFoodStore(d)
	ctrl := controller.NewController(us, fs)
	ctrl.Register(routerApi)

	r.Logger.Fatal(r.Start("127.0.0.1:8080"))
}
