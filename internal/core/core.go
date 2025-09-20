package core

import (
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema"
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/country_alpha2"
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/country_name"
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/currency"
)

var Countries = []schema.CountryInfo{
	{
		CountryName:   country_name.Cyprus,
		CountryAlpha2: country_alpha2.CY,
		Currency:      currency.EUR,
	},
	{
		CountryName:   country_name.Romania,
		CountryAlpha2: country_alpha2.RO,
		Currency:      currency.RON,
	},
}
