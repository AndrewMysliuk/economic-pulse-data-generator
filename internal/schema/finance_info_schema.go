package schema

import (
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/fintech_service"
)

type FinanceInfo struct {
	Banking BankingInfo `json:"banking"`
	Fintech FintechInfo `json:"fintech"`
}

type (
	PersonalAccount struct {
		IsAvailable       bool     `json:"is_personal_account_available"`
		RequiresResidency bool     `json:"does_personal_account_require_residency"`
		Requirements      []string `json:"personal_account_requirements,omitempty"`
		RequirementsNote  string   `json:"personal_account_requirements_note,omitempty"`
	}

	BusinessAccount struct {
		IsAvailable          bool     `json:"is_business_account_available"`
		RequiresLocalCompany bool     `json:"does_business_account_require_local_company"`
		Requirements         []string `json:"business_account_requirements,omitempty"`
		RequirementsNote     string   `json:"business_account_requirements_note,omitempty"`
	}

	BankingInfo struct {
		PersonalAccount PersonalAccount `json:"personal_account"`
		BusinessAccount BusinessAccount `json:"business_account"`
	}

	FintechServiceInfo struct {
		Name              fintech_service.FintechService `json:"fintech_service_name"`
		IsAvailable       bool                           `json:"is_fintech_service_available"`
		RequiresResidency bool                           `json:"does_fintech_service_require_residency"`
		RequirementsNote  string                         `json:"fintech_service_requirements_note,omitempty"`
		LimitationsNote   string                         `json:"fintech_service_limitations_note,omitempty"`
	}

	FintechInfo struct {
		Services    []FintechServiceInfo `json:"fintech_services"`
		FintechNote string               `json:"fintech_note,omitempty"`
	}
)
