package core

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/llm"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/scraper"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/utils/status_range"
)

var inputCountries = []string{
	"usa",
	"china",
	"germany",
	"japan",
	"uk",
	"france",
	"india",
	"brazil",
}

var countryISO = map[string]string{
	"usa":     "US",
	"china":   "CN",
	"germany": "DE",
	"japan":   "JP",
	"uk":      "GB",
	"france":  "FR",
	"india":   "IN",
	"brazil":  "BR",
}

func Generate(llmClient llm.LLMClient) error {
	rootCtx := context.Background()
	today := time.Now().UTC().Format("2006-01-02")

	data := schema.DailyData{
		Date:      today,
		Countries: make(map[string]schema.CountryMetrics, len(inputCountries)),
	}

	for _, c := range inputCountries {
		iso := countryISO[c]
		if iso == "" {
			log.Printf("skip unknown country key: %s", c)
			continue
		}

		cm := schema.InitEmptyCountryMetrics()

		scraped, err := scraper.Scrape(iso)
		if err != nil {
			log.Printf("SCRAPE country=%s failed: %v", iso, err)
		} else {
			schema.MergeCountryMetrics(&cm, &scraped)

			cm.PolicyRate.ComputeAverage()
			cm.Inflation.ComputeAverage()
			cm.Unemployment.ComputeAverage()
			cm.PMI.ComputeAverage()
			cm.EquityIndex.ComputeAverage()
			cm.CurrencyIndex.ComputeAverage()
			cm.BondYield10Y.ComputeAverage()

			setStatusesFromAverages(iso, &cm)

			log.Printf(
				"SCRAPE country=%s OK | policy_rate=%.2f | inflation=%.2f | unemployment=%.2f",
				iso,
				schema.PtrVal(cm.PolicyRate.Average),
				schema.PtrVal(cm.Inflation.Average),
				schema.PtrVal(cm.Unemployment.Average),
			)
		}

		data.Countries[iso] = cm

		time.Sleep(700*time.Millisecond + time.Duration(200+rand.Intn(400))*time.Millisecond)
	}

	if summary, err := llmClient.GenerateSummary(rootCtx, data); err == nil {
		data.Summary = summary
	} else {
		log.Printf("failed to generate summary: %v", err)
	}

	if err := saveJSON(data); err != nil {
		return err
	}

	historyDir := filepath.Join("output", "history")
	if err := os.MkdirAll(historyDir, 0o755); err != nil {
		return err
	}

	historyPath := filepath.Join(historyDir, "history_180.json")
	if err := UpdateHistoryIncremental(historyPath, "output"); err != nil {
		log.Printf("failed to update history: %v", err)
	}

	log.Printf("Data generated for %s and saved.\n", today)
	return nil
}

func saveJSON(data schema.DailyData) error {
	outputDir := "output"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	filePath := filepath.Join(outputDir, data.Date+".json")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

func setStatusesFromAverages(iso string, cm *schema.CountryMetrics) {
	thresholds, ok := status_range.ByISO[iso]
	if !ok {
		return
	}

	if cm.PolicyRate.Average != nil {
		cm.PolicyRate.Status = thresholds.PolicyRate.StatusForValue(*cm.PolicyRate.Average)
	}

	if cm.Inflation.Average != nil {
		cm.Inflation.Status = thresholds.Inflation.StatusForValue(*cm.Inflation.Average)
	}

	if cm.Unemployment.Average != nil {
		cm.Unemployment.Status = thresholds.Unemployment.StatusForValue(*cm.Unemployment.Average)
	}

	if cm.PMI.Average != nil {
		cm.PMI.Status = thresholds.PMI.StatusForValue(*cm.PMI.Average)
	}

	if cm.EquityIndex.Average != nil {
		cm.EquityIndex.Status = thresholds.EquityYoY.StatusForValue(*cm.EquityIndex.Average)
	}

	if cm.BondYield10Y.Average != nil {
		cm.BondYield10Y.Status = thresholds.Bond10Y.StatusForValue(*cm.BondYield10Y.Average)
	}
}
