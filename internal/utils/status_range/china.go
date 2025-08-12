package status_range

import (
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema/enum/metric_status"
)

var CNPolicyRate = schema.RangeSet{
	{Min: nil, Max: schema.PtrValReverse(2.0), Status: metric_status.Excellent},
	{Min: schema.PtrValReverse(2.0), Max: schema.PtrValReverse(3.0), Status: metric_status.Good},
	{Min: schema.PtrValReverse(3.0), Max: schema.PtrValReverse(4.0), Status: metric_status.Moderate},
	{Min: schema.PtrValReverse(4.0), Max: schema.PtrValReverse(5.0), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(5.0), Max: schema.PtrValReverse(6.5), Status: metric_status.Danger},
	{Min: schema.PtrValReverse(6.5), Max: nil, Status: metric_status.Critical},
}

var CNInflation = schema.RangeSet{
	{Min: nil, Max: schema.PtrValReverse(0.0), Status: metric_status.Danger},
	{Min: schema.PtrValReverse(0.0), Max: schema.PtrValReverse(0.5), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(0.5), Max: schema.PtrValReverse(2.5), Status: metric_status.Excellent},
	{Min: schema.PtrValReverse(2.5), Max: schema.PtrValReverse(3.5), Status: metric_status.Good},
	{Min: schema.PtrValReverse(3.5), Max: schema.PtrValReverse(5.0), Status: metric_status.Moderate},
	{Min: schema.PtrValReverse(5.0), Max: schema.PtrValReverse(7.0), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(7.0), Max: schema.PtrValReverse(9.0), Status: metric_status.Danger},
	{Min: schema.PtrValReverse(9.0), Max: nil, Status: metric_status.Critical},
}

var CNUnemployment = schema.RangeSet{
	{Min: schema.PtrValReverse(4.0), Max: schema.PtrValReverse(5.0), Status: metric_status.Excellent},
	{Min: schema.PtrValReverse(5.0), Max: schema.PtrValReverse(5.5), Status: metric_status.Good},
	{Min: schema.PtrValReverse(5.5), Max: schema.PtrValReverse(6.5), Status: metric_status.Moderate},
	{Min: schema.PtrValReverse(6.5), Max: schema.PtrValReverse(7.5), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(7.5), Max: schema.PtrValReverse(9.0), Status: metric_status.Danger},
	{Min: nil, Max: schema.PtrValReverse(7.5), Status: metric_status.Critical},
}

var CNPMI = schema.RangeSet{
	{Min: schema.PtrValReverse(60), Max: nil, Status: metric_status.Excellent},
	{Min: schema.PtrValReverse(55), Max: schema.PtrValReverse(60), Status: metric_status.Good},
	{Min: schema.PtrValReverse(50), Max: schema.PtrValReverse(55), Status: metric_status.Moderate},
	{Min: schema.PtrValReverse(47), Max: schema.PtrValReverse(50), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(44), Max: schema.PtrValReverse(47), Status: metric_status.Danger},
	{Min: nil, Max: schema.PtrValReverse(44), Status: metric_status.Critical},
}

var CNEquityYoY = schema.RangeSet{
	{Min: schema.PtrValReverse(15), Max: nil, Status: metric_status.Excellent},
	{Min: schema.PtrValReverse(5), Max: schema.PtrValReverse(15), Status: metric_status.Good},
	{Min: schema.PtrValReverse(-5), Max: schema.PtrValReverse(5), Status: metric_status.Moderate},
	{Min: schema.PtrValReverse(-10), Max: schema.PtrValReverse(-5), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(-15), Max: schema.PtrValReverse(-10), Status: metric_status.Danger},
	{Min: nil, Max: schema.PtrValReverse(-15), Status: metric_status.Critical},
}

var CNBond10Y = schema.RangeSet{
	{Min: schema.PtrValReverse(1.3), Max: schema.PtrValReverse(2.3), Status: metric_status.Excellent},
	{Min: schema.PtrValReverse(2.3), Max: schema.PtrValReverse(2.8), Status: metric_status.Good},
	{Min: schema.PtrValReverse(2.8), Max: schema.PtrValReverse(3.5), Status: metric_status.Moderate},
	{Min: schema.PtrValReverse(3.5), Max: schema.PtrValReverse(4.5), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(4.5), Max: schema.PtrValReverse(6.0), Status: metric_status.Danger},
	{Min: schema.PtrValReverse(6.0), Max: nil, Status: metric_status.Critical},
}
