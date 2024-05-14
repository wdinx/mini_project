package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"mini_project/exception"
	"mini_project/model/web"
	"mini_project/service"
	"net/http"
)

type MidtransControllerImpl struct {
	midtransService service.MidtransService
	paymentService  service.PaymentService
}

func NewMidtransController(midtransService service.MidtransService, paymentService service.PaymentService) MidtransController {
	return &MidtransControllerImpl{midtransService, paymentService}
}

func (controller *MidtransControllerImpl) PaymentHandler(c echo.Context) error {
	var notificationPayload map[string]interface{}
	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(exception.ErrorHandler(err), map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	orderID, exists := notificationPayload["order_id"].(string)
	fmt.Println(notificationPayload)
	if !exists {
		return c.JSON(http.StatusNotFound, web.NewBaseErrorResponse("order_id not found"))
	}

	success, _ := controller.midtransService.VerifyPayment(orderID)
	if !success {
		return c.JSON(http.StatusPaymentRequired, web.NewBaseErrorResponse("payment not confirmed"))
	}
	fmt.Println("payment handler calling")
	_ = controller.paymentService.ConfirmedPayment(orderID)
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("payment confirmed", nil))
}
