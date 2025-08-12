package parsed_links

const (
	// Процентная ставка (PBoC / LPR)
	CNPolicyRate_Link     = "https://www.tradingview.com/symbols/ECONOMICS-CNINTR/"
	CNPolicyRate_Selector = "#js-category-content > div.tv-react-category-header > div.js-symbol-page-header-root > div > div > div > div.quotesRow-iJMmXWiA.compact-iJMmXWiA.quotesRowEconomic-iJMmXWiA > div:nth-child(1) > div > div.lastContainer-zoF9r75I > span.last-zoF9r75I.js-symbol-last"

	// Инфляция (CPI YoY)
	CNInflation_Link     = "https://www.tradingview.com/symbols/ECONOMICS-CNIRYY/"
	CNInflation_Selector = "#js-category-content > div.tv-react-category-header > div.js-symbol-page-header-root > div > div > div > div.quotesRow-iJMmXWiA.compact-iJMmXWiA.quotesRowEconomic-iJMmXWiA > div:nth-child(1) > div > div.lastContainer-zoF9r75I > span.last-zoF9r75I.js-symbol-last"

	// Безработица
	CNUnemployment_Link     = "https://www.tradingview.com/symbols/ECONOMICS-CNUR/"
	CNUnemployment_Selector = "#js-category-content > div.tv-react-category-header > div.js-symbol-page-header-root > div > div > div > div.quotesRow-iJMmXWiA.compact-iJMmXWiA.quotesRowEconomic-iJMmXWiA > div:nth-child(1) > div > div.lastContainer-zoF9r75I > span.last-zoF9r75I.js-symbol-last"

	// PMI (официальный NBS Manufacturing PMI у TV висит на CNBCOI)
	CNPMI_Link     = "https://www.tradingview.com/symbols/ECONOMICS-CNBCOI/"
	CNPMI_Selector = "#js-category-content > div.tv-react-category-header > div.js-symbol-page-header-root > div > div > div > div.quotesRow-iJMmXWiA.compact-iJMmXWiA.quotesRowEconomic-iJMmXWiA > div:nth-child(1) > div > div.lastContainer-zoF9r75I > span.last-zoF9r75I.js-symbol-last"

	// Фондовый индекс (SSE Composite)
	CNEquityIndex_Link     = "https://www.tradingview.com/symbols/SSE-000001/"
	CNEquityIndex_Selector = "#js-category-content > div.tv-react-category-header > div.js-symbol-page-header-root > div > div > div > div.quotesRow-iJMmXWiA > div:nth-child(1) > div > div.lastContainer-zoF9r75I > span.last-zoF9r75I.js-symbol-last"

	// Валюты (CNY к базовым)
	CNCNYUSD_Link     = "https://www.tradingview.com/symbols/CNYUSD/"
	CNCNYUSD_Selector = "#js-category-content > div.tv-react-category-header > div.js-symbol-page-header-root > div > div > div > div.quotesRow-iJMmXWiA > div:nth-child(1) > div > div.lastContainer-zoF9r75I > span.last-zoF9r75I.js-symbol-last"

	CNCNYEUR_Link     = "https://www.tradingview.com/symbols/CNYEUR/"
	CNCNYEUR_Selector = "#js-category-content > div.tv-react-category-header > div.js-symbol-page-header-root > div > div > div > div.quotesRow-iJMmXWiA > div:nth-child(1) > div > div.lastContainer-zoF9r75I > span.last-zoF9r75I.js-symbol-last"

	CNCNYJPY_Link     = "https://www.tradingview.com/symbols/CNYJPY/"
	CNCNYJPY_Selector = "#js-category-content > div.tv-react-category-header > div.js-symbol-page-header-root > div > div > div > div.quotesRow-iJMmXWiA > div:nth-child(1) > div > div.lastContainer-zoF9r75I > span.last-zoF9r75I.js-symbol-last"

	CNCNYGBP_Link     = "https://www.tradingview.com/symbols/CNYGBP/"
	CNCNYGBP_Selector = "#js-category-content > div.tv-react-category-header > div.js-symbol-page-header-root > div > div > div > div.quotesRow-iJMmXWiA > div:nth-child(1) > div > div.lastContainer-zoF9r75I > span.last-zoF9r75I.js-symbol-last"

	CNCNYINR_Link     = "https://www.tradingview.com/symbols/CNYINR/"
	CNCNYINR_Selector = "#js-category-content > div.tv-react-category-header > div.js-symbol-page-header-root > div > div > div > div.quotesRow-iJMmXWiA > div:nth-child(1) > div > div.lastContainer-zoF9r75I > span.last-zoF9r75I.js-symbol-last"

	CNCNYBRL_Link     = "https://www.tradingview.com/symbols/CNYBRL/"
	CNCNYBRL_Selector = "#js-category-content > div.tv-react-category-header > div.js-symbol-page-header-root > div > div > div > div.quotesRow-iJMmXWiA > div:nth-child(1) > div > div.lastContainer-zoF9r75I > span.last-zoF9r75I.js-symbol-last"

	// Доходность 10-летних облигаций
	CNBond10Y_Link     = "https://www.tradingview.com/symbols/TVC-CN10Y/"
	CNBond10Y_Selector = "#js-category-content > div.tv-react-category-header > div.js-symbol-page-header-root > div > div > div > div.quotesRow-iJMmXWiA > div:nth-child(1) > div > div.lastContainer-zoF9r75I > span.last-zoF9r75I.js-symbol-last"
)
