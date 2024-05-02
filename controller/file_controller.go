package controller

import "github.com/labstack/echo/v4"

type FileController interface {
	ShowFile(c echo.Context) error
}
