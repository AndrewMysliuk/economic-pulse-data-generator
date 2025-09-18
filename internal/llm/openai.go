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

	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
	openai "github.com/sashabaranov/go-openai"
)

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
