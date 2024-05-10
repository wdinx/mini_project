package controller

import (
	"github.com/labstack/echo/v4"
	_interface2 "mini_project/controller/interface"
	"mini_project/model/web"
	"mini_project/service/interface"
	"net/http"
)

type PaymentControllerImpl struct {
	paymentService _interface.PaymentService
}

func NewPaymentController(paymentService _interface.PaymentService) _interface2.PaymentController {
	return &PaymentControllerImpl{paymentService}
}

func (controller *PaymentControllerImpl) InitializePayment(e echo.Context) error {
	var payment web.PaymentRequest
	err := e.Bind(&payment)
	if err != nil {
		return e.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	response, err := controller.paymentService.InitializePayment(&payment)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}
	return e.JSON(http.StatusOK, web.NewBaseSuccessResponse("payment initialized successfully", response))
}
