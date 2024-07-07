package controller

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"mini_project/model/web"
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

	var conversation []openai.ChatCompletionMessage
	conversation = append(conversation, openai.ChatCompletionMessage{
		Role:    "system",
		Content: "you are an english teacher",
	})

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			return err
		}

		conversation = append(conversation, openai.ChatCompletionMessage{
			Role:    "user",
			Content: string(msg),
		})

		response, err := controller.chatBotService.ChatBotRequest(&conversation)
		if err != nil {
			panic(err)
		}

		err = ws.WriteMessage(websocket.TextMessage, []byte((*response)[len(*response)-1].Content))
		if err != nil {
			return err
		}
	}
}

func (controller ChatBotControllerImpl) ChatBotRestFul(c echo.Context) error {
	var conversation []openai.ChatCompletionMessage
	conversation = append(conversation, openai.ChatCompletionMessage{
		Role:    "system",
		Content: "you are an english teacher",
	})

	var userMessage UserMessage
	if err := c.Bind(&userMessage); err != nil {
		return err
	}

	conversation = append(conversation, openai.ChatCompletionMessage{
		Role:    "user",
		Content: userMessage.Message,
	})

	response, err := controller.chatBotService.ChatBotRequest(&conversation)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse(
		"success get data response",
		(*response)[len(*response)-1].Content,
	))
}

type UserMessage struct {
	Message string `json:"message"`
}
