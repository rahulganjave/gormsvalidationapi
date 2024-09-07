package util

// Constants for all supported currencies
const (
	NASD = "NASD"
	NYSE = "NYSE"
	AMEX = "AMEX"
	SEHK = "SEHK"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedExchange(exchange string) bool {
	switch exchange {
	case NASD, NYSE, AMEX, SEHK:
		return true
	}
	return false
}
