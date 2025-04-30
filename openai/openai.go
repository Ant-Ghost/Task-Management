package openai

import (
	"context"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

var OpenAIClient *openai.Client

func InitOpenAIClient() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable not set")
	}

	// Initialize OpenAI client
	OpenAIClient = openai.NewClient(apiKey)

	generateResponseFormat()
	generateSystemMessage()
}

func GetAIResponse(prompt string) (string, error) {

	UserMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	}

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			UserMessage,
			SystemMessage,
		},
		// ResponseFormat: JsonResponseFomat,
	}

	resp, err := OpenAIClient.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
