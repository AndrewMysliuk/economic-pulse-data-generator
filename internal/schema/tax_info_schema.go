package schema

import (
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/other_taxes"
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/type/money"
)

type TaxInfo struct {
	PersonalIncome PersonalIncomeTax `json:"personal_income"`
	Freelance      FreelanceTax      `json:"freelance"`
	Corporate      CorporateTax      `json:"corporate"`
	Dividends      DividendsTax      `json:"dividends"`
	Vat            VatTax            `json:"vat"`
	OtherTaxes     []OtherTax        `json:"other_taxes,omitempty"`
}

type (
	PersonalIncomeTaxRate struct {
		FromIncome  *money.Money `json:"from_income,omitempty"`
		ToIncome    *money.Money `json:"to_income,omitempty"`
		RatePercent float64      `json:"rate_percent"`
	}

	PersonalIncomeTax struct {
		IsAvailable   bool                    `json:"is_personal_income_tax_available"`
		IsProgressive bool                    `json:"is_personal_income_tax_progressive"`
		Rates         []PersonalIncomeTaxRate `json:"personal_income_tax_rates,omitempty"`
		Description   string                  `json:"personal_income_tax_description"`
	}

	FreelanceRegime struct {
		Name                string              `json:"freelance_regime_name"`
		RatePercent         float64             `json:"freelance_regime_rate_percent"`
		IncomeLimit         *RangeOrMoney       `json:"freelance_regime_income_limit,omitempty"`
		SocialContributions SocialContributions `json:"freelance_regime_social_contributions"`
		RequirementsNote    string              `json:"freelance_regime_requirements_note,omitempty"`
		Description         string              `json:"freelance_regime_description"`
	}

	FreelanceTax struct {
		IsAvailable bool              `json:"is_freelance_tax_available"`
		Regimes     []FreelanceRegime `json:"freelance_tax_regimes,omitempty"`
	}

	SpecialRegime struct {
		Name             string  `json:"special_regime_name"`
		EffectiveRate    float64 `json:"special_regime_effective_rate_percent"`
		RequirementsNote string  `json:"special_regime_requirements_note,omitempty"`
		Description      string  `json:"special_regime_description"`
	}

	CorporateTax struct {
		IsAvailable    bool               `json:"is_corporate_tax_available"`
		Rate           RateWithConditions `json:"corporate_tax_rate"`
		SpecialRegimes []SpecialRegime    `json:"corporate_tax_special_regimes,omitempty"`
		Description    string             `json:"corporate_tax_description"`
	}

	DividendsTaxRate struct {
		AppliesTo      string  `json:"applies_to"`
		RatePercent    float64 `json:"rate_percent"`
		ConditionsNote string  `json:"conditions_note,omitempty"`
	}

	DividendsTax struct {
		IsAvailable          bool               `json:"is_dividends_tax_available"`
		Rates                []DividendsTaxRate `json:"dividends_tax_rates,omitempty"`
		IsWithholdingApplied bool               `json:"is_withholding_tax_applied"`
		Description          string             `json:"dividends_tax_description"`
	}

	VatConditionRate struct {
		RatePercent    float64 `json:"rate_percent"`
		AppliesTo      string  `json:"applies_to,omitempty"`
		ConditionsNote string  `json:"conditions_note,omitempty"`
	}

	VatTax struct {
		IsAvailable         bool               `json:"is_vat_available"`
		StandardRatePercent float64            `json:"vat_standard_rate_percent"`
		ConditionRates      []VatConditionRate `json:"vat_condition_rates,omitempty"`
		Description         string             `json:"vat_description"`
	}

	OtherTax struct {
		TaxType     other_taxes.TaxType `json:"other_tax_type"`
		Rate        RateWithConditions  `json:"other_tax_rate"`
		AppliesTo   string              `json:"other_tax_applies_to,omitempty"`
		Description string              `json:"other_tax_description"`
	}
)
