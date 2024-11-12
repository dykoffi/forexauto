package chat

type ChatInterface interface {
}

type ChatService struct{}

var IChatService ChatService

func New() *ChatService {

	return &ChatService{}
}

func (cs *ChatService) SendMessage() error {
	return nil
}
