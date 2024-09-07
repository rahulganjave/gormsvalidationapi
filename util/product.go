package util

// Constants for all supported currencies
const (
	CS  = "CS"
	OPT = "OPT"
	FUR = "FUR"
	CFD = "CFD"
	BO  = "BO"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedProduct(product string) bool {
	switch product {
	case CS, OPT, FUR, CFD, BO:
		return true
	}
	return false
}
