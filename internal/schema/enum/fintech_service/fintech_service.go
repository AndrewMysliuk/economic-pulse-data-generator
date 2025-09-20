package fintech_service

import "github.com/pkg/errors"

type FintechService string

func (s FintechService) String() string {
	return string(s)
}

const (
	// Европа/UK
	Revolut  FintechService = "REVOLUT"
	Wise     FintechService = "WISE"
	N26      FintechService = "N26"
	Monzo    FintechService = "MONZO"
	Starling FintechService = "STARLING"

	// Азия (региональные/глобальные)
	BigPay   FintechService = "BIGPAY"
	YouTrip  FintechService = "YOUTRIP"
	Instarem FintechService = "INSTAREM"
	Aspire   FintechService = "ASPIRE"
	Tonik    FintechService = "TONIK"

	// США / глобальные
	Chime    FintechService = "CHIME"
	Payoneer FintechService = "PAYONEER"
	PayPal   FintechService = "PAYPAL"

	Other   FintechService = "OTHER"
	Unknown FintechService = "UNKNOWN"
)

func New(str string) (FintechService, error) {
	switch str {
	case "REVOLUT":
		return Revolut, nil
	case "WISE":
		return Wise, nil
	case "N26":
		return N26, nil
	case "MONZO":
		return Monzo, nil
	case "STARLING":
		return Starling, nil
	case "BIGPAY":
		return BigPay, nil
	case "YOUTRIP":
		return YouTrip, nil
	case "INSTAREM":
		return Instarem, nil
	case "ASPIRE":
		return Aspire, nil
	case "TONIK":
		return Tonik, nil
	case "CHIME":
		return Chime, nil
	case "PAYONEER":
		return Payoneer, nil
	case "PAYPAL":
		return PayPal, nil
	case "OTHER":
		return Other, nil
	case "UNKNOWN":
		return Unknown, nil
	default:
		return "", errors.Errorf("can't define FintechService from: %v available values are: %v", str, Values())
	}
}

func Values() []FintechService {
	return []FintechService{
		Revolut, Wise, N26, Monzo, Starling,
		BigPay, YouTrip, Instarem, Aspire, Tonik,
		Chime, Payoneer, PayPal,
		Other, Unknown,
	}
}
