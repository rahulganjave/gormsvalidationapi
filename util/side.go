package util

// Constants for all supported sides
const (
	BUY        = 1
	SELL       = 2
	SHORT_SELL = 3
)

// IsSupportedSide returns true if the side is supported
func IsSupportedSide(side int) bool {
	switch side {
	case BUY, SELL, SHORT_SELL:
		return true
	}
	return false
}
