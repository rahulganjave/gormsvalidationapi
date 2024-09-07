package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomAccountNo() string {
	return strconv.FormatInt(RandomInt(0, 1000000000), 10)
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomDate generates a random date (year, month, day)
func RandomDate(minYear, maxYear int) time.Time {
	rand.Seed(time.Now().UnixNano())
	year := minYear + rand.Intn(maxYear-minYear+1)
	month := rand.Intn(12) + 1
	day := rand.Intn(28) + 1 // assuming all months have at least 28 days
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// RandomCurrency generates a random currency code
func RandomExchange() string {
	exchanges := []string{NASD, NYSE, AMEX, SEHK}
	n := len(exchanges)
	return exchanges[rand.Intn(n)]
}

// RandomProduct generates a random product code
func RandomProduct() string {
	products := []string{CS, OPT, FUR, CFD, BO}
	n := len(products)
	return products[rand.Intn(n)]
}

func RandomUnderlying() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	length := 3 + rand.Intn(2) // generate a random length between 3 and 5
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandomMarginType generates a random Margintype
func RandomMarginType() string {
	margintypes := []string{NonMarginProduct, PercentPerTradesValue, ValuePerLot, PercentPerTradeSize}
	n := len(margintypes)
	return margintypes[rand.Intn(n)]
}

// RandomCommissionType generates a random CommissionType
func RandomCommissionType() string {
	commisiontypes := []string{C_PercentPerTradesValue, C_ValuePerLot, C_FixedValuePerTrades}
	n := len(commisiontypes)
	return commisiontypes[rand.Intn(n)]
}

// RandomSide generates a random Side
func RandomSide() string {
	side := []string{BUY, SELL, SHORT_SELL}
	n := len(side)
	return side[rand.Intn(n)]
}

// RandomPositionType generates a random PositionType
func RandomPositionType() string {
	posistiontypes := []string{INTRADAY, HISTORICAL}
	n := len(posistiontypes)
	return posistiontypes[rand.Intn(n)]
}

// RandomPositionStatus generates a random PositionStatus
func RandomPositionStatus() string {
	posistionstatus := []string{POS_OPEN, POS_CLOSED, POS_HALF_CLOSED}
	n := len(posistionstatus)
	return posistionstatus[rand.Intn(n)]
}

// RandomOpenOrClose generates a random open_or_close
func RandomOpenOrClose() string {
	open_or_close := []string{OPEN, CLOSE}
	n := len(open_or_close)
	return open_or_close[rand.Intn(n)]
}

// RandomPutOrCall generates a random put_or_call
func RandomPutOrCall() string {
	put_or_call := []string{PUT, CALL}
	n := len(put_or_call)
	return put_or_call[rand.Intn(n)]
}

// RandomSettlementStatus generates a random Settlementstatus
func RandomSettlementStatus() string {
	settlementstatus := []string{OUTSTANDING, SETTLED}
	n := len(settlementstatus)
	return settlementstatus[rand.Intn(n)]
}

// GenerateRandomID generates a random ID
func GenerateRandomID() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, 20)
	for i := range b {
		b[i] = chars[r.Intn(len(chars))]
	}
	return string(b)
}

func RandomPositionID() string {
	return GenerateRandomID()
}
