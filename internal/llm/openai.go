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

	"github.com/AndrewMysliuk/expath-data-generator/internal/llm/prompts"
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema"
	// openai "github.com/sashabaranov/go-openai"
)

//go:embed json_schema/immigration_info.schema.json
var immigrationInfoSchema []byte

//go:embed json_schema/tax_info.schema.json
var taxInfoSchema []byte

//go:embed json_schema/finance_info.schema.json
var financeInfoSchema []byte

//go:embed json_schema/cost_of_living_info.schema.json
var costOfLivingInfoSchema []byte

//go:embed json_schema/quality_of_life_info.schema.json
var qualityOfLifeInfoSchema []byte

const (
	HTTP_TIMEOUT = 60 * time.Second
)

type openAIClient struct {
	// client *openai.Client
	apiKey string
}

func NewOpenAIClient(apiKey string) LLMClient {
	return &openAIClient{
		apiKey: apiKey,
	}
}

func (c *openAIClient) GetImmigration(ctx context.Context, country schema.CountryInfo) (*schema.ImmigrationInfo, error) {
	prompt := prompts.ImmigrationInfoPrompt(country)

	raw, err := c.CallWithSchema(ctx, prompt, immigrationInfoSchema)
	if err != nil {
		return nil, err
	}

	fmt.Printf("RAW RESPONSE:\n%s\n", string(raw))

	var apiErr schema.APIErrorResponse
	if err := json.Unmarshal(raw, &apiErr); err == nil && apiErr.Error != nil {
		return nil, apiErr.Error
	}

	var resp schema.OpenAIResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal OpenAIResponse: %w", err)
	}

	content, err := resp.FirstText()
	if err != nil {
		return nil, err
	}

	var wrapper struct {
		Immigration schema.ImmigrationInfo `json:"immigration"`
	}
	if err := json.Unmarshal([]byte(content), &wrapper); err != nil {
		return nil, fmt.Errorf("failed to unmarshal immigration JSON: %w\nraw=%s", err, content)
	}

	return &wrapper.Immigration, nil
}

func (c *openAIClient) GetTaxes(ctx context.Context, country schema.CountryInfo) (*schema.TaxInfo, error) {
	prompt := prompts.TaxInfoPrompt(country)

	raw, err := c.CallWithSchema(ctx, prompt, taxInfoSchema)
	if err != nil {
		return nil, err
	}

	var apiErr schema.APIErrorResponse
	if err := json.Unmarshal(raw, &apiErr); err == nil && apiErr.Error != nil {
		return nil, apiErr.Error
	}

	var resp schema.OpenAIResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal OpenAIResponse: %w", err)
	}

	content, err := resp.FirstText()
	if err != nil {
		return nil, err
	}

	var wrapper struct {
		Taxes schema.TaxInfo `json:"taxes"`
	}
	if err := json.Unmarshal([]byte(content), &wrapper); err != nil {
		return nil, fmt.Errorf("unmarshal taxes: %w", err)
	}

	return &wrapper.Taxes, nil
}

func (c *openAIClient) GetFinance(ctx context.Context, country schema.CountryInfo) (*schema.FinanceInfo, error) {
	prompt := prompts.FinanceInfoPrompt(country)

	raw, err := c.CallWithSchema(ctx, prompt, financeInfoSchema)
	if err != nil {
		return nil, err
	}

	var apiErr schema.APIErrorResponse
	if err := json.Unmarshal(raw, &apiErr); err == nil && apiErr.Error != nil {
		return nil, apiErr.Error
	}

	var resp schema.OpenAIResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal OpenAI wrapper: %w", err)
	}

	content, err := resp.FirstText()
	if err != nil {
		return nil, err
	}

	var wrapper struct {
		Finance schema.FinanceInfo `json:"finance"`
	}
	if err := json.Unmarshal([]byte(content), &wrapper); err != nil {
		return nil, fmt.Errorf("unmarshal finance: %w", err)
	}

	return &wrapper.Finance, nil
}

func (c *openAIClient) GetCostOfLiving(ctx context.Context, country schema.CountryInfo) (*schema.CostOfLivingInfo, error) {
	prompt := prompts.CostOfLivingInfoPrompt(country)

	raw, err := c.CallWithSchema(ctx, prompt, costOfLivingInfoSchema)
	if err != nil {
		return nil, err
	}

	var apiErr schema.APIErrorResponse
	if err := json.Unmarshal(raw, &apiErr); err == nil && apiErr.Error != nil {
		return nil, apiErr.Error
	}

	var resp schema.OpenAIResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal OpenAI wrapper: %w", err)
	}

	content, err := resp.FirstText()
	if err != nil {
		return nil, err
	}

	var wrapper struct {
		CostOfLiving schema.CostOfLivingInfo `json:"cost_of_living"`
	}
	if err := json.Unmarshal([]byte(content), &wrapper); err != nil {
		return nil, fmt.Errorf("unmarshal cost_of_living: %w", err)
	}

	return &wrapper.CostOfLiving, nil
}

func (c *openAIClient) GetQualityOfLife(ctx context.Context, country schema.CountryInfo) (*schema.QualityOfLifeInfo, error) {
	prompt := prompts.QualityOfLifeInfoPrompt(country)

	raw, err := c.CallWithSchema(ctx, prompt, qualityOfLifeInfoSchema)
	if err != nil {
		return nil, err
	}

	var apiErr schema.APIErrorResponse
	if err := json.Unmarshal(raw, &apiErr); err == nil && apiErr.Error != nil {
		return nil, apiErr.Error
	}

	var resp schema.OpenAIResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal OpenAI wrapper: %w", err)
	}

	content, err := resp.FirstText()
	if err != nil {
		return nil, err
	}

	var wrapper struct {
		QualityOfLife schema.QualityOfLifeInfo `json:"quality_of_life"`
	}
	if err := json.Unmarshal([]byte(content), &wrapper); err != nil {
		return nil, fmt.Errorf("unmarshal quality_of_life: %w", err)
	}

	return &wrapper.QualityOfLife, nil
}

func (c *openAIClient) CallWithSchema(ctx context.Context, query string, schema []byte) (json.RawMessage, error) {
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
				"name":        "JsonSchemaResponse",
				"description": "Structured response with summary of search results",
				"schema":      json.RawMessage(schema),
				"strict":      true,
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{Timeout: HTTP_TIMEOUT}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	return json.RawMessage(raw), nil
}
