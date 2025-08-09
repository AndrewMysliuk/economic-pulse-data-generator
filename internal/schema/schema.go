package schema

import (
	"regexp"

	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema/enum/metric_status"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema/enum/metric_unit"
	"github.com/pkg/errors"
)

var (
	FX_PAIR_REGEXP = regexp.MustCompile(`^[A-Z]{3}/[A-Z]{3}$`)
)

type MetricSource struct {
	Value      *float64               `json:"value"`       // число; nil если не достали/невалидно
	Date       string                 `json:"date"`        // YYYY-MM-DD — дата наблюдения
	Unit       metric_unit.MetricUnit `json:"unit"`        // "%", "index", "points" и т.п.
	SourceUrl  string                 `json:"source_url"`  // первоисточник (полный URL или короткий id)
	SourceName string                 `json:"source_name"` // человекочитаемое имя источника/серии
}

type MetricDaily struct {
	Sources []MetricSource             `json:"sources"`           // 1..N источников
	Average *float64                   `json:"average,omitempty"` // простой avg; опционально (если 1 источник — можно не ставить)
	Unit    metric_unit.MetricUnit     `json:"unit"`              // итоговая единица для UI (должна совпадать с источниками)
	Status  metric_status.MetricStatus `json:"status"`            // выставляется твоей логикой порогов
	AsOf    string                     `json:"as_of,omitempty"`   // ISO datetime (UTC), актуально для рынков/дневных серий
}

type FxRate struct {
	Pair       string  `json:"pair"`        // конвенция: "USD/EUR" = EUR за 1 USD
	Value      float64 `json:"value"`       // EUR per 1 USD
	AsOf       string  `json:"as_of"`       // ISO datetime (UTC)
	SourceUrl  string  `json:"source_url"`  // источник (Fed H.10, ECB и т.п.)
	SourceName string  `json:"source_name"` // человекочитаемое имя источника
}

type CountryMetrics struct {
	PolicyRate    MetricDaily `json:"policy_rate"`    // ключевая/бенчмарк ставка
	Inflation     MetricDaily `json:"inflation"`      // headline CPI YoY %
	Unemployment  MetricDaily `json:"unemployment"`   // U-3 SA % или местный аналог
	PMI           MetricDaily `json:"pmi"`            // Composite по умолчанию (тип можно уточнять в SourceName)
	EquityIndex   MetricDaily `json:"equity_index"`   // основной фондовый индекс страны (уровень)
	FxRates       []FxRate    `json:"fx_rates"`       // FX пары по конвенции выше
	CurrencyIndex MetricDaily `json:"currency_index"` // если неприменимо — держи status=unknown, sources=[]
	BondYield10Y  MetricDaily `json:"bond_yield_10y"` // 10Y gov yield, % годовых
}

type StructuredLLMResponse struct {
	Text string `json:"summary"` // краткая сводка дня
	Tip  string `json:"tip"`     // короткий actionable инсайт
}

type DailyData struct {
	Date      string                    `json:"date"`      // дата запуска пайплайна (YYYY-MM-DD)
	Countries map[string]CountryMetrics `json:"countries"` // ISO-2 коды стран
	Summary   StructuredLLMResponse     `json:"summary"`   // общая сводка
}

func (m MetricDaily) ValidateUnits() error {
	for i := range m.Sources {
		if m.Sources[i].Unit != m.Unit {
			return errors.Errorf("unit mismatch: daily=%s source[%d]=%s", m.Unit, i, m.Sources[i].Unit)
		}
	}
	return nil
}

func (m *MetricDaily) ComputeAverage() {
	valid := make([]float64, 0, len(m.Sources))
	for _, s := range m.Sources {
		if s.Value != nil {
			valid = append(valid, *s.Value)
		}
	}
	if len(valid) == 0 {
		m.Average = nil
		return
	}
	sum := 0.0
	for _, v := range valid {
		sum += v
	}
	avg := sum / float64(len(valid))
	m.Average = &avg
}

func (fx FxRate) Validate() error {
	if !FX_PAIR_REGEXP.MatchString(fx.Pair) {
		return errors.Errorf("invalid pair: %s", fx.Pair)
	}
	if fx.SourceUrl == "" || fx.SourceName == "" {
		return errors.New("source_url/source_name required")
	}
	return nil
}

func PtrVal(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}
