package util

// Constants for all supported margintypes
const (
	NonMarginProduct      = "0"
	PercentPerTradesValue = "1"
	ValuePerLot           = "3"
	PercentPerTradeSize   = "4"
)

// IsSupportedMargintype returns true if the margintype is supported
func IsSupportedMargintype(margintype string) bool {
	switch margintype {
	case NonMarginProduct, PercentPerTradesValue, ValuePerLot, PercentPerTradeSize:
		return true
	}
	return false
}
