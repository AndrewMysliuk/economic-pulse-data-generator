package core

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/llm"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/model"
)

var countries = []string{
	"usa", "china", "germany", "japan",
	"uk", "france", "india", "brazil",
}

// Generate builds the daily payload skeleton (per country metrics) and
// delegates summary generation to LLM. Real fetching/parsing will replace
// the placeholder metrics later.
func Generate(llmClient llm.LLMClient) error {
	today := time.Now().Format("2006-01-02")
	data := model.DailyData{
		Date:      today,
		Countries: make(map[string]model.CountryMetrics),
	}

	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, country := range countries {
		wg.Add(1)
		go func(country string) {
			defer wg.Done()

			// Placeholder metric with unknown status and no sources/average.
			empty := model.MetricDaily{
				Sources: nil,
				Average: nil,
				Status:  "unknown",
				Comment: "",
			}

			// Initialize all metrics for the country (to be filled by real fetchers).
			m := model.CountryMetrics{
				PolicyRate:   empty, // TODO: replace with fetched value(s)
				Inflation:    empty, // TODO: replace with fetched value(s)
				Unemployment: empty, // TODO: replace with fetched value(s)
				PMI:          empty, // TODO: replace with fetched value(s)
				EquityIndex:  empty, // TODO: replace with fetched value(s)
				FxRate:       empty, // TODO: replace with fetched value(s)
				BondYield10Y: empty, // TODO: replace with fetched value(s)
			}

			mu.Lock()
			data.Countries[country] = m
			mu.Unlock()
		}(country)
	}

	wg.Wait()

	// LLM summary (note: openai.go must read MetricDaily.Average or Sources[0])
	summary, tip, err := llmClient.GenerateSummary(data)
	if err != nil {
		return err
	}
	data.Summary.Text = summary
	data.Summary.Tip = tip

	if err := saveJSON(data); err != nil {
		return err
	}

	log.Printf("Data generated for %s and saved.\n", today)
	return nil
}

func saveJSON(data model.DailyData) error {
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

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
