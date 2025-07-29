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

			// TODO: Replace this with actual fetching logic:
			// Example: use fetcher.Fetch(country, "inflation"), etc.

			empty := model.Metric{
				Value:  nil,
				Status: "unknown",
				Source: "",
			}
			m := model.CountryMetrics{
				PolicyRate:   empty, // TODO: replace with fetched value
				Inflation:    empty, // TODO: replace with fetched value
				Unemployment: empty, // TODO: replace with fetched value
				PMI:          empty, // TODO: replace with fetched value
				EquityIndex:  empty, // TODO: replace with fetched value
				FxRate:       empty, // TODO: replace with fetched value
				BondYield10Y: empty, // TODO: replace with fetched value
			}

			mu.Lock()
			data.Countries[country] = m
			mu.Unlock()
		}(country)
	}

	wg.Wait()

	summary, tip, err := llmClient.GenerateSummary(data)
	if err != nil {
		return err
	}
	data.Summary.Text = summary
	data.Summary.Tip = tip

	err = saveJSON(data)
	if err != nil {
		return err
	}

	log.Printf("Data generated for %s and saved.\n", today)
	return nil
}

func saveJSON(data model.DailyData) error {
	outputDir := "output"
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
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
