package llm

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	summary_prompt "github.com/AndrewMysliuk/economic-pulse-data-generator/internal/llm/prompts/summary"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
	openai "github.com/sashabaranov/go-openai"
)

//go:embed json_schema/summary.schema.json
var summarySchema []byte

//go:embed json_schema/search_and_summarize.schema.json
var searchAndSummarizeSchema []byte

type openAIClient struct {
	client *openai.Client
	apiKey string
}

func NewOpenAIClient(apiKey string) LLMClient {
	return &openAIClient{
		client: openai.NewClient(apiKey),
		apiKey: apiKey,
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

func (c *openAIClient) SearchAndSummarize(ctx context.Context, query string) (string, error) {
	url := "https://api.openai.com/v1/responses"

	payload := map[string]any{
		"model": "gpt-4.1",
		"input": query,
		"tools": []map[string]any{
			{"type": "web_search"},
		},
		"text": map[string]any{
			"format": map[string]any{
				"type":        "json_schema",
				"name":        "SearchAndSummarize",
				"description": "Structured response with summary of search results",
				"schema":      json.RawMessage(searchAndSummarizeSchema),
				"strict":      true,
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{Timeout: 60 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read body: %w", err)
	}

	// If Error Occurs
	// fmt.Println("RAW:", string(raw))

	var parsed schema.SearchAndSummarizeResponse
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return "", fmt.Errorf("invalid JSON: %w\nRaw: %s", err, string(raw))
	}

	var output string
	for _, out := range parsed.Output {
		for _, c := range out.Content {
			if c.Type == "output_text" {
				output += c.Text + "\n"
			}
		}
	}

	return output, nil
}
