package services

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"os"
)

func consumeApi(message openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error) {
	token := os.Getenv("GPT_API_TOKEN")
	client := openai.NewClient(token)
	return client.CreateChatCompletion(
		context.Background(),
		message,
	)
}

func SendMessage(message []openai.ChatCompletionMessage, token uint) (openai.ChatCompletionResponse, error) {
	resp, err := consumeApi(openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		Messages:  message,
		MaxTokens: int(token),
	})

	if err != nil {
		return openai.ChatCompletionResponse{}, err
	}
	//return resp.Choices[0].Message.Content, nil
	return resp, nil
}
