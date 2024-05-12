package controller

import "github.com/labstack/echo/v4"

type TouristAttractionController interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	GetAll(c echo.Context) error
	UpdateBalanceById(c echo.Context) error
	GetById(c echo.Context) error
}
