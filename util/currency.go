package util

const (
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, CAD, EUR:
		return true
	}
	return false
}
