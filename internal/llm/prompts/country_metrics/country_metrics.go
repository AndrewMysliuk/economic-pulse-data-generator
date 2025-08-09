package country_metrics

import "fmt"

func SystemPrompt() string {
	return `You are an economic data extraction AI. 
Your goal is to produce structured JSON describing today's key economic metrics for a given country.

Rules:
- Output MUST strictly follow the provided JSON Schema for CountryMetrics.
- Do NOT add extra fields or omit required ones.
- All numbers must be in the units specified by the "unit" field.
- Dates must be in YYYY-MM-DD for daily values, and RFC3339 UTC datetime for "as_of".
- "sources" must contain only real, verifiable sources; no fictional URLs or names.
- If a value is unknown, set "value" to null, but still provide "date", "unit", "source_url", and "source_name".
- If a metric is not applicable for this country, set its "status" to "UNKNOWN" and "sources" to an empty array.
- "fx_rates" must always use convention XXX/YYY meaning YYY per 1 XXX.
- Use realistic, recent economic values — do not guess unrealistic numbers.
- Use English names for sources, and full URLs where possible.

Your output will be parsed automatically; any deviation from schema will be rejected.`
}

func BuildUserMessage(countryISO, date string) string {
	return fmt.Sprintf(`Provide the latest available macroeconomic data for the country with ISO-2 code "%s".
The target reference date is %s (today's run date).
Return data in strict accordance with the provided CountryMetrics JSON Schema.

Metrics required:
- policy_rate: Key policy interest rate or benchmark rate.
- inflation: Headline CPI YoY in percent.
- unemployment: Unemployment rate (U-3 SA %% or local equivalent).
- pmi: PMI Composite index (or manufacturing if composite unavailable).
- equity_index: Main stock market index level for this country.
- fx_rates: Currency pairs against major currencies, convention XXX/YYY = YYY per 1 XXX.
- currency_index: Country-specific currency index (e.g., DXY for USD) if available.
- bond_yield_10y: 10-year government bond yield in percent.

Guidelines:
- Dates: Use YYYY-MM-DD for 'date', and RFC3339 UTC for 'as_of'.
- Units: Use enum values: PERCENT, INDEX, POINTS, LEVEL, RATE_PCT.
- Values: Use realistic, recent data only; if unknown, set value = null.
- Sources: Use trusted, verifiable public data sources (full URLs).
- fx_rates: Include USD, EUR, JPY, GBP, CNY, INR, BRL at minimum if available.
- If metric is not applicable, set status=UNKNOWN and sources=[].

Important: Output only valid JSON matching the schema. No extra commentary.`, countryISO, date)
}
