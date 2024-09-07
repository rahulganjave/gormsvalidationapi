package util

// Constants for all supported posistionopenorclose
const (
	OPEN  = "0"
	CLOSE = "1"
)

// IsSupportedpositionopenorclose returns true if the positionopenorclose is supported
func IsSupportedpositionopenorclose(positionopenorclose string) bool {
	switch positionopenorclose {
	case OPEN, CLOSE:
		return true
	}
	return false
}
