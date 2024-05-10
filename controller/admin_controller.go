package controller

import "github.com/labstack/echo/v4"

type AdminController interface {
	Login(e echo.Context) error
	Register(e echo.Context) error
}
