package core

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"

	"github.com/AndrewMysliuk/expath-data-generator/internal/llm"
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema"
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/country_alpha2"
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/country_name"
	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/currency"
)

var Countries = []schema.CountryInfo{
	{
		CountryName:   country_name.Cyprus,
		CountryAlpha2: country_alpha2.CY,
		Currency:      currency.EUR,
	},
	{
		CountryName:   country_name.Romania,
		CountryAlpha2: country_alpha2.RO,
		Currency:      currency.RON,
	},
}

type job struct {
	Country schema.CountryInfo
	Field   string
}

type result struct {
	Country schema.CountryInfo
	Field   string
	Value   interface{}
	Err     error
}

type Core struct {
	LLMClient llm.LLMClient
	Countries []schema.CountryInfo
}

func NewCore(LLMClient llm.LLMClient) *Core {
	return &Core{
		LLMClient: LLMClient,
		Countries: Countries,
	}
}

func (c *Core) Run(ctx context.Context, outDir string) error {
	jobs := make(chan job)
	results := make(chan result)

	var wg sync.WaitGroup

	numWorkers := 10
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				res := c.handleJob(ctx, j)
				results <- res
			}
		}()
	}

	go func() {
		for _, country := range c.Countries {
			for _, field := range []string{"immigration", "taxes", "finance", "costOfLiving", "qualityOfLife"} {
				jobs <- job{Country: country, Field: field}
			}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	countryResults := make(map[string]*schema.CountryData)
	pending := make(map[string]int)

	for _, country := range c.Countries {
		pending[country.CountryAlpha2.String()] = 5
		countryResults[country.CountryAlpha2.String()] = &schema.CountryData{Country: country}
	}

	for r := range results {
		if r.Err != nil {
			return errors.Wrapf(r.Err, "failed to get %s for %s", r.Field, r.Country.CountryName)
		}

		code := r.Country.CountryAlpha2.String()
		data := countryResults[code]

		switch r.Field {
		case "immigration":
			data.Immigration = *(r.Value.(*schema.ImmigrationInfo))
		case "taxes":
			data.Taxes = *(r.Value.(*schema.TaxInfo))
		case "finance":
			data.Finance = *(r.Value.(*schema.FinanceInfo))
		case "costOfLiving":
			data.CostOfLiving = *(r.Value.(*schema.CostOfLivingInfo))
		case "qualityOfLife":
			data.QualityOfLife = *(r.Value.(*schema.QualityOfLifeInfo))
		}

		pending[code]--
		if pending[code] == 0 {
			filePath := fmt.Sprintf("%s/%s.json", outDir, r.Country.CountryAlpha2.Lower())
			if err := SaveCountryDataToFile(data, filePath); err != nil {
				return errors.Wrapf(err, "failed to save %s", r.Country.CountryName)
			}
			fmt.Printf("Saved %s data to %s\n", r.Country.CountryName, filePath)
		}
	}

	return nil
}

func (c *Core) handleJob(ctx context.Context, j job) result {
	switch j.Field {
	case "immigration":
		val, err := c.LLMClient.GetImmigration(ctx, j.Country)
		return result{Country: j.Country, Field: j.Field, Value: val, Err: err}
	case "taxes":
		val, err := c.LLMClient.GetTaxes(ctx, j.Country)
		return result{Country: j.Country, Field: j.Field, Value: val, Err: err}
	case "finance":
		val, err := c.LLMClient.GetFinance(ctx, j.Country)
		return result{Country: j.Country, Field: j.Field, Value: val, Err: err}
	case "costOfLiving":
		val, err := c.LLMClient.GetCostOfLiving(ctx, j.Country)
		return result{Country: j.Country, Field: j.Field, Value: val, Err: err}
	case "qualityOfLife":
		val, err := c.LLMClient.GetQualityOfLife(ctx, j.Country)
		return result{Country: j.Country, Field: j.Field, Value: val, Err: err}
	}

	return result{Country: j.Country, Field: j.Field, Err: fmt.Errorf("unknown field %s", j.Field)}
}

func SaveCountryDataToFile(data *schema.CountryData, filePath string) error {
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return fmt.Errorf("failed to create dir for %s: %w", filePath, err)
	}

	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}
