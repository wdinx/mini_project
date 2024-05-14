package controller

import (
	"github.com/labstack/echo/v4"
	"mini_project/exception"
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

func (controller *PaymentControllerImpl) InitializePayment(c echo.Context) error {
	var payment web.PaymentRequest
	err := c.Bind(&payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	response, err := controller.paymentService.InitializePayment(&payment)
	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, web.NewBaseSuccessResponse("payment initialized successfully", response))
}
