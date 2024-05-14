package controller

import "github.com/labstack/echo/v4"

type AdminController interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}
