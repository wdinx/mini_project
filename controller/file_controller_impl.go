package controller

import (
	"github.com/labstack/echo/v4"
)

type FileControllerImpl struct{}

func NewFileController() FileController {
	return &FileControllerImpl{}
}

func (controller FileControllerImpl) ShowFile(c echo.Context) error {
	file := c.Param("image")
	return c.File("public/" + file)
}
