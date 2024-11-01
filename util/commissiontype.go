package util

// Constants for all supported commisiontypes
const (
	C_PercentPerTradesValue = 1
	C_ValuePerLot           = 2
	C_FixedValuePerTrades   = 3
)

// IsSupportedCommisionTypes returns true if the commisiontypes is supported
func IsSupportedCommisionTypes(commisiontypes int) bool {
	switch commisiontypes {
	case C_PercentPerTradesValue, C_ValuePerLot, C_FixedValuePerTrades:
		return true
	}
	return false
}
