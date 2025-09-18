package currency

import "github.com/pkg/errors"

type Currency string

func (c Currency) String() string {
	return string(c)
}

const (
	EUR     Currency = "EUR"
	RON     Currency = "RON"
	Unknown Currency = "UNKNOWN"
)

func New(str string) (Currency, error) {
	switch str {
	case "EUR":
		return EUR, nil
	case "RON":
		return RON, nil
	case "UNKNOWN":
		return Unknown, nil
	default:
		return "", errors.Errorf("can't define Currency from: %v available values are: %v", str, Values())
	}
}

func Values() []Currency {
	return []Currency{EUR, RON, Unknown}
}
