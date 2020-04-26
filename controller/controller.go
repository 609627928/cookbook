package controller

import (
	"cookbook/user"
)

type Controller struct {
	userStore user.Store
}

func NewController(us user.Store) *Controller {
	return &Controller{
		userStore: us,
	}
}
