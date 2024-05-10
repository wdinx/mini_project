package controller

import (
	"github.com/labstack/echo/v4"
	_interface2 "mini_project/controller/interface"
	"mini_project/model/web"
	"mini_project/service/interface"
	"net/http"
)

type TransactionControllerImpl struct {
	TransactionService _interface.TransactionService
}

func NewTransactionController(transactionService _interface.TransactionService) _interface2.TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

func (controller *TransactionControllerImpl) InitializeTransaction(ctx echo.Context) error {
	request := new(web.TransactionCreateRequest)
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := controller.TransactionService.Create(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, response)
}
