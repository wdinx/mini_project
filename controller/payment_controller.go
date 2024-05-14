package controller

import "github.com/labstack/echo/v4"

type PaymentController interface {
	InitializePayment(c echo.Context) error
}
