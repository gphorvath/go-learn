package client

import (
	"context"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIClient struct {
	client *openai.Client
}

func NewOpenAI() *OpenAIClient {
	apiKey, exists := os.LookupEnv("OPENAI_API_KEY")
	if !exists || apiKey == "" {
		panic("OPENAI_API_KEY is not set")
	}

	return &OpenAIClient{
		client: openai.NewClient(
			option.WithAPIKey(apiKey),
		),
	}
}

func (c *OpenAIClient) Chat(ctx context.Context, prompt string) (string, error) {
	completion, err := c.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Seed:  openai.Int(1),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		return "", err
	}

	return completion.Choices[0].Message.Content, nil
}
