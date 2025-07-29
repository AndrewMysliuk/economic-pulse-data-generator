package model

type Metric struct {
	Value   *float64 `json:"value"`             // Actual value, or nil if unavailable
	Status  string   `json:"status"`            // one of: "ok", "warning", "danger", "unknown"
	Source  string   `json:"source"`            // Data source URL
	Comment string   `json:"comment,omitempty"` // Optional explanation or interpretation
}

type CountryMetrics struct {
	PolicyRate   Metric `json:"policy_rate"`    // Interest rate
	Inflation    Metric `json:"inflation"`      // CPI
	Unemployment Metric `json:"unemployment"`   // Unemployment rate
	PMI          Metric `json:"pmi"`            // Purchasing Managers' Index
	EquityIndex  Metric `json:"equity_index"`   // Stock market index
	FxRate       Metric `json:"fx_rate"`        // Currency exchange rate to USD
	BondYield10Y Metric `json:"bond_yield_10y"` // 10-year government bond yield
}

type Summary struct {
	Text string `json:"text"`    // LLM summary
	Tip  string `json:"llm_tip"` // Recommendation or insight from LLM
}

type DailyData struct {
	Date      string                    `json:"date"`      // e.g. "2025-07-29"
	Countries map[string]CountryMetrics `json:"countries"` // Key: country name in English
	Summary   Summary                   `json:"summary"`
}

type StructuredLLMResponse struct {
	Summary string `json:"summary"`
	Tip     string `json:"tip"`
}

func ToFloat64(v float64) *float64 {
	return &v
}
