package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var contextWindow []string // Stores the context window for the conversation

func createOpenAIClient() *openai.Client {
	apiKey, exists := os.LookupEnv("OPENAI_API_KEY")

	if !exists || apiKey == "" {
		panic("OPENAI_API_KEY is not set")
	}

	return openai.NewClient(
		option.WithAPIKey(apiKey),
	)
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			log.Printf("Error reading input: %v", err)
			return ""
		}
	}

	// Trim spaces and newline characters
	input = strings.TrimSpace(input)

	// Handle empty input
	if len(input) == 0 {
		return ""
	}

	return input
}

func processUserInput(input string) string {
	switch strings.ToLower(input) {
	case "exit", "quit":
		fmt.Println("Exiting...")
		os.Exit(0)
		return ""
	case "clear":
		contextWindow = []string{}
		fmt.Println("Context window cleared.")
		return ""
	case "help":
		fmt.Println("Available commands:")
		fmt.Println("  help  - Show this help message")
		fmt.Println("  clear - Clear the context window")
		fmt.Println("  exit or quit  - Exit the program")
		return ""
	default:
		return input
	}
}

func buildPrompt(question string) string {
	contextWindow = append(contextWindow, question)
	return strings.Join(contextWindow, "\n")
}

func main() {

	client := createOpenAIClient()
	ctx := context.Background()

	for {
		userQuestion := getUserInput("Ask a question (or type 'help' for a list of options): ")

		if processUserInput(userQuestion) == "" {
			continue
		}

		llm_prompt := buildPrompt(userQuestion)

		print("> ")
		println(userQuestion)
		println()

		completion, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
			Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(llm_prompt),
			}),
			Seed:  openai.Int(1),
			Model: openai.F(openai.ChatModelGPT4o),
		})
		if err != nil {
			log.Printf("Error getting completion: %v", err)
			continue
		}

		contextWindow = append(contextWindow, completion.Choices[0].Message.Content)

		println(completion.Choices[0].Message.Content)
	}
}
