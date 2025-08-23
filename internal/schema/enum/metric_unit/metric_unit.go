package metric_unit

import "github.com/pkg/errors"

type MetricUnit string

func (enum MetricUnit) Equals(str string) bool {
	return string(enum) == str
}

func (enum MetricUnit) Pointer() *MetricUnit {
	return &enum
}

func (enum MetricUnit) String() string {
	return string(enum)
}

const (
	Percent MetricUnit = "PERCENT"  // "%": инфляция, безработица, доходности
	Index   MetricUnit = "INDEX"    // индексное значение: PMI, фондовые индексы, валютные индексы
	Points  MetricUnit = "POINTS"   // пункты: например, абсолютные изменения индексов
	Level   MetricUnit = "LEVEL"    // уровень (без единиц), например курс валюты или числовое значение индекса без % или шкалы
	RatePct MetricUnit = "RATE_PCT" // процентная ставка в %
)

func New(str string) (MetricUnit, error) {
	switch str {

	case "PERCENT":
		return Percent, nil

	case "INDEX":
		return Index, nil

	case "POINTS":
		return Points, nil

	case "LEVEL":
		return Level, nil

	case "RATE_PCT":
		return RatePct, nil

	default:
		return "", errors.Errorf("can't define MetricUnit from: %v available values are: %v", str, Values())
	}
}

func Values() []MetricUnit {
	return []MetricUnit{Percent, Index, Points, Level, RatePct}
}

func ResolveUnit(name string) MetricUnit {
	switch name {
	case "PolicyRate":
		return RatePct
	case "Inflation", "Unemployment", "Bond10Y", "GDP", "PriceChangeYoY", "RentalYield",
		"ShareManufacturing", "ShareFinance", "ShareLogistics", "ShareOther":
		return Percent
	case "PMI":
		return Index
	case "NetSalary", "LivingWage":
		return Level // может сделать EUR_PER_MONTH отдельно, если критично
	case "PriceCapital", "PriceRegional":
		return Level // или завести "USD_PER_M2" кастомно
	case "Population":
		return Level // в миллионах — либо завести MILLION
	case "BirthRate":
		return Level // per woman
	case "Corruption":
		return Index
	case "PoliticalStability":
		return Level // или завести TEXT_SCORE
	default:
		return Level
	}
}
