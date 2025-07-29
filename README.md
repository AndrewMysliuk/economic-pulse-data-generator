# Economic Pulse — Data Generator

**Economic Pulse** is a backend utility that scrapes and aggregates key macroeconomic indicators, generates a unified JSON file, and produces a short summary using GPT-4o.  
This data is consumed by the [Economic Pulse frontend](https://github.com/AndrewMysliuk/economic-pulse-frontend-app) as part of a static site build (SSG).

## What It Does

- Scrapes public macroeconomic data:
  - Interest rates, inflation, unemployment, PMI, FX rates, equity indices, bond yields
- Covers 8 key economies:
  - US, China, Germany, Japan, UK, France, India, Brazil
- Calls OpenAI GPT-4o with a strict JSON schema to generate:
  - `summary`: brief economic insight
  - `tip`: investment suggestion
- Outputs a clean JSON file: `output/YYYY-MM-DD.json`

## Tech Stack

- Go 1.22+
- OpenAI API (GPT-4o)
- JSON-only output (no DB)
- Future: async scraping, scheduled runs
