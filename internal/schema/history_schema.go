package schema

import (
	"encoding/json"
	"os"
	"sort"
	"time"
)

type MetricPoint struct {
	Date  string   `json:"date"`  // YYYY-MM-DD
	Value *float64 `json:"value"` // null допустим
}

type FxPoint struct {
	Date  string  `json:"date"`  // YYYY-MM-DD
	Value float64 `json:"value"` // YYY per 1 XXX по твоей конвенции
}

type CountryHistory struct {
	Units         map[string]string    `json:"units,omitempty"`
	PolicyRate    []MetricPoint        `json:"policy_rate,omitempty"`
	Inflation     []MetricPoint        `json:"inflation,omitempty"`
	Unemployment  []MetricPoint        `json:"unemployment,omitempty"`
	PMI           []MetricPoint        `json:"pmi,omitempty"`
	EquityIndex   []MetricPoint        `json:"equity_index,omitempty"`
	CurrencyIndex []MetricPoint        `json:"currency_index,omitempty"`
	BondYield10Y  []MetricPoint        `json:"bond_yield_10y,omitempty"`
	FxRates       map[string][]FxPoint `json:"fx_rates,omitempty"`
}

type History struct {
	GeneratedAt string                     `json:"generated_at"` // RFC3339 UTC
	WindowDays  int                        `json:"window_days"`  // напр., 180
	StartDate   string                     `json:"start_date"`   // YYYY-MM-DD
	EndDate     string                     `json:"end_date"`     // YYYY-MM-DD
	Countries   map[string]*CountryHistory `json:"countries"`    // ISO-2 -> серии
}

func NewHistory(windowDays int) History {
	return History{
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		WindowDays:  windowDays,
		Countries:   map[string]*CountryHistory{},
	}
}

func (h *History) ensureCountry(iso string) *CountryHistory {
	ch := h.Countries[iso]
	if ch == nil {
		ch = &CountryHistory{
			Units:   map[string]string{},
			FxRates: map[string][]FxPoint{},
		}
		h.Countries[iso] = ch
	}
	return ch
}

func (h *History) AppendDay(day DailyData) {
	dayDate := day.Date // YYYY-MM-DD

	for iso, cm := range day.Countries {
		ch := h.ensureCountry(iso)

		if cm.PolicyRate.Unit != "" {
			ch.Units["policy_rate"] = string(cm.PolicyRate.Unit)
		}
		if cm.Inflation.Unit != "" {
			ch.Units["inflation"] = string(cm.Inflation.Unit)
		}
		if cm.Unemployment.Unit != "" {
			ch.Units["unemployment"] = string(cm.Unemployment.Unit)
		}
		if cm.PMI.Unit != "" {
			ch.Units["pmi"] = string(cm.PMI.Unit)
		}
		if cm.EquityIndex.Unit != "" {
			ch.Units["equity_index"] = string(cm.EquityIndex.Unit)
		}
		if cm.CurrencyIndex.Unit != "" {
			ch.Units["currency_index"] = string(cm.CurrencyIndex.Unit)
		}
		if cm.BondYield10Y.Unit != "" {
			ch.Units["bond_yield_10y"] = string(cm.BondYield10Y.Unit)
		}

		ch.PolicyRate = append(ch.PolicyRate, MetricPoint{Date: dayDate, Value: cm.PolicyRate.Average})
		ch.Inflation = append(ch.Inflation, MetricPoint{Date: dayDate, Value: cm.Inflation.Average})
		ch.Unemployment = append(ch.Unemployment, MetricPoint{Date: dayDate, Value: cm.Unemployment.Average})
		ch.PMI = append(ch.PMI, MetricPoint{Date: dayDate, Value: cm.PMI.Average})
		ch.EquityIndex = append(ch.EquityIndex, MetricPoint{Date: dayDate, Value: cm.EquityIndex.Average})
		ch.CurrencyIndex = append(ch.CurrencyIndex, MetricPoint{Date: dayDate, Value: cm.CurrencyIndex.Average})
		ch.BondYield10Y = append(ch.BondYield10Y, MetricPoint{Date: dayDate, Value: cm.BondYield10Y.Average})

		for _, r := range cm.FxRates {
			asOf := r.AsOf
			if len(asOf) >= 10 {
				asOf = asOf[:10]
			} else {
				asOf = dayDate
			}
			ch.FxRates[r.Pair] = append(ch.FxRates[r.Pair], FxPoint{
				Date:  asOf,
				Value: r.Value,
			})
		}
	}

	h.GeneratedAt = time.Now().UTC().Format(time.RFC3339)
	h.recomputeDateRange()
}

