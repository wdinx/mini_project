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

type TicketControllerImpl struct {
	ticketService service.TicketService
}

func NewTicketController(ticketService service.TicketService) TicketController {
	return &TicketControllerImpl{
		ticketService: ticketService,
	}
}

func (controller *TicketControllerImpl) FindByUserID(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.Split(authHeader, " ")[1]

	token, err := util.ParsingToken(tokenString)

	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["userId"].(float64))

		tickets, err := controller.ticketService.FindByUserID(userID)

		if err != nil {
			return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("success get tickets", tickets))
	}

	return c.JSON(exception.ErrorHandler(constant.ErrUnauthorized), web.NewBaseErrorResponse(constant.ErrUnauthorized.Error()))
}

func (controller *TicketControllerImpl) FindByTouristAttractionID(c echo.Context) error {

	touristAttractionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}

	tickets, err := controller.ticketService.FindByTouristAttractionID(touristAttractionID)
	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("success get tickets", tickets))

}

func (controller *TicketControllerImpl) FindByID(c echo.Context) error {

	ticketID := c.Param("id")
	ticket, err := controller.ticketService.FindByID(ticketID)
	if err != nil {
		return c.JSON(http.StatusNotFound, web.NewBaseErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("success get ticket", ticket))
}
