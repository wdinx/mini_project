package controller

import (
	"github.com/labstack/echo/v4"
	"mini_project/controller/interface"
)

type FileControllerImpl struct{}

func NewFileController() _interface.FileController {
	return &FileControllerImpl{}
}

func (controller *FileControllerImpl) ShowFile(c echo.Context) error {
	file := c.Param("image")
	return c.File("public/" + file)
}
