package other_taxes

import "github.com/pkg/errors"

type TaxType string

func (t TaxType) String() string {
	return string(t)
}

const (
	CapitalGains TaxType = "CAPITAL_GAINS"
	Property     TaxType = "PROPERTY"
	Inheritance  TaxType = "INHERITANCE"
	Wealth       TaxType = "WEALTH"
	StampDuty    TaxType = "STAMP_DUTY"
	Transfer     TaxType = "TRANSFER"
	Excise       TaxType = "EXCISE"
	Other        TaxType = "OTHER"
	Unknown      TaxType = "UNKNOWN"
)

func New(str string) (TaxType, error) {
	switch str {
	case "CAPITAL_GAINS":
		return CapitalGains, nil
	case "PROPERTY":
		return Property, nil
	case "INHERITANCE":
		return Inheritance, nil
	case "WEALTH":
		return Wealth, nil
	case "STAMP_DUTY":
		return StampDuty, nil
	case "TRANSFER":
		return Transfer, nil
	case "EXCISE":
		return Excise, nil
	case "OTHER":
		return Other, nil
	case "UNKNOWN":
		return Unknown, nil
	default:
		return "", errors.Errorf("can't define TaxType from: %v available values are: %v", str, Values())
	}
}

func Values() []TaxType {
	return []TaxType{
		CapitalGains,
		Property,
		Inheritance,
		Wealth,
		StampDuty,
		Transfer,
		Excise,
		Other,
		Unknown,
	}
}
