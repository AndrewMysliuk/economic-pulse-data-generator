package core

import "github.com/AndrewMysliuk/economic-pulse-data-generator/internal/model"

// mockData builds a DailyData struct with hardcoded example values.
// Each metric uses MetricDaily with one source and status from the new enum.
func mockData(date string) model.DailyData {
	// helper to create a MetricDaily with one source
	oneSource := func(value float64, status model.MetricStatus, source string) model.MetricDaily {
		src := model.MetricSource{
			Value:  model.Ptr(value),
			Date:   date,
			Source: source,
		}
		return model.MetricDaily{
			Sources: []model.MetricSource{src},
			Average: model.Ptr(value),
			Status:  status,
		}
	}

	return model.DailyData{
		Date: date,
		Countries: map[string]model.CountryMetrics{
			"usa": {
				PolicyRate:   oneSource(5.25, model.StatusWarning, "mock"),
				Inflation:    oneSource(3.2, model.StatusModerate, "mock"),
				Unemployment: oneSource(3.6, model.StatusGood, "mock"),
				PMI:          oneSource(49.8, model.StatusModerate, "mock"),
				EquityIndex:  oneSource(4700, model.StatusGood, "mock"),
				FxRate:       oneSource(1.0, model.StatusGood, "mock"),
				BondYield10Y: oneSource(4.2, model.StatusWarning, "mock"),
			},
			"germany": {
				PolicyRate:   oneSource(4.0, model.StatusGood, "mock"),
				Inflation:    oneSource(2.6, model.StatusGood, "mock"),
				Unemployment: oneSource(5.1, model.StatusModerate, "mock"),
				PMI:          oneSource(46.0, model.StatusWarning, "mock"),
				EquityIndex:  oneSource(15800, model.StatusModerate, "mock"),
				FxRate:       oneSource(1.1, model.StatusGood, "mock"),
				BondYield10Y: oneSource(2.3, model.StatusGood, "mock"),
			},
			"china": {
				PolicyRate:   oneSource(3.45, model.StatusModerate, "mock"),
				Inflation:    oneSource(0.2, model.StatusDanger, "mock"),
				Unemployment: oneSource(5.2, model.StatusModerate, "mock"),
				PMI:          oneSource(50.5, model.StatusGood, "mock"),
				EquityIndex:  oneSource(3300, model.StatusModerate, "mock"),
				FxRate:       oneSource(7.2, model.StatusWarning, "mock"),
				BondYield10Y: oneSource(2.6, model.StatusGood, "mock"),
			},
			"japan": {
				PolicyRate:   oneSource(0.1, model.StatusDanger, "mock"),
				Inflation:    oneSource(3.0, model.StatusGood, "mock"),
				Unemployment: oneSource(2.6, model.StatusExcellent, "mock"),
				PMI:          oneSource(49.2, model.StatusModerate, "mock"),
				EquityIndex:  oneSource(39000, model.StatusGood, "mock"),
				FxRate:       oneSource(156.8, model.StatusDanger, "mock"),
				BondYield10Y: oneSource(0.7, model.StatusDanger, "mock"),
			},
			"uk": {
				PolicyRate:   oneSource(5.25, model.StatusWarning, "mock"),
				Inflation:    oneSource(2.9, model.StatusGood, "mock"),
				Unemployment: oneSource(4.2, model.StatusGood, "mock"),
				PMI:          oneSource(51.0, model.StatusGood, "mock"),
				EquityIndex:  oneSource(7900, model.StatusModerate, "mock"),
				FxRate:       oneSource(1.29, model.StatusGood, "mock"),
				BondYield10Y: oneSource(4.1, model.StatusWarning, "mock"),
			},
			"france": {
				PolicyRate:   oneSource(4.0, model.StatusGood, "mock"),
				Inflation:    oneSource(2.4, model.StatusExcellent, "mock"),
				Unemployment: oneSource(7.1, model.StatusWarning, "mock"),
				PMI:          oneSource(47.3, model.StatusWarning, "mock"),
				EquityIndex:  oneSource(7500, model.StatusModerate, "mock"),
				FxRate:       oneSource(1.1, model.StatusGood, "mock"),
				BondYield10Y: oneSource(2.8, model.StatusGood, "mock"),
			},
			"india": {
				PolicyRate:   oneSource(6.5, model.StatusDanger, "mock"),
				Inflation:    oneSource(4.6, model.StatusWarning, "mock"),
				Unemployment: oneSource(7.0, model.StatusWarning, "mock"),
				PMI:          oneSource(57.5, model.StatusExcellent, "mock"),
				EquityIndex:  oneSource(74000, model.StatusExcellent, "mock"),
				FxRate:       oneSource(83.1, model.StatusDanger, "mock"),
				BondYield10Y: oneSource(7.1, model.StatusDanger, "mock"),
			},
			"brazil": {
				PolicyRate:   oneSource(10.5, model.StatusCritical, "mock"),
				Inflation:    oneSource(3.9, model.StatusModerate, "mock"),
				Unemployment: oneSource(7.8, model.StatusWarning, "mock"),
				PMI:          oneSource(50.2, model.StatusGood, "mock"),
				EquityIndex:  oneSource(126000, model.StatusGood, "mock"),
				FxRate:       oneSource(5.3, model.StatusModerate, "mock"),
				BondYield10Y: oneSource(11.2, model.StatusCritical, "mock"),
			},
		},
	}
}
