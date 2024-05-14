package controller

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"mini_project/constant"
	"mini_project/exception"
	"mini_project/model/web"
	"mini_project/service"
	"mini_project/util"
	"net/http"
	"strconv"
	"strings"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

func (controller *TransactionControllerImpl) InitializeTransaction(c echo.Context) error {
	request := new(web.TransactionCreateRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}

	response, err := controller.TransactionService.Create(request)
	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, web.NewBaseSuccessResponse("transaction created successfully", response))
}

func (controller *TransactionControllerImpl) GetByUserID(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.Split(authHeader, " ")[1]

	token, err := util.ParsingToken(tokenString)

	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["userId"].(float64))

		transactions, err := controller.TransactionService.GetByUserID(userID)

		if err != nil {
			return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("success get transactions", transactions))
	}
	return c.JSON(exception.ErrorHandler(constant.ErrInvalidToken), web.NewBaseErrorResponse(constant.ErrInvalidToken.Error()))
}

func (controller *TransactionControllerImpl) GetByID(c echo.Context) error {
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}

	transaction, err := controller.TransactionService.GetByID(transactionID)
	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("success get transaction", transaction))
}
