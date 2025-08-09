package llm

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"

	country_metrics_prompt "github.com/AndrewMysliuk/economic-pulse-data-generator/internal/llm/prompts/country_metrics"
	summary_prompt "github.com/AndrewMysliuk/economic-pulse-data-generator/internal/llm/prompts/summary"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
	openai "github.com/sashabaranov/go-openai"
)

//go:embed json_schema/country_metrics.schema.json
var countryMetricsSchema []byte

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

func (c *openAIClient) GenerateCountryMetrics(ctx context.Context, countryISO string, date string) (schema.CountryMetrics, error) {
	userMsg := country_metrics_prompt.BuildUserMessage(countryISO, date)
	schemaDef := json.RawMessage(countryMetricsSchema)

	apiResp, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: country_metrics_prompt.SystemPrompt()},
			{Role: "user", Content: userMsg},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
				Name:        "CountryMetrics",
				Description: "Country-level macro metrics with sources and units",
				Schema:      schemaDef,
				Strict:      true,
			},
		},
	})
	if err != nil {
		return schema.CountryMetrics{}, err
	}
	if len(apiResp.Choices) == 0 {
		return schema.CountryMetrics{}, fmt.Errorf("empty choices from LLM")
	}

	var parsed schema.CountryMetrics
	if err := json.Unmarshal([]byte(apiResp.Choices[0].Message.Content), &parsed); err != nil {
		return schema.CountryMetrics{},
			fmt.Errorf("invalid JSON from LLM: %w\nRaw: %s", err, apiResp.Choices[0].Message.Content)
	}

	return parsed, nil
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
