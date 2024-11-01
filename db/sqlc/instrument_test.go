package db

import (
	"context"
	"math/big"
	"strconv"
	"testing"
	"time"

	"github.com/gormsvalidationapi/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTestInstrument(t *testing.T) Instrument {
	arg := CreateInstrumentParams{
		ProductCode:       pgtype.Text{String: util.RandomProduct(), Valid: true}.String,
		ExchangeCode:      pgtype.Text{String: util.RandomExchange(), Valid: true}.String,
		InstrumentCode:    pgtype.Text{String: util.RandomUnderlying(), Valid: true}.String,
		UnderlyingCode:    pgtype.Text{String: util.RandomUnderlying(), Valid: true},
		InstrumentDesc:    pgtype.Text{String: "E-Mini S&P 500", Valid: true},
		MaturityMonthyear: pgtype.Text{String: "202305", Valid: true},
		MaturityDay:       pgtype.Text{String: "15", Valid: true},
		PutOrCall:         pgtype.Text{String: strconv.Itoa(util.RandomPutOrCall()), Valid: true},
		StrikePrice:       pgtype.Numeric{Int: big.NewInt(util.RandomMoney()), Valid: true},
		Symbol:            pgtype.Text{String: util.RandomUnderlying(), Valid: true},
		SymbolSfx:         pgtype.Text{String: "F", Valid: true},
		Cusip:             pgtype.Text{String: util.RandomUnderlying(), Valid: true},
		Isin:              pgtype.Text{String: util.RandomUnderlying(), Valid: true},
		Ric:               pgtype.Text{String: util.RandomUnderlying(), Valid: true},
		Sedol:             pgtype.Text{String: util.RandomUnderlying(), Valid: true},
		Blpsyn:            pgtype.Text{String: util.RandomUnderlying(), Valid: true},
		Cficode:           pgtype.Text{String: util.RandomUnderlying(), Valid: true},
		PriceCode:         pgtype.Text{String: util.RandomExchange() + util.RandomUnderlying(), Valid: true},
		EffectiveDate:     pgtype.Timestamp{Time: time.Now(), Valid: true},
		Minqty:            int32(util.RandomMoney()),
		Sector:            pgtype.Text{String: "Eq", Valid: true},
		ContractSize:      pgtype.Numeric{Int: big.NewInt(util.RandomMoney()), Valid: true},
		CorporateAction:   pgtype.Text{String: "None", Valid: true},
		TradingSessionID:  pgtype.Text{String: util.RandomExchange() + util.RandomProduct(), Valid: true},
		TradingCurr:       pgtype.Text{String: util.RandomCurrency(), Valid: true},
		PriceDec:          pgtype.Numeric{Int: big.NewInt(3.0), Valid: true},
		HandlInst:         pgtype.Text{String: "0", Valid: true},
		TickID:            int32(1),
		MarketSlippage:    pgtype.Numeric{Int: big.NewInt(0.0), Valid: true},
		RefPx:             pgtype.Numeric{Int: big.NewInt(10.0), Valid: true},
		RefPxType:         int16(1),
		Tplusone:          pgtype.Int4{1, true},
		CreatedAt:         time.Now(),
		UpdatedAt:         pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	instrument, err := testQueries.CreateInstrument(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, instrument)

	return instrument
}

func TestCreateInstrument(t *testing.T) {
	instrument := createTestInstrument(t)

	assert.Equal(t, instrument.ProductCode, instrument.ProductCode)
	assert.Equal(t, instrument.ExchangeCode, instrument.ExchangeCode)
	assert.Equal(t, instrument.InstrumentCode, instrument.InstrumentCode)
}

func TestGetInstrument(t *testing.T) {
	instrument := createTestInstrument(t)

	arg := GetInstrumentParams{
		ProductCode:    instrument.ProductCode,
		ExchangeCode:   instrument.ExchangeCode,
		InstrumentCode: instrument.InstrumentCode,
	}

	gotInstrument, err := testQueries.GetInstrument(context.Background(), arg)
	require.NoError(t, err)
	assert.Equal(t, instrument, gotInstrument)
}

/*
	func TestUpdateInstrument(t *testing.T) {
		instrument := createTestInstrument(t)

		arg := UpdateInstrumentParams{
			ProductCode:       instrument.ProductCode,
			UnderlyingCode:    pgtype.Text{String: util.RandomUnderlying(), Valid: true},
			InstrumentDesc:    pgtype.Text{String: "E-Mini S&P 500", Valid: true},
			MaturityMonthyear: pgtype.Text{String: "202305", Valid: true},
			MaturityDay:       pgtype.Text{String: "15", Valid: true},
			PutOrCall:         pgtype.Text{String: util.RandomPutOrCall(), Valid: true},
			StrikePrice:       pgtype.Numeric{Int: big.NewInt(util.RandomMoney()), Valid: true},
			Symbol:            pgtype.Text{String: util.RandomUnderlying(), Valid: true},
			SymbolSfx:         pgtype.Text{String: "F", Valid: true},
			Cusip:             pgtype.Text{String: util.RandomUnderlying(), Valid: true},
			Isin:              pgtype.Text{String: util.RandomUnderlying(), Valid: true},
			Ric:               pgtype.Text{String: util.RandomUnderlying(), Valid: true},
			Sedol:             pgtype.Text{String: util.RandomUnderlying(), Valid: true},
			Blpsyn:            pgtype.Text{String: util.RandomUnderlying(), Valid: true},
			Cficode:           pgtype.Text{String: util.RandomUnderlying(), Valid: true},
			PriceCode:         pgtype.Text{String: util.RandomExchange() + util.RandomUnderlying(), Valid: true},
			EffectiveDate:     pgtype.Timestamp{Time: time.Now(), Valid: true},
			Minqty:            pgtype.Int4{Int32: int32(util.RandomMoney()), Valid: false},
			Sector:            pgtype.Text{String: "Eq", Valid: true},
			ContractSize:      pgtype.Numeric{Int: big.NewInt(util.RandomMoney()), Valid: true},
			CorporateAction:   pgtype.Text{String: "None", Valid: true},
			TradingSessionID:  pgtype.Text{String: util.RandomExchange() + util.RandomProduct(), Valid: true},
			TradingCurr:       pgtype.Text{String: util.RandomCurrency(), Valid: true},
			PriceDec:          pgtype.Numeric{Int: big.NewInt(3.0), Valid: true},
			HandlInst:         pgtype.Text{String: "0", Valid: true},
			TickID:            pgtype.Int4{1, true},
			MarketSlippage:    pgtype.Numeric{Int: big.NewInt(0.0), Valid: true},
			RefPx:             pgtype.Numeric{Int: big.NewInt(10.0), Valid: true},
			RefPxType:         pgtype.Int4{1, true},
			Tplusone:          pgtype.Int4{0, true},
		}

		updatedInstrument, err := testQueries.UpdateInstrument(context.Background(), arg)
		require.NoError(t, err)
		assert.Equal(t, instrument.ProductCode, updatedInstrument.ProductCode)
		assert.Equal(t, instrument.ExchangeCode, updatedInstrument.ExchangeCode)
		assert.Equal(t, instrument.InstrumentCode, updatedInstrument.InstrumentCode)
	}
*/
func TestDeleteInstrument(t *testing.T) {
	instrument := createTestInstrument(t)

	arg := DeleteInstrumentParams{
		ProductCode:    instrument.ProductCode,
		ExchangeCode:   instrument.ExchangeCode,
		InstrumentCode: instrument.InstrumentCode,
	}

	err := testQueries.DeleteInstrument(context.Background(), arg)
	require.NoError(t, err)
}

func TestListInstruments(t *testing.T) {
	createTestInstrument(t)

	arg := ListInstrumentsParams{
		Limit:  10,
		Offset: 0,
	}

	instruments, err := testQueries.ListInstruments(context.Background(), arg)
	require.NoError(t, err)
	assert.NotEmpty(t, instruments)
}
