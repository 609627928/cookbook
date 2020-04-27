package controller

import (
	"cookbook/model"
	"cookbook/utils"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func (ctrl *Controller) CreateFood(c echo.Context) error {
	var f model.Food
	req := &foodCreateRequest{}
	if err := req.bind(c, &f); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	err := ctrl.foodStore.CreateFood(&f)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newFoodResponse(c, &f))
}

func (ctrl *Controller) Foods(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	foods, count, err := ctrl.foodStore.List(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, newFoodListResponse(foods, count))
}
