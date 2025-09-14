package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

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

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := llmClient.SearchAndSummarize(ctx, "Актуальные новости Германии по экономике, 3 штуки, укажи дату каждой новости")
	if err != nil {
		log.Fatalf("web search failed: %v", err)
	}

	fmt.Println(result)

	// err = core.Generate(llmClient)
	// if err != nil {
	// 	log.Fatalf("generation failed: %v", err)
	// }

	// if err := scraper.Scrape("US"); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Generation complete")
}
