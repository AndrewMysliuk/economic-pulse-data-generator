package llm

import (
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/model"
)

type LLMClient interface {
	GenerateSummary(data model.DailyData) (summary string, tip string, err error)
}
