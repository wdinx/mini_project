package controller

import "github.com/labstack/echo/v4"

type TouristAttractionTypeController interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetAll(c echo.Context) error
	FindByID(c echo.Context) error
}
