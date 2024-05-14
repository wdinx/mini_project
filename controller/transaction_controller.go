package controller

import "github.com/labstack/echo/v4"

type TransactionController interface {
	InitializeTransaction(c echo.Context) error
	GetByUserID(c echo.Context) error
	GetByID(c echo.Context) error
}
