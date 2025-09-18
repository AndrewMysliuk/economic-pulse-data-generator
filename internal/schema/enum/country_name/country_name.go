package country_name

import "github.com/pkg/errors"

type CountryName string

func (c CountryName) String() string {
	return string(c)
}

const (
	Cyprus  CountryName = "Cyprus"
	Romania CountryName = "Romania"
	Unknown CountryName = "Unknown"
)

func New(str string) (CountryName, error) {
	switch str {
	case "Cyprus":
		return Cyprus, nil
	case "Romania":
		return Romania, nil
	case "Unknown":
		return Unknown, nil
	default:
		return "", errors.Errorf("can't define CountryName from: %v available values are: %v", str, Values())
	}
}

func Values() []CountryName {
	return []CountryName{Cyprus, Romania, Unknown}
}
