package schema

type QualityOfLifeInfo struct {
	Safety   SafetyInfo   `json:"safety"`
	Climate  ClimateInfo  `json:"climate"`
	Ecology  EcologyInfo  `json:"ecology"`
	Language LanguageInfo `json:"language"`
}

type (
	SafetyInfo struct {
		SafetyIndex int    `json:"quality_of_life_safety_index"`
		CrimeIndex  int    `json:"quality_of_life_crime_index"`
		Note        string `json:"quality_of_life_safety_note,omitempty"`
	}

	ClimateInfo struct {
		AvgTempC         int    `json:"quality_of_life_avg_temp_c"`
		SunnyDaysPerYear int    `json:"quality_of_life_sunny_days_per_year"`
		RainDaysPerYear  int    `json:"quality_of_life_rain_days_per_year"`
		AvgHumidityPct   int    `json:"quality_of_life_avg_humidity_percent"`
		Note             string `json:"quality_of_life_climate_note,omitempty"`
	}

	EcologyInfo struct {
		AirQualityIndex   int    `json:"quality_of_life_air_quality_index"`
		WaterQualityIndex int    `json:"quality_of_life_water_quality_index"`
		Note              string `json:"quality_of_life_ecology_note,omitempty"`
	}

	LanguageInfo struct {
		OfficialLanguages       []string `json:"quality_of_life_official_languages,omitempty"`
		EnglishProficiencyIndex int      `json:"quality_of_life_english_proficiency_index"`
		Note                    string   `json:"quality_of_life_language_note,omitempty"`
	}
)
