// Country thresholds metadata
// Last updated: 2025-08-12
// Review schedule:
//   - US, DE, JP, UK, FR → review every 1–2 years
//   - CN, IN, BR → review every 6–12 months
// Triggers for earlier review:
//   - Major global or local economic shock (e.g., pandemic, war, energy crisis)
//   - Policy interest rate change > 2% within 6 months
//   - Inflation deviation > 3% YoY from historical average
//   - Currency volatility > 10% vs USD/EUR within 3 months

package status_range

import (
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
)

type CountryThresholds struct {
	PolicyRate   schema.RangeSet
	Inflation    schema.RangeSet
	Unemployment schema.RangeSet
	PMI          schema.RangeSet
	EquityYoY    schema.RangeSet
	Bond10Y      schema.RangeSet
}

var US = CountryThresholds{
	PolicyRate:   USPolicyRate,
	Inflation:    USInflation,
	Unemployment: USUnemployment,
	PMI:          USPMI,
	EquityYoY:    USEquityYoY,
	Bond10Y:      USBond10Y,
}

var CN = CountryThresholds{
	PolicyRate:   CNPolicyRate,
	Inflation:    CNInflation,
	Unemployment: CNUnemployment,
	PMI:          CNPMI,
	EquityYoY:    CNEquityYoY,
	Bond10Y:      CNBond10Y,
}

var DE = CountryThresholds{
	PolicyRate:   DEPolicyRate,
	Inflation:    DEInflation,
	Unemployment: DEUnemployment,
	PMI:          DEPMI,
	EquityYoY:    DEEquityYoY,
	Bond10Y:      DEBond10Y,
}

var JP = CountryThresholds{
	PolicyRate:   JPPolicyRate,
	Inflation:    JPInflation,
	Unemployment: JPUnemployment,
	PMI:          JPPMI,
	EquityYoY:    JPEquityYoY,
	Bond10Y:      JPBond10Y,
}

var GB = CountryThresholds{
	PolicyRate:   UKPolicyRate,
	Inflation:    UKInflation,
	Unemployment: UKUnemployment,
	PMI:          UKPMI,
	EquityYoY:    UKEqutiyYoY,
	Bond10Y:      UKBond10Y,
}

var FR = CountryThresholds{
	PolicyRate:   FRPolicyRate,
	Inflation:    FRInflation,
	Unemployment: FRUnemployment,
	PMI:          FRPMI,
	EquityYoY:    FREquityYoY,
	Bond10Y:      FRBond10Y,
}

var IN = CountryThresholds{
	PolicyRate:   INPolicyRate,
	Inflation:    INInflation,
	Unemployment: INUnemployment,
	PMI:          INPMI,
	EquityYoY:    INEquityYoY,
	Bond10Y:      INBond10Y,
}

var BR = CountryThresholds{
	PolicyRate:   BRPolicyRate,
	Inflation:    BRInflation,
	Unemployment: BRUnemployment,
	PMI:          BRPMI,
	EquityYoY:    BREquityYoY,
	Bond10Y:      BRBond10Y,
}

var ByISO = map[string]CountryThresholds{
	"US": US,
	"CN": CN,
	"DE": DE,
	"JP": JP,
	"GB": GB,
	"FR": FR,
	"IN": IN,
	"BR": BR,
}

func SetStatusesFromAverages(iso string, cm *schema.CountryMetrics) {
	thresholds, ok := ByISO[iso]
	if !ok {
		return
	}

	if cm.PolicyRate.Average != nil {
		cm.PolicyRate.Status = thresholds.PolicyRate.StatusForValue(*cm.PolicyRate.Average)
	}

	if cm.Inflation.Average != nil {
		cm.Inflation.Status = thresholds.Inflation.StatusForValue(*cm.Inflation.Average)
	}

	if cm.Unemployment.Average != nil {
		cm.Unemployment.Status = thresholds.Unemployment.StatusForValue(*cm.Unemployment.Average)
	}

	if cm.PMI.Average != nil {
		cm.PMI.Status = thresholds.PMI.StatusForValue(*cm.PMI.Average)
	}

	if cm.EquityIndex.Average != nil {
		cm.EquityIndex.Status = thresholds.EquityYoY.StatusForValue(*cm.EquityIndex.Average)
	}

	if cm.BondYield10Y.Average != nil {
		cm.BondYield10Y.Status = thresholds.Bond10Y.StatusForValue(*cm.BondYield10Y.Average)
	}
}
