package model

// One value from a specific source
type MetricSource struct {
	Value  *float64 `json:"value"`  // numeric value, nil if unavailable
	Date   string   `json:"date"`   // ISO date (yyyy-mm-dd) of the underlying observation
	Source string   `json:"source"` // short id or full URL
}

type MetricStatus string

const (
	StatusExcellent MetricStatus = "excellent" // лучше нормы, устойчивый рост
	StatusGood      MetricStatus = "good"      // в пределах нормы
	StatusModerate  MetricStatus = "moderate"  // лёгкое ухудшение
	StatusWarning   MetricStatus = "warning"   // заметное ухудшение
	StatusDanger    MetricStatus = "danger"    // серьёзная угроза
	StatusCritical  MetricStatus = "critical"  // кризисный уровень
	StatusUnknown   MetricStatus = "unknown"   // нет данных / не рассчитано
)

// Aggregated metric for daily.json: multiple sources + average
type MetricDaily struct {
	Sources []MetricSource `json:"sources"`           // multiple sources per metric (can be 1..N)
	Average *float64       `json:"average"`           // precomputed simple average across sources; nil if no valid values
	Status  MetricStatus   `json:"status"`            // ok | warning | danger | unknown (filled later by thresholds)
	Comment string         `json:"comment,omitempty"` // optional note
}

// Country-level metrics (keys in daily.json)
type CountryMetrics struct {
	PolicyRate   MetricDaily `json:"policy_rate"`    // policy/benchmark interest rate
	Inflation    MetricDaily `json:"inflation"`      // CPI YoY %
	Unemployment MetricDaily `json:"unemployment"`   // unemployment rate %
	PMI          MetricDaily `json:"pmi"`            // PMI (manufacturing/services if you split later)
	EquityIndex  MetricDaily `json:"equity_index"`   // main equity index level
	FxRate       MetricDaily `json:"fx_rate"`        // FX to USD (units per USD or USD per unit — decide globally)
	BondYield10Y MetricDaily `json:"bond_yield_10y"` // 10-year gov bond yield %
}

// Daily payload written to output/YYYY-MM-DD.json
type Summary struct {
	Text string `json:"text"`    // LLM summary
	Tip  string `json:"llm_tip"` // LLM "tip of the day"
}

type DailyData struct {
	Date      string                    `json:"date"`      // run date (yyyy-mm-dd)
	Countries map[string]CountryMetrics `json:"countries"` // country code -> metrics
	Summary   Summary                   `json:"summary"`
}

// Trend payload for a single metric (kept for up to 180 points)
type TrendPoint struct {
	Date  string  `json:"date"`  // ISO date (yyyy-mm-dd)
	Value float64 `json:"value"` // numeric value
}

// LLM response schema (unchanged)
type StructuredLLMResponse struct {
	Summary string `json:"summary"`
	Tip     string `json:"tip"`
}

// Helpers
func Avg(vals []float64) *float64 {
	if len(vals) == 0 {
		return nil
	}
	sum := 0.0
	for _, v := range vals {
		sum += v
	}
	avg := sum / float64(len(vals))
	return &avg
}

func Ptr(f float64) *float64 { return &f }
