package main

import (
	"context"
	"fmt"

	"github.com/markettools-ai/poggers"
	"github.com/sashabaranov/go-openai"
)

var openAIAPIKey *string

// Simple function to send a request to the OpenAI API
func SendMessages(messages []poggers.Message, model string) (string, error) {
	fmt.Println("Sending messages to OpenAI API...", messages)
	client := openai.NewClient(*openAIAPIKey)
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: castToOpenAIMessage(messages),
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to create chat completion: %w", err)
	}

	return response.Choices[0].Message.Content, nil
}

// Helper function to cast the messages to the OpenAI format
func castToOpenAIMessage(messages []poggers.Message) []openai.ChatCompletionMessage {
	var openAIMessages []openai.ChatCompletionMessage
	for _, message := range messages {
		openAIMessages = append(openAIMessages, openai.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
		})
	}
	return openAIMessages
}
