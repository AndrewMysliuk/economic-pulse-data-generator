package core

import "github.com/AndrewMysliuk/economic-pulse-data-generator/internal/model"

func mockData(date string) model.DailyData {
	return model.DailyData{
		Date: date,
		Countries: map[string]model.CountryMetrics{
			"usa": {
				PolicyRate:   model.Metric{Value: model.ToFloat64(5.25), Status: "high"},
				Inflation:    model.Metric{Value: model.ToFloat64(3.2), Status: "moderate"},
				Unemployment: model.Metric{Value: model.ToFloat64(3.6), Status: "normal"},
				PMI:          model.Metric{Value: model.ToFloat64(49.8), Status: "contracting"},
				EquityIndex:  model.Metric{Value: model.ToFloat64(4700), Status: "bullish"},
				FxRate:       model.Metric{Value: model.ToFloat64(1.0), Status: "base"},
				BondYield10Y: model.Metric{Value: model.ToFloat64(4.2), Status: "elevated"},
			},
			"germany": {
				PolicyRate:   model.Metric{Value: model.ToFloat64(4.0), Status: "normal"},
				Inflation:    model.Metric{Value: model.ToFloat64(2.6), Status: "low"},
				Unemployment: model.Metric{Value: model.ToFloat64(5.1), Status: "elevated"},
				PMI:          model.Metric{Value: model.ToFloat64(46.0), Status: "contracting"},
				EquityIndex:  model.Metric{Value: model.ToFloat64(15800), Status: "neutral"},
				FxRate:       model.Metric{Value: model.ToFloat64(1.1), Status: "stable"},
				BondYield10Y: model.Metric{Value: model.ToFloat64(2.3), Status: "low"},
			},
			"china": {
				PolicyRate:   model.Metric{Value: model.ToFloat64(3.45), Status: "loose"},
				Inflation:    model.Metric{Value: model.ToFloat64(0.2), Status: "deflation"},
				Unemployment: model.Metric{Value: model.ToFloat64(5.2), Status: "elevated"},
				PMI:          model.Metric{Value: model.ToFloat64(50.5), Status: "expanding"},
				EquityIndex:  model.Metric{Value: model.ToFloat64(3300), Status: "recovering"},
				FxRate:       model.Metric{Value: model.ToFloat64(7.2), Status: "weak"},
				BondYield10Y: model.Metric{Value: model.ToFloat64(2.6), Status: "low"},
			},
			"japan": {
				PolicyRate:   model.Metric{Value: model.ToFloat64(0.1), Status: "ultralow"},
				Inflation:    model.Metric{Value: model.ToFloat64(3.0), Status: "normal"},
				Unemployment: model.Metric{Value: model.ToFloat64(2.6), Status: "low"},
				PMI:          model.Metric{Value: model.ToFloat64(49.2), Status: "contracting"},
				EquityIndex:  model.Metric{Value: model.ToFloat64(39000), Status: "bullish"},
				FxRate:       model.Metric{Value: model.ToFloat64(156.8), Status: "very weak"},
				BondYield10Y: model.Metric{Value: model.ToFloat64(0.7), Status: "ultralow"},
			},
			"uk": {
				PolicyRate:   model.Metric{Value: model.ToFloat64(5.25), Status: "high"},
				Inflation:    model.Metric{Value: model.ToFloat64(2.9), Status: "declining"},
				Unemployment: model.Metric{Value: model.ToFloat64(4.2), Status: "normal"},
				PMI:          model.Metric{Value: model.ToFloat64(51.0), Status: "expanding"},
				EquityIndex:  model.Metric{Value: model.ToFloat64(7900), Status: "neutral"},
				FxRate:       model.Metric{Value: model.ToFloat64(1.29), Status: "stable"},
				BondYield10Y: model.Metric{Value: model.ToFloat64(4.1), Status: "elevated"},
			},
			"france": {
				PolicyRate:   model.Metric{Value: model.ToFloat64(4.0), Status: "normal"},
				Inflation:    model.Metric{Value: model.ToFloat64(2.4), Status: "target"},
				Unemployment: model.Metric{Value: model.ToFloat64(7.1), Status: "high"},
				PMI:          model.Metric{Value: model.ToFloat64(47.3), Status: "contracting"},
				EquityIndex:  model.Metric{Value: model.ToFloat64(7500), Status: "neutral"},
				FxRate:       model.Metric{Value: model.ToFloat64(1.1), Status: "stable"},
				BondYield10Y: model.Metric{Value: model.ToFloat64(2.8), Status: "normal"},
			},
			"india": {
				PolicyRate:   model.Metric{Value: model.ToFloat64(6.5), Status: "tight"},
				Inflation:    model.Metric{Value: model.ToFloat64(4.6), Status: "moderate"},
				Unemployment: model.Metric{Value: model.ToFloat64(7.0), Status: "elevated"},
				PMI:          model.Metric{Value: model.ToFloat64(57.5), Status: "strong"},
				EquityIndex:  model.Metric{Value: model.ToFloat64(74000), Status: "bullish"},
				FxRate:       model.Metric{Value: model.ToFloat64(83.1), Status: "weak"},
				BondYield10Y: model.Metric{Value: model.ToFloat64(7.1), Status: "high"},
			},
			"brazil": {
				PolicyRate:   model.Metric{Value: model.ToFloat64(10.5), Status: "very high"},
				Inflation:    model.Metric{Value: model.ToFloat64(3.9), Status: "normal"},
				Unemployment: model.Metric{Value: model.ToFloat64(7.8), Status: "elevated"},
				PMI:          model.Metric{Value: model.ToFloat64(50.2), Status: "borderline"},
				EquityIndex:  model.Metric{Value: model.ToFloat64(126000), Status: "growing"},
				FxRate:       model.Metric{Value: model.ToFloat64(5.3), Status: "stable"},
				BondYield10Y: model.Metric{Value: model.ToFloat64(11.2), Status: "very high"},
			},
		},
	}
}
