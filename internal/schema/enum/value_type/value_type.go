package value_type

import "github.com/pkg/errors"

type ValueType string

func (v ValueType) String() string {
	return string(v)
}

const (
	Percent ValueType = "PERCENT"
	Money   ValueType = "MONEY"
	Unknown ValueType = "UNKNOWN"
)

func New(str string) (ValueType, error) {
	switch str {
	case "PERCENT":
		return Percent, nil
	case "MONEY":
		return Money, nil
	case "UNKNOWN":
		return Unknown, nil
	default:
		return "", errors.Errorf("can't define ValueType from: %v available values are: %v", str, Values())
	}
}

func Values() []ValueType {
	return []ValueType{Percent, Money, Unknown}
}
