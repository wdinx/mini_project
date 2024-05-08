package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
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

func (controller *MidtransControllerImpl) PaymentHandler(e echo.Context) error {
	var notificationPayload map[string]interface{}
	if err := e.Bind(&notificationPayload); err != nil {
		return e.JSON(400, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	orderID, exists := notificationPayload["order_id"].(string)
	fmt.Println(notificationPayload)
	if !exists {
		return e.JSON(http.StatusBadRequest, web.NewBaseErrorResponse("order_id not found"))
	}

	success, _ := controller.midtransService.VerifyPayment(orderID)
	if !success {
		return e.JSON(http.StatusBadRequest, web.NewBaseErrorResponse("payment not confirmed"))
	}
	fmt.Println("payment handler calling")
	_ = controller.paymentService.ConfirmedPayment(orderID)
	return e.JSON(http.StatusOK, web.NewBaseSuccessResponse("payment confirmed", nil))
}
