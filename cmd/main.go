package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/AndrewMysliuk/expath-data-generator/internal/llm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	openaiKey := os.Getenv("OPENAI_API_KEY")
	if openaiKey == "" {
		log.Fatal("OPENAI_API_KEY not set")
	}

	llmClient := llm.NewOpenAIClient(openaiKey)

	_ = llmClient
}
