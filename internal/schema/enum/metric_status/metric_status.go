package metric_status

import "github.com/pkg/errors"

type MetricStatus string

func (enum MetricStatus) Equals(str string) bool {
	return string(enum) == str
}

func (enum MetricStatus) Pointer() *MetricStatus {
	return &enum
}

func (enum MetricStatus) String() string {
	return string(enum)
}

const (
	Excellent MetricStatus = "EXCELLENT"
	Good      MetricStatus = "GOOD"
	Moderate  MetricStatus = "MODERATE"
	Warning   MetricStatus = "WARNING"
	Danger    MetricStatus = "DANGER"
	Critical  MetricStatus = "CRITICAL"
	Unknown   MetricStatus = "UNKNOWN"
)

func New(str string) (MetricStatus, error) {
	switch str {

	case "EXCELLENT":
		return Excellent, nil

	case "GOOD":
		return Good, nil

	case "MODERATE":
		return Moderate, nil

	case "WARNING":
		return Warning, nil

	case "DANGER":
		return Danger, nil

	case "CRITICAL":
		return Critical, nil

	case "UNKNOWN":
		return Unknown, nil

	default:
		return "", errors.Errorf("can't define MetricStatus from: %v available values are: %v", str, Values())
	}
}

func Values() []MetricStatus {
	return []MetricStatus{Excellent, Good, Moderate, Warning, Danger, Critical, Unknown}
}
