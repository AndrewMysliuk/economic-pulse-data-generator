package schema

import (
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema/enum/metric_status"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema/enum/metric_unit"
)

type MetricSource struct {
	Value      *float64               `json:"value,omitempty"` // Число, если точное значение
	From       *float64               `json:"from,omitempty"`  // Начало диапазона
	To         *float64               `json:"to,omitempty"`    // Конец диапазона
	Unit       metric_unit.MetricUnit `json:"unit"`            // "%", "$/m²", "€/month", и т.п.
	SourceUrl  string                 `json:"source_url"`      // Полная ссылка или ID
	SourceName string                 `json:"source_name"`     // Название источника
	Raw        string                 `json:"raw"`             // Оригинальное строковое значение (например, "€2,800 – €3,300 /month")
}

type MetricDaily struct {
	Sources MetricSource               `json:"sources"` // 1..N источников
	Status  metric_status.MetricStatus `json:"status"`  // выставляется твоей логикой порогов
}

type CountryMetrics struct {
	// Income & Living
	IncomeAverageNetMonthlyEUR  MetricDaily `json:"income_average_net_monthly_eur"`
	IncomeLivingWageEstimateEUR MetricDaily `json:"income_living_wage_estimate_eur"`

	// Real Estate
	RealEstatePriceCapitalUSDPerM2  MetricDaily `json:"real_estate_price_capital_usd_per_m2"`
	RealEstatePriceRegionalUSDPerM2 MetricDaily `json:"real_estate_price_regional_usd_per_m2"`
	RealEstatePriceChangeYoYPercent MetricDaily `json:"real_estate_price_change_yoy_percent"`
	RealEstateRentalYieldPercent    MetricDaily `json:"real_estate_rental_yield_percent"`

	// Macroeconomics
	MacroPolicyRatePercent        MetricDaily `json:"macro_policy_rate_percent"`
	MacroInflationCPIYoYPercent   MetricDaily `json:"macro_inflation_cpi_yoy_percent"`
	MacroUnemploymentRatePercent  MetricDaily `json:"macro_unemployment_rate_percent"`
	MacroPMIIndex                 MetricDaily `json:"macro_pmi_index"`
	MacroBondYield10YPercent      MetricDaily `json:"macro_bond_yield_10y_percent"`
	MacroGDPGrowthForecastPercent MetricDaily `json:"macro_gdp_growth_forecast_percent"`

	// Economic Structure (Sector Shares)
	EconStructShareManufacturingPercent         MetricDaily `json:"econ_struct_share_manufacturing_percent"`
	EconStructShareInfoFinancialServicesPercent MetricDaily `json:"econ_struct_share_info_financial_services_percent"`
	EconStructShareTradeLogisticsPercent        MetricDaily `json:"econ_struct_share_trade_logistics_percent"`
	EconStructShareOtherSectorsPercent          MetricDaily `json:"econ_struct_share_other_sectors_percent"`

	// Society & Politics
	SocietyPopulationMillion        MetricDaily `json:"society_population_million"`
	SocietyBirthRatePerWoman        MetricDaily `json:"society_birth_rate_per_woman"`
	SocietyCorruptionIndex100Scale  MetricDaily `json:"society_corruption_index_100_scale"`
	SocietyPoliticalStabilityRating MetricDaily `json:"society_political_stability_rating"`
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

func PtrVal(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func PtrValReverse(f float64) *float64 { return &f }

func NewEmptyMetric(unit metric_unit.MetricUnit) MetricDaily {
	return MetricDaily{
		Sources: MetricSource{
			Value:      nil,
			From:       nil,
			To:         nil,
			Unit:       unit,
			SourceUrl:  "",
			SourceName: "",
			Raw:        "",
		},
		Status: metric_status.Unknown,
	}
}

func InitEmptyCountryMetrics() CountryMetrics {
	return CountryMetrics{
		// Income & Living
		IncomeAverageNetMonthlyEUR:  NewEmptyMetric(metric_unit.Level),
		IncomeLivingWageEstimateEUR: NewEmptyMetric(metric_unit.Level),

		// Real Estate
		RealEstatePriceCapitalUSDPerM2:  NewEmptyMetric(metric_unit.Level),
		RealEstatePriceRegionalUSDPerM2: NewEmptyMetric(metric_unit.Level),
		RealEstatePriceChangeYoYPercent: NewEmptyMetric(metric_unit.Percent),
		RealEstateRentalYieldPercent:    NewEmptyMetric(metric_unit.Percent),

		// Macroeconomics
		MacroPolicyRatePercent:        NewEmptyMetric(metric_unit.RatePct),
		MacroInflationCPIYoYPercent:   NewEmptyMetric(metric_unit.Percent),
		MacroUnemploymentRatePercent:  NewEmptyMetric(metric_unit.Percent),
		MacroPMIIndex:                 NewEmptyMetric(metric_unit.Index),
		MacroBondYield10YPercent:      NewEmptyMetric(metric_unit.Percent),
		MacroGDPGrowthForecastPercent: NewEmptyMetric(metric_unit.Percent),

		// Economic Structure
		EconStructShareManufacturingPercent:         NewEmptyMetric(metric_unit.Percent),
		EconStructShareInfoFinancialServicesPercent: NewEmptyMetric(metric_unit.Percent),
		EconStructShareTradeLogisticsPercent:        NewEmptyMetric(metric_unit.Percent),
		EconStructShareOtherSectorsPercent:          NewEmptyMetric(metric_unit.Percent),

		// Society & Politics
		SocietyPopulationMillion:        NewEmptyMetric(metric_unit.Level),
		SocietyBirthRatePerWoman:        NewEmptyMetric(metric_unit.Level),
		SocietyCorruptionIndex100Scale:  NewEmptyMetric(metric_unit.Index),
		SocietyPoliticalStabilityRating: NewEmptyMetric(metric_unit.Level),
	}
}

func MergeMetric(dst *MetricDaily, src *MetricDaily) {
	if src == nil {
		return
	}

	if src.Sources.Raw != "" || src.Sources.Value != nil || src.Sources.From != nil || src.Sources.To != nil {
		dst.Sources = src.Sources
	}

	if src.Status != "" {
		dst.Status = src.Status
	}
}

func MergeCountryMetrics(dst, src *CountryMetrics) {
	MergeMetric(&dst.IncomeAverageNetMonthlyEUR, &src.IncomeAverageNetMonthlyEUR)
	MergeMetric(&dst.IncomeLivingWageEstimateEUR, &src.IncomeLivingWageEstimateEUR)

	MergeMetric(&dst.RealEstatePriceCapitalUSDPerM2, &src.RealEstatePriceCapitalUSDPerM2)
	MergeMetric(&dst.RealEstatePriceRegionalUSDPerM2, &src.RealEstatePriceRegionalUSDPerM2)
	MergeMetric(&dst.RealEstatePriceChangeYoYPercent, &src.RealEstatePriceChangeYoYPercent)
	MergeMetric(&dst.RealEstateRentalYieldPercent, &src.RealEstateRentalYieldPercent)

	MergeMetric(&dst.MacroPolicyRatePercent, &src.MacroPolicyRatePercent)
	MergeMetric(&dst.MacroInflationCPIYoYPercent, &src.MacroInflationCPIYoYPercent)
	MergeMetric(&dst.MacroUnemploymentRatePercent, &src.MacroUnemploymentRatePercent)
	MergeMetric(&dst.MacroPMIIndex, &src.MacroPMIIndex)
	MergeMetric(&dst.MacroBondYield10YPercent, &src.MacroBondYield10YPercent)
	MergeMetric(&dst.MacroGDPGrowthForecastPercent, &src.MacroGDPGrowthForecastPercent)

	MergeMetric(&dst.EconStructShareManufacturingPercent, &src.EconStructShareManufacturingPercent)
	MergeMetric(&dst.EconStructShareInfoFinancialServicesPercent, &src.EconStructShareInfoFinancialServicesPercent)
	MergeMetric(&dst.EconStructShareTradeLogisticsPercent, &src.EconStructShareTradeLogisticsPercent)
	MergeMetric(&dst.EconStructShareOtherSectorsPercent, &src.EconStructShareOtherSectorsPercent)

	MergeMetric(&dst.SocietyPopulationMillion, &src.SocietyPopulationMillion)
	MergeMetric(&dst.SocietyBirthRatePerWoman, &src.SocietyBirthRatePerWoman)
	MergeMetric(&dst.SocietyCorruptionIndex100Scale, &src.SocietyCorruptionIndex100Scale)
	MergeMetric(&dst.SocietyPoliticalStabilityRating, &src.SocietyPoliticalStabilityRating)
}
