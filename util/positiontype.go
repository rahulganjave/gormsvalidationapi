package util

// Constants for all supported positiontype
const (
	INTRADAY   = "1"
	HISTORICAL = "2"
)

// IsSupportedpositiontype returns true if the positiontype is supported
func IsSupportedPositiontype(positiontype string) bool {
	switch positiontype {
	case INTRADAY, HISTORICAL:
		return true
	}
	return false
}
