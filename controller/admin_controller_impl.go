package controller

import (
	"github.com/labstack/echo/v4"
	"mini_project/exception"
	"mini_project/model/web"
	"mini_project/service"
	"net/http"
)

type AdminControllerImpl struct {
	adminService service.AdminService
}

func NewAdminController(adminService service.AdminService) AdminController {
	return &AdminControllerImpl{
		adminService: adminService,
	}
}

func (controller *AdminControllerImpl) Login(c echo.Context) error {
	adminLoginRequest := web.AdminLoginRequest{}
	if err := c.Bind(&adminLoginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	admin, err := controller.adminService.Login(adminLoginRequest)
	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("login successfully", admin))
}

func (controller *AdminControllerImpl) Register(c echo.Context) error {
	adminRegisterRequest := web.AdminRegisterRequest{}
	if err := c.Bind(&adminRegisterRequest); err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}

	result, err := controller.adminService.Register(&adminRegisterRequest)
	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, web.NewBaseSuccessResponse("admin created successfully", result))
}
