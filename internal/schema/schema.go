package schema

import (
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/country_alpha2"
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/country_name"
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/currency"
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/value_type"
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/type/money"
)

type CountryInfo struct {
	CountryName   country_name.CountryName
	CountryAlpha2 country_alpha2.CountryAlpha2
	Currency      currency.Currency
}

type CountryData struct {
	Country       CountryInfo       `json:"country"`
	Immigration   ImmigrationInfo   `json:"immigration"`
	Taxes         TaxInfo           `json:"taxes"`
	Finance       FinanceInfo       `json:"finance"`
	CostOfLiving  CostOfLivingInfo  `json:"cost_of_living"`
	QualityOfLife QualityOfLifeInfo `json:"quality_of_life"`
}

type (
	RangeOrValue struct {
		Value       *int `json:"value,omitempty"`
		Min         *int `json:"min,omitempty"`
		Max         *int `json:"max,omitempty"`
		IsUnlimited bool `json:"is_unlimited,omitempty"`
	}

	RangeOrMoney struct {
		Value       *money.Money `json:"value,omitempty"`
		Min         *money.Money `json:"min,omitempty"`
		Max         *money.Money `json:"max,omitempty"`
		IsUnlimited bool         `json:"is_unlimited,omitempty"`
	}

	RateWithConditions struct {
		Range          *RangeOrValue `json:"range,omitempty"`
		ConditionsNote string        `json:"conditions_note,omitempty"`
	}

	MonetaryRateWithConditions struct {
		Range          *RangeOrMoney `json:"range,omitempty"`
		ConditionsNote string        `json:"conditions_note,omitempty"`
	}

	SocialContributions struct {
		Type   value_type.ValueType `json:"contribution_type"`
		Amount int64                `json:"contribution_amount"`
		Note   string               `json:"contribution_note,omitempty"`
	}
)
