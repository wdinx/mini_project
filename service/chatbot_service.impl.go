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

func (service ChatBotServiceImpl) ChatBotRequest(message string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_KEY"))
	conversation := new([]openai.ChatCompletionMessage)

	*conversation = append(*conversation, openai.ChatCompletionMessage{
		Role:    "system",
		Content: "kamu adalah seorang pemandu wisata. user akan mencari tempat wisata sesuai dengan kondisi atau preferensinya, misalnya: tenang, tenang, cocok untuk petualangan, dekat dengan alam, cocok untuk keluarga, memiliki pemandangan indah, dll. kamu tidak perlu memberi tahu tempat wisatanya dengan spesifik, cukup jenisnya  seperti gunung, pantai, danau, hutan, dll beserta alasan kamu merekomendasikan tempat dengan jenis seperti itu.",
	})

	getResponse := func(userMessage string) (string, error) {
		*conversation = append(*conversation, openai.ChatCompletionMessage{
			Role:    "user",
			Content: userMessage,
		})
		response, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo,
				Messages: *conversation,
			},
		)
		if err != nil {
			return "", constant.ErrInternalServer
		}

		reply := response.Choices[0].Message.Content
		*conversation = append(*conversation, openai.ChatCompletionMessage{
			Role:    "assistant",
			Content: reply,
		})

		return reply, nil
	}

	reply, err := getResponse(message)
	if err != nil {
		fmt.Println("Error:", err)
		return "", constant.ErrInternalServer
	}
	fmt.Println("AI:", reply)

	return reply, nil
}
