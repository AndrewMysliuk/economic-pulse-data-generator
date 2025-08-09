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
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema/enum/metric_status"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema/enum/metric_unit"
)

var inputCountries = []string{
	"usa",
	// "china",
	// "germany",
	// "japan",
	// "uk",
	// "france",
	// "india",
	// "brazil",
}

var countryISO = map[string]string{
	"usa": "US",
	// "china":   "CN",
	// "germany": "DE",
	// "japan":   "JP",
	// "uk":      "GB",
	// "france":  "FR",
	// "india":   "IN",
	// "brazil":  "BR",
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

			cm := initEmptyCountryMetrics()

			ctx, cancel := context.WithTimeout(rootCtx, 25*time.Second)
			defer cancel()

			filled, err := llmClient.GenerateCountryMetrics(ctx, iso, today)
			if err != nil {
				log.Printf("LLM country=%s failed: %v", iso, err)
			} else {
				mergeCountryMetrics(&cm, &filled)

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

func NewEmptyMetric(unit metric_unit.MetricUnit) schema.MetricDaily {
	return schema.MetricDaily{
		Sources: nil,
		Average: nil,
		Unit:    unit,
		Status:  metric_status.Unknown,
	}
}

func initEmptyCountryMetrics() schema.CountryMetrics {
	return schema.CountryMetrics{
		PolicyRate:    NewEmptyMetric(metric_unit.RatePct),
		Inflation:     NewEmptyMetric(metric_unit.Percent),
		Unemployment:  NewEmptyMetric(metric_unit.Percent),
		PMI:           NewEmptyMetric(metric_unit.Index),
		EquityIndex:   NewEmptyMetric(metric_unit.Index),
		CurrencyIndex: NewEmptyMetric(metric_unit.Index),
		BondYield10Y:  NewEmptyMetric(metric_unit.Percent),
	}
}

func mergeMetric(dst *schema.MetricDaily, src *schema.MetricDaily) {
	if src == nil {
		return
	}
	if len(src.Sources) > 0 {
		dst.Sources = src.Sources
	}

	if src.Unit != "" {
		dst.Unit = src.Unit
	}
}

func mergeCountryMetrics(dst, src *schema.CountryMetrics) {
	mergeMetric(&dst.PolicyRate, &src.PolicyRate)
	mergeMetric(&dst.Inflation, &src.Inflation)
	mergeMetric(&dst.Unemployment, &src.Unemployment)
	mergeMetric(&dst.PMI, &src.PMI)
	mergeMetric(&dst.EquityIndex, &src.EquityIndex)
	mergeMetric(&dst.CurrencyIndex, &src.CurrencyIndex)
	mergeMetric(&dst.BondYield10Y, &src.BondYield10Y)

	if len(src.FxRates) > 0 {
		dst.FxRates = src.FxRates
	}
}
