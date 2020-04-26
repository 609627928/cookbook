package controller

import (
	"cookbook/model"
	"cookbook/utils"
)

type userResponse struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	} `json:"user"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Token = utils.GenerateJWT(u.ID)
	return r
}
