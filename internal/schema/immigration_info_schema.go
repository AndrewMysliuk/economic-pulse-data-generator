package schema

import (
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/permit_type"
)

type ImmigrationInfo struct {
	ResidencePermits   []ResidencePermit  `json:"residence_permits"`
	PermanentResidency PermanentResidency `json:"permanent_residency"`
	Citizenship        Citizenship        `json:"citizenship"`
	DigitalNomadVisa   *DigitalNomadVisa  `json:"digital_nomad_visa,omitempty"`
	OtherOptions       []OtherOption      `json:"other_options,omitempty"`
}

type (
	ResidencePermit struct {
		PermitType         permit_type.PermitType `json:"permit_type"`
		PermitDuration     *RangeOrValue          `json:"permit_duration_years,omitempty"`
		IsPermitRenewable  bool                   `json:"is_permit_renewable"`
		PermitRequirements PermitRequirements     `json:"permit_requirements"`
		PermitDescription  string                 `json:"permit_description"`
	}

	PermanentResidency struct {
		YearsOfResidenceRequired  *RangeOrValue `json:"years_of_residence_required"`
		IsLanguageTestRequired    bool          `json:"is_language_test_required"`
		IsIntegrationTestRequired bool          `json:"is_integration_test_required"`
		Description               string        `json:"permanent_residency_description"`
	}

	Citizenship struct {
		YearsUntilCitizenship    *RangeOrValue `json:"years_until_citizenship"`
		IsDualCitizenshipAllowed bool          `json:"is_dual_citizenship_allowed"`
		IsLanguageRequired       bool          `json:"is_language_required_for_citizenship"`
		OtherConditions          string        `json:"other_citizenship_conditions,omitempty"`
		Description              string        `json:"citizenship_description"`
	}

	DigitalNomadVisa struct {
		IsAvailable          bool          `json:"is_digital_nomad_visa_available"`
		MinIncomeRequirement *RangeOrMoney `json:"min_income_requirement,omitempty"`
		DurationYears        *RangeOrValue `json:"digital_nomad_visa_duration_years,omitempty"`
		IsRenewable          bool          `json:"is_digital_nomad_visa_renewable"`
		Description          string        `json:"digital_nomad_visa_description"`
	}

	OtherOption struct {
		OptionType         string             `json:"option_type"`
		OptionRequirements PermitRequirements `json:"option_requirements"`
		OptionDescription  string             `json:"option_description"`
	}
)

type (
	PermitRequirements struct {
		MinIncomeRequirement  *RangeOrMoney `json:"min_income_requirement,omitempty"`
		InvestmentAmount      *RangeOrMoney `json:"investment_amount,omitempty"`
		OtherRequirementsNote string        `json:"other_requirements_note,omitempty"`
	}
)
