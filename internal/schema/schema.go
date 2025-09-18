package schema

import (
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema/enum/country_alpha2"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema/enum/country_name"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema/enum/currency"
)

type CountryInfo struct {
	CountryName   country_name.CountryName
	CountryAlpha2 country_alpha2.CountryAlpha2
	Currency      currency.Currency
}

type (
	CountryData struct {
		Country       CountryInfo       `json:"country"`
		Immigration   ImmigrationInfo   `json:"immigration"`
		Taxes         TaxInfo           `json:"taxes"`
		Finance       FinanceInfo       `json:"finance"`
		CostOfLiving  CostOfLivingInfo  `json:"cost_of_living"`
		QualityOfLife QualityOfLifeInfo `json:"quality_of_life"`
	}

	ImmigrationInfo struct{}

	TaxInfo struct{}

	FinanceInfo struct{}

	CostOfLivingInfo struct{}

	QualityOfLifeInfo struct{}
)
