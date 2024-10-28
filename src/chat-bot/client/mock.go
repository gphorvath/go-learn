package client

import (
	"context"
	"fmt"
)

type MockClient struct {
	Responses []string
	CallCount int
}

func NewMock(responses []string) *MockClient {
	return &MockClient{
		Responses: responses,
	}
}

func (c *MockClient) Chat(ctx context.Context, prompt string) (string, error) {
	if c.CallCount >= len(c.Responses) {
		return "", fmt.Errorf("no more mock responses available")
	}
	response := c.Responses[c.CallCount]
	c.CallCount++
	return response, nil
}
