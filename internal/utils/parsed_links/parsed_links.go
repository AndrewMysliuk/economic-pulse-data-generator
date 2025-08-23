package parsed_links

type MetricLink struct {
	URL      string
	Selector string
}

type CountryLinks struct {
	PolicyRate   MetricLink
	Inflation    MetricLink
	Unemployment MetricLink
	PMI          MetricLink
	EquityIndex  MetricLink
	Bond10Y      MetricLink
}

var Countries = map[string]CountryLinks{
	"US": {
		PolicyRate: MetricLink{
			URL:      USPolicyRate_Link,
			Selector: USPolicyRate_Selector,
		},
		Inflation: MetricLink{
			URL:      USInflation_Link,
			Selector: USInflation_Selector,
		},
		Unemployment: MetricLink{
			URL:      USUnemployment_Link,
			Selector: USUnemployment_Selector,
		},
		PMI: MetricLink{
			URL:      USPMI_Link,
			Selector: USPMI_Selector,
		},
		EquityIndex: MetricLink{
			URL:      USEquityIndex_Link,
			Selector: USEquityIndex_Selector,
		},
		Bond10Y: MetricLink{
			URL:      USBond10Y_Link,
			Selector: USBond10Y_Selector,
		},
	},
	"CN": {
		PolicyRate: MetricLink{
			URL:      CNPolicyRate_Link,
			Selector: CNPolicyRate_Selector,
		},
		Inflation: MetricLink{
			URL:      CNInflation_Link,
			Selector: CNInflation_Selector,
		},
		Unemployment: MetricLink{
			URL:      CNUnemployment_Link,
			Selector: CNUnemployment_Selector,
		},
		PMI: MetricLink{
			URL:      CNPMI_Link,
			Selector: CNPMI_Selector,
		},
		EquityIndex: MetricLink{
			URL:      CNEquityIndex_Link,
			Selector: CNEquityIndex_Selector,
		},
		Bond10Y: MetricLink{
			URL:      CNBond10Y_Link,
			Selector: CNBond10Y_Selector,
		},
	},
	"DE": {
		PolicyRate: MetricLink{
			URL:      DEPolicyRate_Link,
			Selector: DEPolicyRate_Selector,
		},
		Inflation: MetricLink{
			URL:      DEInflation_Link,
			Selector: DEInflation_Selector,
		},
		Unemployment: MetricLink{
			URL:      DEUnemployment_Link,
			Selector: DEUnemployment_Selector,
		},
		PMI: MetricLink{
			URL:      DEPMI_Link,
			Selector: DEPMI_Selector,
		},
		EquityIndex: MetricLink{
			URL:      DEEquityIndex_Link,
			Selector: DEEquityIndex_Selector,
		},
		Bond10Y: MetricLink{
			URL:      DEBond10Y_Link,
			Selector: DEBond10Y_Selector,
		},
	},
	"JP": {
		PolicyRate: MetricLink{
			URL:      JPPolicyRate_Link,
			Selector: JPPolicyRate_Selector,
		},
		Inflation: MetricLink{
			URL:      JPInflation_Link,
			Selector: JPInflation_Selector,
		},
		Unemployment: MetricLink{
			URL:      JPUnemployment_Link,
			Selector: JPUnemployment_Selector,
		},
		PMI: MetricLink{
			URL:      JPPMI_Link,
			Selector: JPPMI_Selector,
		},
		EquityIndex: MetricLink{
			URL:      JPEquityIndex_Link,
			Selector: JPEquityIndex_Selector,
		},
		Bond10Y: MetricLink{
			URL:      JPBond10Y_Link,
			Selector: JPBond10Y_Selector,
		},
	},
	"GB": {
		PolicyRate: MetricLink{
			URL:      UKPolicyRate_Link,
			Selector: UKPolicyRate_Selector,
		},
		Inflation: MetricLink{
			URL:      UKInflation_Link,
			Selector: UKInflation_Selector,
		},
		Unemployment: MetricLink{
			URL:      UKUnemployment_Link,
			Selector: UKUnemployment_Selector,
		},
		PMI: MetricLink{
			URL:      UKPMI_Link,
			Selector: UKPMI_Selector,
		},
		EquityIndex: MetricLink{
			URL:      UKEequityIndex_Link,
			Selector: UKEequityIndex_Selector,
		},
		Bond10Y: MetricLink{
			URL:      UKBond10Y_Link,
			Selector: UKBond10Y_Selector,
		},
	},
	"FR": {
		PolicyRate: MetricLink{
			URL:      FRPolicyRate_Link,
			Selector: FRPolicyRate_Selector,
		},
		Inflation: MetricLink{
			URL:      FRInflation_Link,
			Selector: FRInflation_Selector,
		},
		Unemployment: MetricLink{
			URL:      FRUnemployment_Link,
			Selector: FRUnemployment_Selector,
		},
		PMI: MetricLink{
			URL:      FRPMI_Link,
			Selector: FRPMI_Selector,
		},
		EquityIndex: MetricLink{
			URL:      FREquityIndex_Link,
			Selector: FREquityIndex_Selector,
		},
		Bond10Y: MetricLink{
			URL:      FRBond10Y_Link,
			Selector: FRBond10Y_Selector,
		},
	},
	"IN": {
		PolicyRate: MetricLink{
			URL:      INPolicyRate_Link,
			Selector: INPolicyRate_Selector,
		},
		Inflation: MetricLink{
			URL:      INInflation_Link,
			Selector: INInflation_Selector,
		},
		Unemployment: MetricLink{
			URL:      INUnemployment_Link,
			Selector: INUnemployment_Selector,
		},
		PMI: MetricLink{
			URL:      INPMI_Link,
			Selector: INPMI_Selector,
		},
		EquityIndex: MetricLink{
			URL:      INEquityIndex_Link,
			Selector: INEquityIndex_Selector,
		},
		Bond10Y: MetricLink{
			URL:      INBond10Y_Link,
			Selector: INBond10Y_Selector,
		},
	},
	"BR": {
		PolicyRate: MetricLink{
			URL:      BRPolicyRate_Link,
			Selector: BRPolicyRate_Selector,
		},
		Inflation: MetricLink{
			URL:      BRInflation_Link,
			Selector: BRInflation_Selector,
		},
		Unemployment: MetricLink{
			URL:      BRUnemployment_Link,
			Selector: BRUnemployment_Selector,
		},
		PMI: MetricLink{
			URL:      BRPMI_Link,
			Selector: BRPMI_Selector,
		},
		EquityIndex: MetricLink{
			URL:      BREquityIndex_Link,
			Selector: BREquityIndex_Selector,
		},
		Bond10Y: MetricLink{
			URL:      BRBond10Y_Link,
			Selector: BRBond10Y_Selector,
		},
	},
}
