package llm

import (
	"context"

	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
)

type LLMClient interface {
	GenerateCountryMetrics(ctx context.Context, countryISO string, date string) (schema.CountryMetrics, error)
	GenerateSummary(ctx context.Context, data schema.DailyData) (schema.StructuredLLMResponse, error)
}
