package controller

import (
	"cookbook/food"
	"cookbook/user"
)

type Controller struct {
	userStore user.Store
	foodStore food.Store
}

func NewController(us user.Store, fs food.Store) *Controller {
	return &Controller{
		userStore: us,
		foodStore: fs,
	}
}
