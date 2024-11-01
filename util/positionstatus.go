package util

// Constants for all supported posistionstatus
const (
	POS_OPEN        = 1
	POS_CLOSED      = 2
	POS_HALF_CLOSED = 3
)

// IsSupportedpositionstatus returns true if the posistionstatus is supported
func IsSupportedPositionstatus(positionstatus int) bool {
	switch positionstatus {
	case POS_OPEN, POS_CLOSED, POS_HALF_CLOSED:
		return true
	}
	return false
}
