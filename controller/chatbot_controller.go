package controller

import "github.com/labstack/echo/v4"

type ChatBotController interface {
	ChatBot(c echo.Context) error
}
