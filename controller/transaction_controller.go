package controller

import "github.com/labstack/echo/v4"

type TransactionController interface {
	InitializeTransaction(ctx echo.Context) error
	GetByUserID(ctx echo.Context) error
	GetByID(ctx echo.Context) error
}
