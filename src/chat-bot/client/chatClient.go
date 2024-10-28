package client

import "context"

// ChatClient defines the interface for any chat-based LLM client
type ChatClient interface {
	Chat(ctx context.Context, prompt string) (string, error)
}
