package llm

import (
	"context"

	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
)

type LLMClient interface {
	GenerateSummary(ctx context.Context, data schema.DailyData) (schema.StructuredLLMResponse, error)
}
