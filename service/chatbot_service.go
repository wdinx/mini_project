package service

type ChatBotService interface {
	ChatBotRequest(message string) (string, error)
}
