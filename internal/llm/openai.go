package llm

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"

	summary_prompt "github.com/AndrewMysliuk/economic-pulse-data-generator/internal/llm/prompts/summary"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
	openai "github.com/sashabaranov/go-openai"
)

//go:embed json_schema/summary.schema.json
var summarySchema []byte

type openAIClient struct {
	client *openai.Client
}

func NewOpenAIClient(apiKey string) LLMClient {
	return &openAIClient{
		client: openai.NewClient(apiKey),
	}
}

func (c *openAIClient) GenerateSummary(ctx context.Context, data schema.DailyData) (resp schema.StructuredLLMResponse, err error) {
	userMessage := summary_prompt.BuildUserMessage(data)
	schemaDef := json.RawMessage(summarySchema)

	apiResp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: summary_prompt.SystemPrompt()},
				{Role: "user", Content: userMessage},
			},
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
				JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
					Name:        "EconomicSummary",
					Description: "JSON response with 'summary' and 'tip' fields",
					Schema:      schemaDef,
					Strict:      true,
				},
			},
		},
	)
	if err != nil {
		return schema.StructuredLLMResponse{}, err
	}

	var parsed schema.StructuredLLMResponse
	if err := json.Unmarshal([]byte(apiResp.Choices[0].Message.Content), &parsed); err != nil {
		return schema.StructuredLLMResponse{},
			fmt.Errorf("invalid JSON from LLM: %w\nRaw: %s", err, apiResp.Choices[0].Message.Content)
	}

	return parsed, nil
}
