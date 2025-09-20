package money

import (
	"fmt"
	"strings"

	"github.com/AndrewMysliuk/expath-data-generator/internal/schema/enum/currency"
)

type Money struct {
	Amount   int64             `json:"amount"`
	Currency currency.Currency `json:"currency"`
}

func (m Money) ToFloat() float64 {
	return float64(m.Amount) / 100
}

func (m Money) String() string {
	return fmt.Sprintf("%.2f %s", m.ToFloat(), m.Currency)
}

func (m Money) Formatted() string {
	val := fmt.Sprintf("%.2f", m.ToFloat())
	parts := strings.Split(val, ".")
	intPart := parts[0]
	decPart := parts[1]

	var result []string
	for i, r := range reverse(intPart) {
		if i > 0 && i%3 == 0 {
			result = append(result, ",")
		}
		result = append(result, string(r))
	}
	revIntPart := reverse(strings.Join(result, ""))

	return fmt.Sprintf("%s.%s %s", revIntPart, decPart, m.Currency)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
