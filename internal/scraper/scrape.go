package scraper

import (
	"log"
)

func RunScraperDemo() {
	InitHTTPClient()

	url := "https://tradingeconomics.com/united-states/inflation-cpi"
	selector := "#actual"

	resp, err := GetClient().R().Get(url)
	if err != nil {
		log.Fatal("Request failed:", err)
	}

	value, err := ExtractFromHTML(string(resp.Body()), selector, "")
	if err != nil {
		log.Fatal("Extraction failed:", err)
	}

	log.Println("Extracted value:", value)
}
