package controller

import (
	"cookbook/model"
	"cookbook/utils"
	"github.com/labstack/echo"
	"time"
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

//type foodResponse struct {
//	Food struct {
//		Name  string `json:"name"`
//		Price int64  `json:"price"`
//		Desc  string `json:"desc"`
//	} `json:"food"`
//}

//func newFoodResponse(f *model.Food) *foodResponse {
//	r := new(foodResponse)
//	r.Food.Name = f.Name
//	r.Food.Price = f.Price
//	r.Food.Desc = f.Desc
//	return r
//}

type foodResponse struct {
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type singleFoodResponse struct {
	Food *foodResponse `json:"food"`
}

type foodListResponse struct {
	Foods      []*foodResponse `json:"foods"`
	FoodsCount int             `json:"foodsCount"`
}

func newFoodResponse(c echo.Context, f *model.Food) *singleFoodResponse {
	fr := new(foodResponse)
	fr.Name = f.Name
	fr.Price = f.Price
	fr.Desc = f.Desc
	return &singleFoodResponse{fr}
}

func newFoodListResponse(foods []model.Food, count int) *foodListResponse {
	r := new(foodListResponse)
	r.Foods = make([]*foodResponse, 0)
	for _, food := range foods {
		fr := new(foodResponse)
		fr.Name = food.Name
		fr.Price = food.Price
		fr.Desc = food.Desc
		r.Foods = append(r.Foods, fr)
	}
	r.FoodsCount = count
	return r
}
