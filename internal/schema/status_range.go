package schema

import (
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema/enum/metric_status"
)

type Interval struct {
	Min    *float64
	Max    *float64
	Status metric_status.MetricStatus
}

type RangeSet []Interval

func (rs RangeSet) StatusForValue(val float64) metric_status.MetricStatus {
	for _, r := range rs {
		if (r.Min == nil || val >= *r.Min) &&
			(r.Max == nil || val < *r.Max) {
			return r.Status
		}
	}
	return metric_status.Unknown
}
