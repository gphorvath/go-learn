package client

import (
	"context"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

type AnthropicClient struct {
	client *anthropic.Client
}

func NewAnthropic() *AnthropicClient {
	apiKey, exists := os.LookupEnv("ANTHROPIC_API_KEY")
	if !exists || apiKey == "" {
		panic("ANTHROPIC_API_KEY is not set")
	}

	return &AnthropicClient{
		client: anthropic.NewClient(
			option.WithAPIKey(apiKey),
		),
	}
}

func (c *AnthropicClient) Chat(ctx context.Context, prompt string) (string, error) {
	message, err := c.client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.F(anthropic.ModelClaude_3_5_Sonnet_20240620),
		MaxTokens: anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock("What is a quaternion?")),
		}),
	})

	if err != nil {
		return "", err
	}

	return message.Content[0].Text, nil
}
