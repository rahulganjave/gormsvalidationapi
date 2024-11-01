package util

// Constants for all supported authorizedordertype
const (
	C_Buy    = 1
	C_Sell   = 2
	C_Limit  = 3
	C_Market = 4
)

// IsSupportedAuthorizedOrderType returns true if the authorizedordertype is supported
func IsSupportedAuthorizedOrderType(authorizedordertype int) bool {
	switch authorizedordertype {
	case C_Buy, C_Sell, C_Limit, C_Market:
		return true
	}
	return false
}
