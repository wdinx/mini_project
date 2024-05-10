package controller

import (
	"github.com/labstack/echo/v4"
	"mini_project/model/web"
	"mini_project/service"
	"net/http"
	"strconv"
)

type TouristAttractionControllerImpl struct {
	touristAttractionService service.TouristAttractionService
}

func NewTouristAttractionController(touristAttractionService service.TouristAttractionService) TouristAttractionController {
	return &TouristAttractionControllerImpl{touristAttractionService: touristAttractionService}
}

func (controller *TouristAttractionControllerImpl) Create(c echo.Context) error {
	var err error
	var touristAttractionCreateRequest web.TouristAttractionRequest
	if err = c.Bind(&touristAttractionCreateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	touristAttractionCreateRequest.Image, err = c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	result, err := controller.touristAttractionService.Create(&touristAttractionCreateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("tourist attraction created successfully", result))
}

func (controller *TouristAttractionControllerImpl) Update(c echo.Context) error {
	var err error
	var touristAttractionUpdateRequest web.TouristAttractionUpdateRequest
	touristAttractionUpdateRequest.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	if err = c.Bind(&touristAttractionUpdateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	result, err := controller.touristAttractionService.Update(&touristAttractionUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("tourist attraction updated successfully", result))
}

func (controller *TouristAttractionControllerImpl) GetAll(c echo.Context) error {
	result, err := controller.touristAttractionService.GetAllTouristAttraction()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("get all tourist attraction success", *result))
}

func (controller *TouristAttractionControllerImpl) UpdateBalanceById(c echo.Context) error {
	// TODO: Implement this method
	panic("implement me")
	//var err error
	//var touristAttractionUpdateRequest web.TouristAttractionUpdateRequest
	//touristAttractionUpdateRequest.ID, err = strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	//}
	//if err = c.Bind(&touristAttractionUpdateRequest); err != nil {
	//	return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	//}
	//result, err := controller.touristAttractionService.UpdateBalanceById(&touristAttractionUpdateRequest)
	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	//}
	//return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("tourist attraction balance updated successfully", result))
}
