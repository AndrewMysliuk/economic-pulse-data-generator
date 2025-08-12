package status_range

import (
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema/enum/metric_status"
)

var BRPolicyRate = schema.RangeSet{
	{Min: nil, Max: schema.PtrValReverse(5.0), Status: metric_status.Excellent},
	{Min: schema.PtrValReverse(5.0), Max: schema.PtrValReverse(8.0), Status: metric_status.Good},
	{Min: schema.PtrValReverse(8.0), Max: schema.PtrValReverse(10.5), Status: metric_status.Moderate},
	{Min: schema.PtrValReverse(10.5), Max: schema.PtrValReverse(12.5), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(12.5), Max: schema.PtrValReverse(15.0), Status: metric_status.Danger},
	{Min: schema.PtrValReverse(15.0), Max: nil, Status: metric_status.Critical},
}

var BRInflation = schema.RangeSet{
	{Min: nil, Max: schema.PtrValReverse(0.0), Status: metric_status.Danger},
	{Min: schema.PtrValReverse(0.0), Max: schema.PtrValReverse(0.5), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(0.5), Max: schema.PtrValReverse(3.0), Status: metric_status.Good},
	{Min: schema.PtrValReverse(3.0), Max: schema.PtrValReverse(4.5), Status: metric_status.Excellent},
	{Min: schema.PtrValReverse(4.5), Max: schema.PtrValReverse(6.0), Status: metric_status.Moderate},
	{Min: schema.PtrValReverse(6.0), Max: schema.PtrValReverse(8.0), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(8.0), Max: schema.PtrValReverse(10.0), Status: metric_status.Danger},
	{Min: schema.PtrValReverse(10.0), Max: nil, Status: metric_status.Critical},
}

var BRUnemployment = schema.RangeSet{
	{Min: schema.PtrValReverse(5.5), Max: schema.PtrValReverse(6.5), Status: metric_status.Excellent},
	{Min: schema.PtrValReverse(6.5), Max: schema.PtrValReverse(7.5), Status: metric_status.Good},
	{Min: schema.PtrValReverse(7.5), Max: schema.PtrValReverse(9.0), Status: metric_status.Moderate},
	{Min: schema.PtrValReverse(9.0), Max: schema.PtrValReverse(11.0), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(11.0), Max: schema.PtrValReverse(13.0), Status: metric_status.Danger},
	{Min: schema.PtrValReverse(13.0), Max: nil, Status: metric_status.Critical},
}

var BRPMI = schema.RangeSet{
	{Min: schema.PtrValReverse(60), Max: nil, Status: metric_status.Excellent},
	{Min: schema.PtrValReverse(55), Max: schema.PtrValReverse(60), Status: metric_status.Good},
	{Min: schema.PtrValReverse(50), Max: schema.PtrValReverse(55), Status: metric_status.Moderate},
	{Min: schema.PtrValReverse(47), Max: schema.PtrValReverse(50), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(44), Max: schema.PtrValReverse(47), Status: metric_status.Danger},
	{Min: nil, Max: schema.PtrValReverse(44), Status: metric_status.Critical},
}

var BREquityYoY = schema.RangeSet{
	{Min: schema.PtrValReverse(15), Max: nil, Status: metric_status.Excellent},
	{Min: schema.PtrValReverse(5), Max: schema.PtrValReverse(15), Status: metric_status.Good},
	{Min: schema.PtrValReverse(-5), Max: schema.PtrValReverse(5), Status: metric_status.Moderate},
	{Min: schema.PtrValReverse(-10), Max: schema.PtrValReverse(-5), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(-15), Max: schema.PtrValReverse(-10), Status: metric_status.Danger},
	{Min: nil, Max: schema.PtrValReverse(-15), Status: metric_status.Critical},
}

var BRBond10Y = schema.RangeSet{
	{Min: schema.PtrValReverse(6.0), Max: schema.PtrValReverse(8.0), Status: metric_status.Excellent},
	{Min: schema.PtrValReverse(8.0), Max: schema.PtrValReverse(9.5), Status: metric_status.Good},
	{Min: schema.PtrValReverse(9.5), Max: schema.PtrValReverse(11.0), Status: metric_status.Moderate},
	{Min: schema.PtrValReverse(11.0), Max: schema.PtrValReverse(13.0), Status: metric_status.Warning},
	{Min: schema.PtrValReverse(13.0), Max: schema.PtrValReverse(16.0), Status: metric_status.Danger},
	{Min: schema.PtrValReverse(16.0), Max: nil, Status: metric_status.Critical},
}
