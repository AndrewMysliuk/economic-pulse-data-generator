package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/AndrewMysliuk/expath-data-generator/internal/core"
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

	ctx := context.Background()
	coreApp := core.NewCore(llmClient)

	if err := coreApp.Run(ctx, "output"); err != nil {
		fmt.Printf("error: %+v\n", err)
	}
}
