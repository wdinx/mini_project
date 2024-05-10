package controller

import (
	"github.com/labstack/echo/v4"
	_interface2 "mini_project/controller/interface"
	"mini_project/model/web"
	"mini_project/service/interface"
	"net/http"
)

type UserControllerImpl struct {
	userService _interface.UserService
}

func NewUserController(userService _interface.UserService) _interface2.UserController {
	return &UserControllerImpl{userService: userService}
}

func (controller *UserControllerImpl) Login(c echo.Context) error {
	var err error
	userLoginRequest := web.UserLoginRequest{}
	if err := c.Bind(&userLoginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}

	result, err := controller.userService.Login(userLoginRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("login success", result))

}

func (controller *UserControllerImpl) Register(c echo.Context) error {
	var err error
	userRegisterRequest := web.UserRegisterRequest{}
	if err = c.Bind(&userRegisterRequest); err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}

	userRegisterRequest.ProfilePicture, err = c.FormFile("profile_picture")
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}

	result, err := controller.userService.Register(userRegisterRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("admin created successfully", result))
}
