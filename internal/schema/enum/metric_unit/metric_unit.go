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
