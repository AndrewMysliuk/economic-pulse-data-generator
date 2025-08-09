package summary

import (
	"fmt"
	"strings"

	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
)

func SystemPrompt() string {
	return `You are an economic data analyst. 
You receive fresh daily macroeconomic metrics for multiple countries.

Your task:
- Summarize the overall global and cross-country economic situation in a **concise and factual** way.
- Provide a short actionable tip or insight based on the data.
- Focus only on what is present in the input, do not hallucinate or invent numbers.
- Keep the summary no longer than 3–4 sentences, the tip no longer than 1 sentence.

Output format: strictly match the provided JSON Schema with:
{
  "summary": string, // concise overview
  "tip": string      // short actionable insight
}`
}

func BuildUserMessage(data schema.DailyData) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Date: %s\n", data.Date))
	sb.WriteString("Daily macroeconomic metrics by country:\n\n")

	for country, metrics := range data.Countries {
		sb.WriteString(fmt.Sprintf("Country: %s\n", country))
		sb.WriteString(fmt.Sprintf("  PolicyRate: %s\n", metricValue(metrics.PolicyRate)))
		sb.WriteString(fmt.Sprintf("  Inflation: %s\n", metricValue(metrics.Inflation)))
		sb.WriteString(fmt.Sprintf("  Unemployment: %s\n", metricValue(metrics.Unemployment)))
		sb.WriteString(fmt.Sprintf("  PMI: %s\n", metricValue(metrics.PMI)))
		sb.WriteString(fmt.Sprintf("  EquityIndex: %s\n", metricValue(metrics.EquityIndex)))
		sb.WriteString(fmt.Sprintf("  CurrencyIndex: %s\n", metricValue(metrics.CurrencyIndex)))
		sb.WriteString(fmt.Sprintf("  BondYield10Y: %s\n", metricValue(metrics.BondYield10Y)))

		if len(metrics.FxRates) > 0 {
			sb.WriteString("  FxRates:\n")
			for _, fx := range metrics.FxRates {
				sb.WriteString(fmt.Sprintf("    %s: %.4f (as of %s)\n", fx.Pair, fx.Value, fx.AsOf))
			}
		}
		sb.WriteString("\n")
	}

	sb.WriteString("\nGenerate a concise summary of the global economic situation and a short actionable tip based on this data.")
	return sb.String()
}

func metricValue(m schema.MetricDaily) string {
	if m.Average != nil {
		return fmt.Sprintf("%.2f %s", *m.Average, m.Unit)
	}
	return "N/A"
}
