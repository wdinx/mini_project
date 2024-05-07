package controller

import (
	"github.com/labstack/echo/v4"
	"mini_project/model/web"
	"mini_project/service"
	"net/http"
)

type PaymentControllerImpl struct {
	paymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) PaymentController {
	return &PaymentControllerImpl{paymentService}
}

func (controller PaymentControllerImpl) InitializePayment(e echo.Context) error {
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

func (controller PaymentControllerImpl) ConfirmedPayment(e echo.Context) error {
	//TODO implement me
	panic("implement me")
}
