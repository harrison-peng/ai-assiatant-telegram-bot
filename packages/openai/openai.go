package openai

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	client *openai.Client
	ctx    context.Context
}

func NewOpenAI(token string) *OpenAI {
	return &OpenAI{
		client: openai.NewClient(token),
		ctx:    context.Background(),
	}
}

func (c *OpenAI) Chat(messages []openai.ChatCompletionMessage) (string, error) {
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 300,
		Messages:  messages,
	}
	resp, err := c.client.CreateChatCompletion(c.ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
