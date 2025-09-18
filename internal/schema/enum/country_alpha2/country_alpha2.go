package country_alpha2

import "github.com/pkg/errors"

type CountryAlpha2 string

func (c CountryAlpha2) String() string {
	return string(c)
}

const (
	CY      CountryAlpha2 = "CY"
	RO      CountryAlpha2 = "RO"
	Unknown CountryAlpha2 = "UNKNOWN"
)

func New(str string) (CountryAlpha2, error) {
	switch str {
	case "CY":
		return CY, nil
	case "RO":
		return RO, nil
	case "UNKNOWN":
		return Unknown, nil
	default:
		return "", errors.Errorf("can't define CountryAlpha2 from: %v available values are: %v", str, Values())
	}
}

func Values() []CountryAlpha2 {
	return []CountryAlpha2{CY, RO, Unknown}
}
