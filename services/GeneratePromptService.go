package services

import (
	"github.com/sashabaranov/go-openai"
	"log"
)

func GetListOfPossibleProject() {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "¿Puedes proporcionarme una lista de posibles proyectos?",
		},
	}

	ms, err := SendMessage(messages, 30)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(ms)

}
