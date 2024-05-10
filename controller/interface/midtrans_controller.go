package _interface

import "github.com/labstack/echo/v4"

type MidtransController interface {
	PaymentHandler(e echo.Context) error
}
