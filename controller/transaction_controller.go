package controller

import "github.com/labstack/echo/v4"

type TransactionController interface {
	InitializeTransaction(ctx echo.Context) error
}
