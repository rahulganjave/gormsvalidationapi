package util

// Constants for all supported posistionopenorclose
const (
	OPEN  = 'O'
	CLOSE = 'C'
)

// IsSupportedpositionopenorclose returns true if the positionopenorclose is supported
func IsSupportedpositionopenorclose(positionopenorclose rune) bool {
	switch positionopenorclose {
	case OPEN, CLOSE:
		return true
	}
	return false
}
