package main

import (
	"context"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {

	apiKey, exists := os.LookupEnv("OPENAI_API_KEY")

	if !exists || apiKey == "" {
		panic("OPENAI_API_KEY is not set")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	ctx := context.Background()

	question := "Write me a haiku"

	print("> ")
	println(question)
	println()

	completion, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
		}),
		Seed:  openai.Int(1),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		panic(err)
	}

	println(completion.Choices[0].Message.Content)
}
