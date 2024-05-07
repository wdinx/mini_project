package controller

import "github.com/labstack/echo/v4"

type PaymentController interface {
	InitializePayment(e echo.Context) error
	ConfirmedPayment(e echo.Context) error
}
