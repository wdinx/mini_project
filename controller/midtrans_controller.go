package controller

import "github.com/labstack/echo/v4"

type MidtransController interface {
	PaymentHandler(c echo.Context) error
}
