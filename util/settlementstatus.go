package util

// Constants for all supported Settlementstatus
const (
	OUTSTANDING = 1
	SETTLED     = 2
)

// IsSupportedSettlementstatus returns true if the Settlementstatus is supported
func IsSupportedSettlementstatus(settlementstatus int) bool {
	switch settlementstatus {
	case OUTSTANDING, SETTLED:
		return true
	}
	return false
}
