package chat

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/chat/v1"
	"google.golang.org/api/option"
)

type ChatInterface interface {
}

type ChatService struct {
	client *chat.Service
}

var IChatService ChatService

func New() *ChatService {

	if (IChatService != ChatService{}) {
		return &IChatService
	}

	client, err := chat.NewService(context.Background(), option.WithCredentialsFile(".credentials.json"))

	if err != nil {
		log.Fatalln(err)
	}

	IChatService = ChatService{
		client: client,
	}

	return &IChatService
}

func (cs *ChatService) SendMessage() {
	spacesService := chat.NewSpacesService(cs.client)
	response, err := spacesService.List().Do()
	if err != nil {
		log.Fatalf("Erreur lors de la récupération des espaces : %v", err)
	}

	// Afficher les IDs des espaces
	for _, space := range response.Spaces {
		fmt.Printf("Nom de l'espace : %s, ID de l'espace : %s\n", space.DisplayName, space.Name)
	}
}
