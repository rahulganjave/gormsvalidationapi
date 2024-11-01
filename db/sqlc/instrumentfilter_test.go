package db

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/gormsvalidationapi/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomInstrumentfilter(t *testing.T) Instrumentfilter {
	// Create a new Instrumentfilter
	params := CreateInstrumentfilterParams{
		ProductCode:         pgtype.Text{String: util.RandomProduct(), Valid: true}.String,
		ExchangeCode:        pgtype.Text{String: util.RandomExchange(), Valid: true}.String,
		InstrumentCode:      pgtype.Text{String: util.RandomUnderlying(), Valid: true}.String,
		UnderlyingCode:      pgtype.Text{String: util.RandomUnderlying(), Valid: true},
		FilterOption:        int32(1),
		MaxOrderQty:         pgtype.Numeric{Int: big.NewInt(100)},
		MaxOrderPrice:       pgtype.Numeric{Int: big.NewInt(100.0)},
		MaxOrderValue:       pgtype.Numeric{Int: big.NewInt(1000.0)},
		BuyLimit:            pgtype.Numeric{Int: big.NewInt(10.0)},
		SellLimit:           pgtype.Numeric{Int: big.NewInt(20.0)},
		ExposureFactor:      pgtype.Numeric{Int: big.NewInt(int64(0.5 * 100))},
		BuyExposure:         pgtype.Numeric{Int: big.NewInt(10.0)},
		SellExposure:        pgtype.Numeric{Int: big.NewInt(20.0)},
		BuyQtyLimit:         pgtype.Numeric{Int: big.NewInt(10)},
		BuyQtyExposure:      pgtype.Numeric{Int: big.NewInt(10.0)},
		SellQtyLimit:        pgtype.Numeric{Int: big.NewInt(20)},
		SellQtyExposure:     pgtype.Numeric{Int: big.NewInt(20.0)},
		NetQtyLimit:         pgtype.Numeric{Int: big.NewInt(10)},
		NetQtyExposure:      pgtype.Numeric{Int: big.NewInt(10.0)},
		MarginType:          int32(util.RandomMarginType()),
		InitialMargin:       pgtype.Numeric{Int: big.NewInt(10.0)},
		MaintenanceMargin:   pgtype.Numeric{Int: big.NewInt(20.0)},
		MarginableRatio:     pgtype.Numeric{Int: big.NewInt(int64(0.5 * 100))},
		MarginCurr:          pgtype.Text{String: util.RandomCurrency(), Valid: true},
		CommissionType:      int32(util.RandomCommissionType()),
		Commission:          pgtype.Numeric{Int: big.NewInt(10.0)},
		MinCommission:       pgtype.Numeric{Int: big.NewInt(5.0)},
		CommissionCurr:      pgtype.Text{String: util.RandomCurrency(), Valid: true},
		AuthorizedOrderType: int64(util.RandomAuthorizedOrderType()),
		TradingCurr:         pgtype.Text{String: util.RandomCurrency(), Valid: true},
		PriceCode:           pgtype.Text{String: util.RandomExchange() + util.RandomUnderlying(), Valid: true},
		RefPx:               pgtype.Numeric{Int: big.NewInt(10.0)},
		RefPxType:           int16(1),
		ContractSize:        pgtype.Numeric{Int: big.NewInt(100)},
	}

	instrumentfilter, err := testQueries.CreateInstrumentfilter(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, instrumentfilter)

	require.Equal(t, params.InstrumentCode, instrumentfilter.InstrumentCode)
	require.Equal(t, params.ExchangeCode, instrumentfilter.ExchangeCode)
	require.Equal(t, params.ProductCode, instrumentfilter.ProductCode)

	//require.NotZero(t, account.ID)
	//require.NotZero(t, account.CreatedAt)

	return instrumentfilter
}

func TestCreateInstrumentfilter(t *testing.T) {
	createRandomInstrumentfilter(t)
}

/*
func TestGetInstrumentfilter(t *testing.T) {
	instrumentfilter1 := createRandomInstrumentfilter(t)
	params := GetInstrumentfilterParams{
		ProductCode:    pgtype.Text{String: util.RandomProduct(), Valid: true}.String,
		ExchangeCode:   pgtype.Text{String: util.RandomExchange(), Valid: true}.String,
		InstrumentCode: pgtype.Text{String: util.RandomUnderlying(), Valid: true}.String,
	}
	instrumentfilter2, err := testQueries.GetInstrumentfilter(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, instrumentfilter2)

	require.Equal(t, instrumentfilter1.ProductCode, instrumentfilter2.ProductCode)
	require.Equal(t, instrumentfilter1.ExchangeCode, instrumentfilter2.ExchangeCode)
	require.Equal(t, instrumentfilter1.InstrumentCode, instrumentfilter2.InstrumentCode)
	require.Equal(t, instrumentfilter1.TradingCurr, instrumentfilter2.TradingCurr)
	require.WithinDuration(t, instrumentfilter1.CreatedAt, instrumentfilter2.CreatedAt, time.Second)
}*/

func TestUpdateInstrumentfilter(t *testing.T) {
	instrumentfilter1 := createRandomInstrumentfilter(t)

	params := UpdateInstrumentfilterParams{
		ProductCode:    instrumentfilter1.ProductCode,
		ExchangeCode:   instrumentfilter1.ExchangeCode,
		InstrumentCode: instrumentfilter1.InstrumentCode,
		BuyExposure:    instrumentfilter1.BuyExposure,
		SellExposure:   instrumentfilter1.SellExposure,
	}

	instrumentfilter2, err := testQueries.UpdateInstrumentfilter(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, instrumentfilter2)

	require.Equal(t, instrumentfilter1.ProductCode, instrumentfilter2.ProductCode)
	require.Equal(t, instrumentfilter1.ExchangeCode, instrumentfilter2.ExchangeCode)
	require.Equal(t, instrumentfilter1.InstrumentCode, instrumentfilter2.InstrumentCode)
	require.Equal(t, instrumentfilter1.TradingCurr, instrumentfilter2.TradingCurr)
	require.WithinDuration(t, instrumentfilter1.CreatedAt, instrumentfilter2.CreatedAt, time.Second)
}

func TestDeleteInstrumentfilter(t *testing.T) {
	instrumentfilter1 := createRandomInstrumentfilter(t)
	delete_params := DeleteInstrumentfilterParams{
		ProductCode:    instrumentfilter1.ProductCode,
		ExchangeCode:   instrumentfilter1.ExchangeCode,
		InstrumentCode: instrumentfilter1.InstrumentCode,
	}
	err := testQueries.DeleteInstrumentfilter(context.Background(), delete_params)
	require.NoError(t, err)

	params := GetInstrumentfilterParams{
		ProductCode:    pgtype.Text{String: util.RandomProduct(), Valid: true}.String,
		ExchangeCode:   pgtype.Text{String: util.RandomExchange(), Valid: true}.String,
		InstrumentCode: pgtype.Text{String: util.RandomUnderlying(), Valid: true}.String,
	}
	instrumentfilter2, err := testQueries.GetInstrumentfilter(context.Background(), params)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, instrumentfilter2)
}
