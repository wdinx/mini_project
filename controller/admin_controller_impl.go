package controller

import (
	"github.com/labstack/echo/v4"
	_interface2 "mini_project/controller/interface"
	"mini_project/model/web"
	"mini_project/service/interface"
	"net/http"
)

type AdminControllerImpl struct {
	adminService _interface.AdminService
}

func NewAdminController(adminService _interface.AdminService) _interface2.AdminController {
	return &AdminControllerImpl{
		adminService: adminService,
	}
}

func (controller *AdminControllerImpl) Login(e echo.Context) error {
	adminLoginRequest := web.AdminLoginRequest{}
	if err := e.Bind(&adminLoginRequest); err != nil {
		return e.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	admin, err := controller.adminService.Login(adminLoginRequest)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, web.NewBaseSuccessResponse("login successfully", admin))
}

func (controller *AdminControllerImpl) Register(e echo.Context) error {
	adminRegisterRequest := web.AdminRegisterRequest{}
	if err := e.Bind(&adminRegisterRequest); err != nil {
		return e.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}

	result, err := controller.adminService.Register(&adminRegisterRequest)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, web.NewBaseSuccessResponse("admin created successfully", result))
}
