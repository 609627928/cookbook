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

func (ctrl *Controller) UpdateFood(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	f, err := ctrl.foodStore.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if f == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	req := &foodUpdateRequest{}
	//req.populate(f) // 为了不传参数时，不将数据库内置空，这里将数据库的值填充至结构体
	if err := req.bind(c, f); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = ctrl.foodStore.UpdateFood(f); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newFoodResponse(c, f))
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

func (ctrl *Controller) DeleteFood(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	f, err := ctrl.foodStore.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if f == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	err = ctrl.foodStore.DeleteFood(f)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}
