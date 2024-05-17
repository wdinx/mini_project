package controller

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"mini_project/service"
	"net/http"
)

type ChatBotControllerImpl struct {
	chatBotService service.ChatBotService
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewChatBotController(chatBotService service.ChatBotService) ChatBotController {
	return &ChatBotControllerImpl{
		chatBotService: chatBotService,
	}
}

func (controller ChatBotControllerImpl) ChatBot(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			return err
		}

		response, err := controller.chatBotService.ChatBotRequest(string(msg))
		if err != nil {
			panic(err)
		}

		err = ws.WriteMessage(websocket.TextMessage, []byte(response))
		if err != nil {
			return err
		}
	}
}
