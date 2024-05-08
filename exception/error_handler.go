package exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"mini_project/constant"
	"mini_project/model/web"
	"net/http"
)

func ErrorHandler(e echo.Context, err interface{}) {

}

func validationErrors(e echo.Context, err interface{}) bool {
	_, ok := err.(validator.ValidationErrors)
	if ok {
		e.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(constant.ErrInputData.Error()))
		return true
	} else {
		return false
	}
}
