package controller

import "github.com/labstack/echo/v4"

type TicketController interface {
	FindByUserID(c echo.Context) error
	FindByTouristAttractionID(c echo.Context) error
	FindByID(c echo.Context) error
}
