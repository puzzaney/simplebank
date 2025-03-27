package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

func IsSupportedCurency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
