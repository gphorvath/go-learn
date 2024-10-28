package main

import (
	"bufio"
	"context"
	"fmt"
	"go-learn/chat-bot/client"
	"io"
	"log"
	"os"
	"strings"
)

type Dispatcher struct {
	activeClient string
	clients      map[string]client.ChatClient
	commands     map[string]func(string) string
	context      []string
}

func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		activeClient: "mock",
		clients: map[string]client.ChatClient{
			"anthropic": client.NewAnthropic(),
			"openai":    client.NewOpenAI(),
			"mock": client.NewMock([]string{
				"You are using the mock client. It doesn't do much, but it's here to help you test the chat bot.",
				"Sure, I can help with that. What do you need?",
				"Let me see what I can do. One moment please...",
				"Here is the information you requested.",
				"Is there anything else I can assist you with?",
			}),
		},
		context: make([]string, 0),
	}

	d.commands = map[string]func(string) string{
		"anthropic": d.switchClient,
		"openai":    d.switchClient,
		"mock":      d.switchClient,
		"clear":     d.clearContext,
		"help":      d.showHelp,
		"exit":      d.exit,
		"quit":      d.exit,
	}

	return d
}

func (d *Dispatcher) switchClient(clientName string) string {
	d.activeClient = clientName
	fmt.Printf("Switched to %s client.\n", clientName)
	return ""
}

func (d *Dispatcher) clearContext(_ string) string {
	d.context = []string{}
	fmt.Println("Context window cleared.")
	return ""
}

func (d *Dispatcher) exit(_ string) string {
	fmt.Println("Exiting...")
	os.Exit(0)
	return ""
}

func (d *Dispatcher) showHelp(_ string) string {
	fmt.Println("Available commands:")
	fmt.Println("  anthropic - Switch to the Anthropic client")
	fmt.Println("  openai    - Switch to the OpenAI client")
	fmt.Println("  mock      - Switch to the mock client")
	fmt.Println("  help      - Show this help message")
	fmt.Println("  clear     - Clear the context window")
	fmt.Println("  exit or quit  - Exit the program")
	return ""
}

func (d *Dispatcher) getCurrentClient() client.ChatClient {
	return d.clients[d.activeClient]
}

func (d *Dispatcher) dispatch(input string) string {
	cmd := strings.ToLower(input)
	if handler, exists := d.commands[cmd]; exists {
		return handler(cmd)
	}
	return input
}

func (d *Dispatcher) buildPrompt(question string) string {
	d.context = append(d.context, question)
	return strings.Join(d.context, "\n")
}

func (d *Dispatcher) processResponse(response string) {
	d.context = append(d.context, response)
	println(response)
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

	return strings.TrimSpace(input)
}

func main() {
	dispatcher := NewDispatcher()
	ctx := context.Background()

	for {
		userQuestion := getUserInput("Ask a question (or type 'help' for a list of options): ")

		if processed := dispatcher.dispatch(userQuestion); processed == "" {
			continue
		}

		prompt := dispatcher.buildPrompt(userQuestion)

		print("> ")
		println(userQuestion)
		println()

		chatClient := dispatcher.getCurrentClient()
		response, err := chatClient.Chat(ctx, prompt)
		if err != nil {
			log.Printf("Error getting completion: %v", err)
			continue
		}

		dispatcher.processResponse(response)
	}
}
