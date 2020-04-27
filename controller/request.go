package controller

import (
	"cookbook/model"
	"github.com/labstack/echo"
)

type userRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userRegisterRequest) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Username = r.User.Username
	u.Email = r.User.Email
	hashPswd, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Password = hashPswd
	return nil
}

type userLoginRequest struct {
	User struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userLoginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type foodCreateRequest struct {
	Food struct {
		Name  string `json:"name" validate:"required"`
		Price int64  `json:"price" validate:"required"`
		Desc  string `json:"desc"`
	} `json:"food"`
}

func (r *foodCreateRequest) bind(c echo.Context, f *model.Food) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	f.Name = r.Food.Name
	f.Price = r.Food.Price
	f.Desc = r.Food.Desc
	return nil
}
