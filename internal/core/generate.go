package core

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/llm"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
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

	eg, _ := errgroup.WithContext(rootCtx)
	eg.SetLimit(4)

	for _, c := range inputCountries {
		cc := c
		eg.Go(func() error {
			iso := countryISO[cc]
			if iso == "" {
				log.Printf("skip unknown country key: %s", cc)
				return nil
			}

			cm := schema.InitEmptyCountryMetrics()

			ctx, cancel := context.WithTimeout(rootCtx, 120*time.Second)
			defer cancel()

			filled, err := llmClient.GenerateCountryMetrics(ctx, iso, today)
			if err != nil {
				log.Printf("LLM country=%s failed: %v", iso, err)
			} else {
				schema.MergeCountryMetrics(&cm, &filled)

				cm.PolicyRate.ComputeAverage()
				cm.Inflation.ComputeAverage()
				cm.Unemployment.ComputeAverage()
				cm.PMI.ComputeAverage()
				cm.EquityIndex.ComputeAverage()
				cm.CurrencyIndex.ComputeAverage()
				cm.BondYield10Y.ComputeAverage()

				log.Printf(
					"LLM country=%s OK | policy_rate=%.2f | inflation=%.2f | unemployment=%.2f",
					iso,
					schema.PtrVal(cm.PolicyRate.Average),
					schema.PtrVal(cm.Inflation.Average),
					schema.PtrVal(cm.Unemployment.Average),
				)
			}

			data.Countries[iso] = cm
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return err
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
