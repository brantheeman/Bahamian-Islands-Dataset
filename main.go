package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/ollama/ollama/api"
)

func main() {
	// Create Ollama client
	client := api.NewClient(&url.URL{
		Scheme: "http",
		Host:   "localhost:11434",
	}, &http.Client{})

	// Make a chat request
	req := &api.ChatRequest{
		Model: "llama3.1", // Make sure this matches your `ollama list`
		Messages: []api.Message{
			{
				Role:    "user",
				Content: "Hello, can you help me analyze JSON data about the Bahamas?",
			},
		},
	}

	ctx := context.Background()
	err := client.Chat(ctx, req, func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to chat with Ollama: %v", err)
	}

	fmt.Println() // Add newline at the end
}
