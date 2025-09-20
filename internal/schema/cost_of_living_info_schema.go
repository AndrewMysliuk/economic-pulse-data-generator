package schema

type CostOfLivingInfo struct {
	Housing    HousingInfo    `json:"housing"`
	Groceries  GroceriesInfo  `json:"groceries"`
	Healthcare HealthcareInfo `json:"healthcare"`
	Transport  TransportInfo  `json:"transport"`
	Internet   InternetInfo   `json:"internet"`
}

type (
	HousingInfo struct {
		RentInCapital        RangeOrMoney `json:"cost_of_living_rent_in_capital"`
		RentInProvince       RangeOrMoney `json:"cost_of_living_rent_in_province"`
		BuyPriceM2InCapital  RangeOrMoney `json:"cost_of_living_buy_price_m2_in_capital"`
		BuyPriceM2InProvince RangeOrMoney `json:"cost_of_living_buy_price_m2_in_province"`
		HousingNote          string       `json:"cost_of_living_housing_note,omitempty"`
	}

	GroceriesInfo struct {
		MonthlyBasket  RangeOrMoney `json:"cost_of_living_groceries_monthly_basket"`
		FamilyBasket   RangeOrMoney `json:"cost_of_living_groceries_family_basket"`
		GroceriesIndex int          `json:"cost_of_living_groceries_index"`
		GroceriesNote  string       `json:"cost_of_living_groceries_note,omitempty"`
	}

	HealthcareInfo struct {
		IsAvailable      bool         `json:"is_healthcare_available"`
		InsuranceMonthly RangeOrMoney `json:"cost_of_living_healthcare_insurance_monthly"`
		DoctorVisit      RangeOrMoney `json:"cost_of_living_healthcare_doctor_visit"`
		HealthcareNote   string       `json:"cost_of_living_healthcare_note,omitempty"`
	}

	TransportInfo struct {
		IsAvailable       bool         `json:"is_transport_available"`
		PublicMonthlyPass RangeOrMoney `json:"cost_of_living_transport_public_monthly_pass"`
		GasolineLiter     RangeOrMoney `json:"cost_of_living_transport_gasoline_liter"`
		TaxiStart         RangeOrMoney `json:"cost_of_living_transport_taxi_start"`
		TransportNote     string       `json:"cost_of_living_transport_note,omitempty"`
	}

	InternetInfo struct {
		IsAvailable   bool         `json:"is_internet_available"`
		HomeMonthly   RangeOrMoney `json:"cost_of_living_internet_home_monthly"`
		MobileMonthly RangeOrMoney `json:"cost_of_living_internet_mobile_monthly"`
		InternetNote  string       `json:"cost_of_living_internet_note,omitempty"`
	}
)
