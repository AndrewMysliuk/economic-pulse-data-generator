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

		sb.WriteString(fmt.Sprintf("  Policy Rate: %s\n", metricValue(metrics.MacroPolicyRatePercent)))
		sb.WriteString(fmt.Sprintf("  Inflation (CPI YoY): %s\n", metricValue(metrics.MacroInflationCPIYoYPercent)))
		sb.WriteString(fmt.Sprintf("  Unemployment: %s\n", metricValue(metrics.MacroUnemploymentRatePercent)))
		sb.WriteString(fmt.Sprintf("  PMI: %s\n", metricValue(metrics.MacroPMIIndex)))
		sb.WriteString(fmt.Sprintf("  Bond Yield 10Y: %s\n", metricValue(metrics.MacroBondYield10YPercent)))
		sb.WriteString(fmt.Sprintf("  GDP Growth Forecast: %s\n", metricValue(metrics.MacroGDPGrowthForecastPercent)))

		sb.WriteString(fmt.Sprintf("  Avg Net Salary: %s\n", metricValue(metrics.IncomeAverageNetMonthlyEUR)))
		sb.WriteString(fmt.Sprintf("  Living Wage: %s\n", metricValue(metrics.IncomeLivingWageEstimateEUR)))

		sb.WriteString(fmt.Sprintf("  Price Capital: %s\n", metricValue(metrics.RealEstatePriceCapitalUSDPerM2)))
		sb.WriteString(fmt.Sprintf("  Price Regional: %s\n", metricValue(metrics.RealEstatePriceRegionalUSDPerM2)))
		sb.WriteString(fmt.Sprintf("  Price Change YoY: %s\n", metricValue(metrics.RealEstatePriceChangeYoYPercent)))
		sb.WriteString(fmt.Sprintf("  Rental Yield: %s\n", metricValue(metrics.RealEstateRentalYieldPercent)))

		sb.WriteString(fmt.Sprintf("  Share Manufacturing: %s\n", metricValue(metrics.EconStructShareManufacturingPercent)))
		sb.WriteString(fmt.Sprintf("  Share Info & Financial: %s\n", metricValue(metrics.EconStructShareInfoFinancialServicesPercent)))
		sb.WriteString(fmt.Sprintf("  Share Trade & Logistics: %s\n", metricValue(metrics.EconStructShareTradeLogisticsPercent)))
		sb.WriteString(fmt.Sprintf("  Share Other Sectors: %s\n", metricValue(metrics.EconStructShareOtherSectorsPercent)))

		sb.WriteString(fmt.Sprintf("  Population: %s\n", metricValue(metrics.SocietyPopulationMillion)))
		sb.WriteString(fmt.Sprintf("  Birth Rate: %s\n", metricValue(metrics.SocietyBirthRatePerWoman)))
		sb.WriteString(fmt.Sprintf("  Corruption Index: %s\n", metricValue(metrics.SocietyCorruptionIndex100Scale)))
		sb.WriteString(fmt.Sprintf("  Political Stability: %s\n", metricValue(metrics.SocietyPoliticalStabilityRating)))

		sb.WriteString("\n")
	}

	sb.WriteString("Generate a concise summary of the global economic situation and a short actionable tip based on this data.")
	return sb.String()
}

func metricValue(m schema.MetricDaily) string {
	src := m.Sources

	switch {
	case src.Value != nil:
		return fmt.Sprintf("%.2f %s", *src.Value, src.Unit)
	case src.From != nil && src.To != nil:
		return fmt.Sprintf("%.2f–%.2f %s", *src.From, *src.To, src.Unit)
	case src.Raw != "":
		return src.Raw
	default:
		return "N/A"
	}
}