func trimSlice[T any](s []T, n int) []T {
	if n <= 0 {
		return nil
	}
	if len(s) <= n {
		return s
	}
	return s[len(s)-n:]
}

func (h *History) TrimToWindow() {
	keep := h.WindowDays

	for _, ch := range h.Countries {
		sort.Slice(ch.PolicyRate, func(i, j int) bool { return ch.PolicyRate[i].Date < ch.PolicyRate[j].Date })
		sort.Slice(ch.Inflation, func(i, j int) bool { return ch.Inflation[i].Date < ch.Inflation[j].Date })
		sort.Slice(ch.Unemployment, func(i, j int) bool { return ch.Unemployment[i].Date < ch.Unemployment[j].Date })
		sort.Slice(ch.PMI, func(i, j int) bool { return ch.PMI[i].Date < ch.PMI[j].Date })
		sort.Slice(ch.EquityIndex, func(i, j int) bool { return ch.EquityIndex[i].Date < ch.EquityIndex[j].Date })
		sort.Slice(ch.CurrencyIndex, func(i, j int) bool { return ch.CurrencyIndex[i].Date < ch.CurrencyIndex[j].Date })
		sort.Slice(ch.BondYield10Y, func(i, j int) bool { return ch.BondYield10Y[i].Date < ch.BondYield10Y[j].Date })

		ch.PolicyRate = trimSlice(ch.PolicyRate, keep)
		ch.Inflation = trimSlice(ch.Inflation, keep)
		ch.Unemployment = trimSlice(ch.Unemployment, keep)
		ch.PMI = trimSlice(ch.PMI, keep)
		ch.EquityIndex = trimSlice(ch.EquityIndex, keep)
		ch.CurrencyIndex = trimSlice(ch.CurrencyIndex, keep)
		ch.BondYield10Y = trimSlice(ch.BondYield10Y, keep)

		for pair, s := range ch.FxRates {
			sort.Slice(s, func(i, j int) bool { return s[i].Date < s[j].Date })
			ch.FxRates[pair] = trimSlice(s, keep)
		}
	}

	h.recomputeDateRange()
}

func (h *History) recomputeDateRange() {
	maxDate := ""
	consider := func(mp []MetricPoint) {
		if len(mp) == 0 {
			return
		}
		d := mp[len(mp)-1].Date
		if d > maxDate {
			maxDate = d
		}
	}
	for _, ch := range h.Countries {
		consider(ch.PolicyRate)
		consider(ch.Inflation)
		consider(ch.Unemployment)
		consider(ch.PMI)
		consider(ch.EquityIndex)
		consider(ch.CurrencyIndex)
		consider(ch.BondYield10Y)
		for _, s := range ch.FxRates {
			if len(s) == 0 {
				continue
			}
			d := s[len(s)-1].Date
			if d > maxDate {
				maxDate = d
			}
		}
	}
	if maxDate == "" {
		h.StartDate, h.EndDate = "", ""
		return
	}
	h.EndDate = maxDate
	if t, err := time.Parse("2006-01-02", maxDate); err == nil {
		h.StartDate = t.AddDate(0, 0, -(h.WindowDays - 1)).Format("2006-01-02")
	}
}

func SaveHistory(h History, path string) error {
	b, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

func LoadHistory(path string) (History, error) {
	var h History
	b, err := os.ReadFile(path)
	if err != nil {
		return h, err
	}
	if err := json.Unmarshal(b, &h); err != nil {
		return h, err
	}
	return h, nil
}
