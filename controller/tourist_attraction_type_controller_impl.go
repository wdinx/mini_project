package controller

import (
	"github.com/labstack/echo/v4"
	_interface2 "mini_project/controller/interface"
	"mini_project/model/web"
	"mini_project/service/interface"
	"net/http"
	"strconv"
)

type TouristAttractionTypeControllerImpl struct {
	touristAttractionTypeService _interface.TouristAttractionTypeService
}

func NewTouristAttractionTypeController(touristAttractionTypeService _interface.TouristAttractionTypeService) _interface2.TouristAttractionTypeController {
	return &TouristAttractionTypeControllerImpl{touristAttractionTypeService: touristAttractionTypeService}
}

func (controller *TouristAttractionTypeControllerImpl) Create(c echo.Context) error {
	var err error
	var touristAttractionTypeCreateRequest web.TouristAttractionTypeCreateRequest
	err = c.Bind(&touristAttractionTypeCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	result, err := controller.touristAttractionTypeService.Create(&touristAttractionTypeCreateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("tourist attraction type created successfully", result))
}

func (controller *TouristAttractionTypeControllerImpl) Update(c echo.Context) error {
	var err error
	var touristAttractionTypeUpdateRequest web.TouristAttractionTypeUpdateRequest
	err = c.Bind(&touristAttractionTypeUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	result, err := controller.touristAttractionTypeService.Update(&touristAttractionTypeUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("tourist attraction type updated successfully", result))
}

func (controller *TouristAttractionTypeControllerImpl) Delete(c echo.Context) error {
	var err error
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}
	err = controller.touristAttractionTypeService.Delete(idx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("tourist attraction type deleted successfully", nil))
}

func (controller *TouristAttractionTypeControllerImpl) GetAll(c echo.Context) error {
	//var err error
	result := controller.touristAttractionTypeService.GetAll()
	if len(*result) == 0 {
		return c.JSON(http.StatusNotFound, web.NewBaseErrorResponse("tourist attraction type not found"))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("tourist attraction type found", result))
}
