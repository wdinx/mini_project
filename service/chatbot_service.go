package service

import "github.com/sashabaranov/go-openai"

type ChatBotService interface {
	ChatBotRequest(*[]openai.ChatCompletionMessage) (*[]openai.ChatCompletionMessage, error)
}
