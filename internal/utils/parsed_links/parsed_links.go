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
	Currencies   map[string]MetricLink
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
		Currencies: map[string]MetricLink{
			"USD/EUR": {URL: USUSDEUR_Link, Selector: USUSDEUR_Selector},
			"USD/JPY": {URL: USUSDJPY_Link, Selector: USUSDJPY_Selector},
			"USD/GBP": {URL: USUSDGBP_Link, Selector: USUSDGBP_Selector},
			"USD/CNY": {URL: USUSDCNY_Link, Selector: USUSDCNY_Selector},
			"USD/INR": {URL: USUSDINR_Link, Selector: USUSDINR_Selector},
			"USD/BRL": {URL: USUSDBRL_Link, Selector: USUSDBRL_Selector},
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
		Currencies: map[string]MetricLink{
			"CNY/USD": {URL: CNCNYUSD_Link, Selector: CNCNYUSD_Selector},
			"CNY/EUR": {URL: CNCNYEUR_Link, Selector: CNCNYEUR_Selector},
			"CNY/JPY": {URL: CNCNYJPY_Link, Selector: CNCNYJPY_Selector},
			"CNY/GBP": {URL: CNCNYGBP_Link, Selector: CNCNYGBP_Selector},
			"CNY/INR": {URL: CNCNYINR_Link, Selector: CNCNYINR_Selector},
			"CNY/BRL": {URL: CNCNYBRL_Link, Selector: CNCNYBRL_Selector},
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
		Currencies: map[string]MetricLink{
			"EUR/USD": {URL: DEEUREURUSD_Link, Selector: DEEUREURUSD_Selector},
			"EUR/JPY": {URL: DEEUREURJPY_Link, Selector: DEEUREURJPY_Selector},
			"EUR/GBP": {URL: DEEUREURGBP_Link, Selector: DEEUREURGBP_Selector},
			"EUR/CNY": {URL: DEEUREURCNY_Link, Selector: DEEUREURCNY_Selector},
			"EUR/INR": {URL: DEEUREURINR_Link, Selector: DEEUREURINR_Selector},
			"EUR/BRL": {URL: DEEUREURBRL_Link, Selector: DEEUREURBRL_Selector},
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
		Currencies: map[string]MetricLink{
			"JPY/USD": {URL: JPJPYUSD_Link, Selector: JPJPYUSD_Selector},
			"JPY/EUR": {URL: JPJPYEUR_Link, Selector: JPJPYEUR_Selector},
			"JPY/GBP": {URL: JPJPYGBP_Link, Selector: JPJPYGBP_Selector},
			"JPY/CNY": {URL: JPJPYCNY_Link, Selector: JPJPYCNY_Selector},
			"JPY/INR": {URL: JPJPYINR_Link, Selector: JPJPYINR_Selector},
			"JPY/BRL": {URL: JPJPYBRL_Link, Selector: JPJPYBRL_Selector},
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
		Currencies: map[string]MetricLink{
			"GBP/USD": {URL: UKGBPUSD_Link, Selector: UKGBPUSD_Selector},
			"GBP/EUR": {URL: UKGBPEUR_Link, Selector: UKGBPEUR_Selector},
			"GBP/JPY": {URL: UKGBPJPY_Link, Selector: UKGBPJPY_Selector},
			"GBP/CNY": {URL: UKGBPCNY_Link, Selector: UKGBPCNY_Selector},
			"GBP/INR": {URL: UKGBPINR_Link, Selector: UKGBPINR_Selector},
			"GBP/BRL": {URL: UKGBPBRL_Link, Selector: UKGBPBRL_Selector},
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
		Currencies: map[string]MetricLink{
			"EUR/USD": {URL: FREURUSD_Link, Selector: FREURUSD_Selector},
			"EUR/GBP": {URL: FREURGBP_Link, Selector: FREURGBP_Selector},
			"EUR/JPY": {URL: FREURJPY_Link, Selector: FREURJPY_Selector},
			"EUR/CNY": {URL: FREURCNY_Link, Selector: FREURCNY_Selector},
			"EUR/INR": {URL: FREURINR_Link, Selector: FREURINR_Selector},
			"EUR/BRL": {URL: FREURBRL_Link, Selector: FREURBRL_Selector},
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
		Currencies: map[string]MetricLink{
			"INR/USD": {URL: ININRUSD_Link, Selector: ININRUSD_Selector},
			"INR/EUR": {URL: ININREUR_Link, Selector: ININREUR_Selector},
			"INR/JPY": {URL: ININRJPY_Link, Selector: ININRJPY_Selector},
			"INR/GBP": {URL: ININRGBP_Link, Selector: ININRGBP_Selector},
			"INR/CNY": {URL: ININRCNY_Link, Selector: ININRCNY_Selector},
			"INR/BRL": {URL: ININRBRL_Link, Selector: ININRBRL_Selector},
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
		Currencies: map[string]MetricLink{
			"BRL/USD": {URL: BRBRLUSD_Link, Selector: BRBRLUSD_Selector},
			"BRL/EUR": {URL: BRBRLEUR_Link, Selector: BRBRLEUR_Selector},
			"BRL/JPY": {URL: BRBRLJPY_Link, Selector: BRBRLJPY_Selector},
			"BRL/GBP": {URL: BRBRLGBP_Link, Selector: BRBRLGBP_Selector},
			"BRL/CNY": {URL: BRBRLCNY_Link, Selector: BRBRLCNY_Selector},
			"BRL/INR": {URL: BRBRLINR_Link, Selector: BRBRLINR_Selector},
		},
		Bond10Y: MetricLink{
			URL:      BRBond10Y_Link,
			Selector: BRBond10Y_Selector,
		},
	},
}
