package prompts

import (
	"fmt"

	"github.com/AndrewMysliuk/expath-data-generator/internal/schema"
)

func ImmigrationInfoPrompt(country schema.CountryInfo) string {
	return fmt.Sprintf(`You are an expert in immigration law and visa policies.
Provide immigration information for the country %s (%s).
Use the local currency %s when specifying monetary requirements.

The information must be returned strictly in JSON format according to the provided schema.
Do not include any text outside of JSON. 
If some fields are unknown, fill them with sensible defaults:
- For strings: "" (empty string)
- For numbers (including monetary amounts): 0
- For booleans: false
- For enums: "UNKNOWN"
- IMPORTANT: All monetary "amount" fields must be expressed as integer values in CENTS (int64).
  Example: 123456 means 1234.56 %s.

Make sure to cover:
- Residence permits (work, freelance, student, investor, family, asylum, retirement, other)
- Permanent residency
- Citizenship
- Digital nomad visa (if applicable)
- Other possible options

Output ONLY a JSON object with the root key "immigration".`,
		country.CountryName, country.CountryAlpha2, country.Currency, country.Currency)
}

func TaxInfoPrompt(country schema.CountryInfo) string {
	return fmt.Sprintf(`You are an expert in international taxation and fiscal policy.
Provide detailed tax information for the country %s (%s).
Use the local currency %s when specifying monetary amounts.

The information must be returned strictly in JSON format according to the provided schema.
Do not include any text outside of JSON. 
If some fields are unknown, fill them with sensible defaults:
- For strings: "" (empty string)
- For numbers (including monetary amounts): 0
- For booleans: false
- For enums: "UNKNOWN"
- IMPORTANT: All monetary "amount" fields must be expressed as integer values in CENTS (int64).
  Example: 123456 means 1234.56 %s.

Make sure to cover:
- Personal income tax (progressive or flat, tax brackets, description)
- Freelance/self-employment tax regimes
- Corporate tax (base rate, special regimes, description)
- Dividend taxation (rates, withholding, conditions)
- Value-added tax (standard rate, reduced/exempt rates, description)
- Other taxes (capital gains, property, inheritance, excise, etc.)

Output ONLY a JSON object with the root key "taxes".`,
		country.CountryName, country.CountryAlpha2, country.Currency, country.Currency)
}

func FinanceInfoPrompt(country schema.CountryInfo) string {
	return fmt.Sprintf(`You are an expert in international banking and fintech services.
Provide detailed finance information for the country %s (%s).

The information must be returned strictly in JSON format according to the provided schema.
Do not include any text outside of JSON. 
If some fields are unknown, fill them with sensible defaults:
- For strings: "" (empty string)
- For numbers (including monetary amounts): 0
- For booleans: false
- For enums: "UNKNOWN"
- IMPORTANT: All monetary "amount" fields must be expressed as integer values in CENTS (int64).
  Example: 123456 means 1234.56 %s.

Make sure to cover:
- Banking: personal accounts (availability, residency requirements, requirements list, notes)
- Banking: business accounts (availability, local company requirements, requirements list, notes)
- Fintech: available services (Revolut, Wise, N26, PayPal, etc.), their availability, residency requirements, limitations, notes

Output ONLY a JSON object with the root key "finance".`,
		country.CountryName, country.CountryAlpha2, country.Currency)
}

func CostOfLivingInfoPrompt(country schema.CountryInfo) string {
	return fmt.Sprintf(`You are an expert in international economics and cost of living analysis.
Provide cost of living information for the country %s (%s).
Use the local currency %s when specifying monetary amounts.

The information must be returned strictly in JSON format according to the provided schema.
Do not include any text outside of JSON. 
If some fields are unknown, fill them with sensible defaults:
- For strings: "" (empty string)
- For numbers (including monetary amounts): 0
- For booleans: false
- For enums: "UNKNOWN"
- IMPORTANT: All monetary "amount" fields must be expressed as integer values in CENTS (int64).
  Example: 123456 means 1234.56 %s.

Make sure to cover:
- Housing costs (rent in capital/province, price per m2 in capital/province, housing note)
- Groceries (monthly basket, family basket, index, groceries note)
- Healthcare (availability, insurance monthly, doctor visit, healthcare note)
- Transport (availability, monthly public transport pass, gasoline liter, taxi start, transport note)
- Internet (availability, home monthly, mobile monthly, internet note)

Output ONLY a JSON object with the root key "cost_of_living".`,
		country.CountryName, country.CountryAlpha2, country.Currency, country.Currency)
}

func QualityOfLifeInfoPrompt(country schema.CountryInfo) string {
	return fmt.Sprintf(`You are an expert in international quality of life studies.
Provide quality of life information for the country %s (%s).

The information must be returned strictly in JSON format according to the provided schema.
Do not include any text outside of JSON. 
If some fields are unknown, fill them with sensible defaults:
- For strings: "" (empty string)
- For numbers (including monetary amounts): 0
- For booleans: false
- For enums: "UNKNOWN"
- IMPORTANT: All monetary "amount" fields must be expressed as integer values in CENTS (int64).
  Example: 123456 means 1234.56 %s.

Make sure to cover:
- Safety (safety index, crime index, safety note)
- Climate (average temperature °C, sunny days per year, rainy days per year, average humidity percent, climate note)
- Ecology (air quality index, water quality index, ecology note)
- Language (official languages, English proficiency index, language note)

Output ONLY a JSON object with the root key "quality_of_life".`,
		country.CountryName, country.CountryAlpha2, country.Currency)
}
