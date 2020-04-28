package controller

import (
	"cookbook/food"
	"cookbook/user"
	"cookbook/utils"
	"fmt"
)

type Controller struct {
	userStore user.Store
	foodStore food.Store
	//foodStore *store.FoodStore
}

//func NewController(us *store.UserStore, fs *store.FoodStore) *Controller {
func NewController(us user.Store, fs food.Store) *Controller {
	fmt.Println(utils.TypeOf(us))
	return &Controller{
		userStore: us,
		foodStore: fs,
	}
}
