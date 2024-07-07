package service

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"mini_project/constant"
	"os"
)

type ChatBotServiceImpl struct {
}

func NewChatBotService() ChatBotService {
	return &ChatBotServiceImpl{}
}

func (service ChatBotServiceImpl) ChatBotRequest(conversation *[]openai.ChatCompletionMessage) (*[]openai.ChatCompletionMessage, error) {
	client := openai.NewClient(os.Getenv("OPENAI_KEY"))
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: *conversation,
		},
	)
	if err != nil {
		return nil, constant.ErrInternalServer
	}

	reply := response.Choices[0].Message.Content
	*conversation = append(*conversation, openai.ChatCompletionMessage{
		Role:    "assistant",
		Content: reply,
	})
	fmt.Println(reply)

	return conversation, nil
}
