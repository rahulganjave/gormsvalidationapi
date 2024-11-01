package util

// Constants for all supported posistionputorcall
const (
	PUT  = 0
	CALL = 1
)

// IsSupportedpositionputorcall returns true if the positionputorcall is supported
func IsSupportedpositionputorcall(positionputorcall int) bool {
	switch positionputorcall {
	case PUT, CALL:
		return true
	}
	return false
}
