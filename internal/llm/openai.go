package llm

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/model"
	openai "github.com/sashabaranov/go-openai"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func (c *openAIClient) GenerateSummary(data model.DailyData) (string, string, error) {
	ctx := context.Background()

	userMessage := buildUserMessage(data)

	schema := json.RawMessage(summarySchema)

	resp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "system",
					Content: systemPrompt(),
				},
				{
					Role:    "user",
					Content: userMessage,
				},
			},
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
				JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
					Name:        "EconomicSummary",
					Description: "JSON response with 'summary' and 'tip' fields",
					Schema:      schema,
					Strict:      true,
				},
			},
		},
	)
	if err != nil {
		return "", "", err
	}

	var parsed model.StructuredLLMResponse
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &parsed); err != nil {
		return "", "", fmt.Errorf("invalid JSON from LLM: %w\nRaw: %s", err, resp.Choices[0].Message.Content)
	}

	return parsed.Summary, parsed.Tip, nil
}

func systemPrompt() string {
	return "You are a professional macroeconomist. Write clear, concise summaries of global economic trends. Avoid buzzwords. If data is inconclusive, say so."
}

func buildUserMessage(data model.DailyData) string {
	var b strings.Builder

	caser := cases.Title(language.English)
	b.WriteString("Macroeconomic indicators as of " + data.Date + ":\n\n")

	for country, metrics := range data.Countries {
		b.WriteString("Country: " + caser.String(country) + "\n")
		b.WriteString("- Policy Rate: " + floatOrNA(metrics.PolicyRate.Value) + "\n")
		b.WriteString("- Inflation: " + floatOrNA(metrics.Inflation.Value) + "\n")
		b.WriteString("- Unemployment: " + floatOrNA(metrics.Unemployment.Value) + "\n")
		b.WriteString("- PMI: " + floatOrNA(metrics.PMI.Value) + "\n")
		b.WriteString("- Equity Index: " + floatOrNA(metrics.EquityIndex.Value) + "\n")
		b.WriteString("- FX Rate to USD: " + floatOrNA(metrics.FxRate.Value) + "\n")
		b.WriteString("- 10Y Bond Yield: " + floatOrNA(metrics.BondYield10Y.Value) + "\n\n")
	}

	b.WriteString("Please provide a concise summary of the global economic condition and one actionable insight for investors.")
	return b.String()
}

func floatOrNA(v *float64) string {
	if v == nil {
		return "N/A"
	}
	return formatFloat(*v)
}

func formatFloat(f float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", f), "0"), ".")
}
