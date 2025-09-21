package core

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

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

type Core struct {
	LLMClient llm.LLMClient
	Countries []schema.CountryInfo
}

func (c *Core) Run(ctx context.Context, outDir string) error {
	for _, country := range c.Countries {
		data, err := c.BuildCountryData(ctx, country)
		if err != nil {
			return errors.Wrapf(err, "failed to build data for %s", country.CountryName)
		}

		filePath := fmt.Sprintf("%s/%s.json", outDir, country.CountryAlpha2.Lower())
		if err := SaveCountryDataToFile(data, filePath); err != nil {
			return errors.Wrapf(err, "failed to save data for %s", country.CountryName)
		}

		fmt.Printf("Saved %s data to %s\n", country.CountryName, filePath)
	}
	return nil
}

func (c *Core) BuildCountryData(ctx context.Context, country schema.CountryInfo) (*schema.CountryData, error) {
	immigration, err := c.LLMClient.GetImmigration(ctx, country)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get immigration data")
	}

	taxes, err := c.LLMClient.GetTaxes(ctx, country)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get taxes data")
	}

	finance, err := c.LLMClient.GetFinance(ctx, country)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get finance data")
	}

	costOfLiving, err := c.LLMClient.GetCostOfLiving(ctx, country)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get cost of living data")
	}

	qualityOfLife, err := c.LLMClient.GetQualityOfLife(ctx, country)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get quality of life data")
	}

	return &schema.CountryData{
		Country:       country,
		Immigration:   *immigration,
		Taxes:         *taxes,
		Finance:       *finance,
		CostOfLiving:  *costOfLiving,
		QualityOfLife: *qualityOfLife,
	}, nil
}

func SaveCountryDataToFile(data *schema.CountryData, filePath string) error {
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

func NewCore(LLMClient llm.LLMClient) *Core {
	return &Core{
		LLMClient: LLMClient,
		Countries: Countries,
	}
}
