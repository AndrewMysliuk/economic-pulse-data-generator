package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/core"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/llm"
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

	err = core.Generate(llmClient)
	if err != nil {
		log.Fatalf("generation failed: %v", err)
	}

	// if err := scraper.Scrape("US"); err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("Generation complete")
}
