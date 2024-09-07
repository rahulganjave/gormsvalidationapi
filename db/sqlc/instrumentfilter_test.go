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
		Product:             pgtype.Text{String: util.RandomProduct(), Valid: true}.String,
		Exchange:            pgtype.Text{String: util.RandomExchange(), Valid: true}.String,
		Instrument:          pgtype.Text{String: util.RandomUnderlying(), Valid: true}.String,
		Underlying:          pgtype.Text{String: util.RandomUnderlying(), Valid: true},
		Filter:              pgtype.Int4{1, true},
		MaxOrderQty:         pgtype.Numeric{Int: big.NewInt(100)},
		MaxOrderPrice:       pgtype.Numeric{Int: big.NewInt(100.0)},
		MaxOrderValue:       pgtype.Numeric{Int: big.NewInt(1000.0)},
		BuyLmt:              pgtype.Numeric{Int: big.NewInt(10.0)},
		SellLmt:             pgtype.Numeric{Int: big.NewInt(20.0)},
		ExposureFactor:      pgtype.Numeric{Int: big.NewInt(int64(0.5 * 100))},
		BuyExposure:         pgtype.Numeric{Int: big.NewInt(10.0)},
		SellExposure:        pgtype.Numeric{Int: big.NewInt(20.0)},
		BuyQtyLmt:           pgtype.Numeric{Int: big.NewInt(10)},
		BuyQtyExposure:      pgtype.Numeric{Int: big.NewInt(10.0)},
		SellQtyLmt:          pgtype.Numeric{Int: big.NewInt(20)},
		SellQtyExposure:     pgtype.Numeric{Int: big.NewInt(20.0)},
		NetQtyLmt:           pgtype.Numeric{Int: big.NewInt(10)},
		NetQtyExposure:      pgtype.Numeric{Int: big.NewInt(10.0)},
		MarginType:          pgtype.Text{String: util.RandomMarginType(), Valid: true},
		Im:                  pgtype.Numeric{Int: big.NewInt(10.0)},
		Mm:                  pgtype.Numeric{Int: big.NewInt(20.0)},
		MarginableRatio:     pgtype.Numeric{Int: big.NewInt(int64(0.5 * 100))},
		MarginCurr:          pgtype.Text{String: util.RandomCurrency(), Valid: true},
		CommissionType:      pgtype.Text{String: util.RandomCommissionType(), Valid: true},
		Commission:          pgtype.Numeric{Int: big.NewInt(10.0)},
		MinCommission:       pgtype.Numeric{Int: big.NewInt(5.0)},
		CommissionCurr:      pgtype.Text{String: util.RandomCurrency(), Valid: true},
		AuthorizedOrderType: pgtype.Int8{1, true},
		TradingCurr:         pgtype.Text{String: util.RandomCurrency(), Valid: true},
		PriceCode:           pgtype.Text{String: util.RandomExchange() + util.RandomUnderlying(), Valid: true},
		RefPx:               pgtype.Numeric{Int: big.NewInt(10.0)},
		RefPxType:           pgtype.Int2{1, true},
		ContractSize:        pgtype.Numeric{Int: big.NewInt(100)},
	}

	instrumentfilter, err := testQueries.CreateInstrumentfilter(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, instrumentfilter)

	require.Equal(t, params.Instrument, instrumentfilter.Instrument)
	require.Equal(t, params.Exchange, instrumentfilter.Exchange)
	require.Equal(t, params.Product, instrumentfilter.Product)

	//require.NotZero(t, account.ID)
	//require.NotZero(t, account.CreatedAt)

	return instrumentfilter
}

func TestCreateInstrumentfilter(t *testing.T) {
	createRandomInstrumentfilter(t)
}

func TestGetInstrumentfilter(t *testing.T) {
	instrumentfilter1 := createRandomInstrumentfilter(t)
	params := GetInstrumentfilterParams{
		Product:    pgtype.Text{String: util.RandomProduct(), Valid: true}.String,
		Exchange:   pgtype.Text{String: util.RandomExchange(), Valid: true}.String,
		Instrument: pgtype.Text{String: util.RandomUnderlying(), Valid: true}.String,
	}
	instrumentfilter2, err := testQueries.GetInstrumentfilter(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, instrumentfilter2)

	require.Equal(t, instrumentfilter1.Product, instrumentfilter2.Product)
	require.Equal(t, instrumentfilter1.Exchange, instrumentfilter2.Exchange)
	require.Equal(t, instrumentfilter1.Instrument, instrumentfilter2.Instrument)
	require.Equal(t, instrumentfilter1.TradingCurr, instrumentfilter2.TradingCurr)
	require.WithinDuration(t, instrumentfilter1.CreatedAt, instrumentfilter2.CreatedAt, time.Second)
}

func TestUpdateInstrumentfilter(t *testing.T) {
	instrumentfilter1 := createRandomInstrumentfilter(t)

	params := UpdateInstrumentfilterParams{
		Product:      instrumentfilter1.Product,
		Exchange:     instrumentfilter1.Exchange,
		Instrument:   instrumentfilter1.Instrument,
		BuyExposure:  instrumentfilter1.BuyExposure,
		SellExposure: instrumentfilter1.SellExposure,
	}

	instrumentfilter2, err := testQueries.UpdateInstrumentfilter(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, instrumentfilter2)

	require.Equal(t, instrumentfilter1.Product, instrumentfilter2.Product)
	require.Equal(t, instrumentfilter1.Exchange, instrumentfilter2.Exchange)
	require.Equal(t, instrumentfilter1.Instrument, instrumentfilter2.Instrument)
	require.Equal(t, instrumentfilter1.TradingCurr, instrumentfilter2.TradingCurr)
	require.WithinDuration(t, instrumentfilter1.CreatedAt, instrumentfilter2.CreatedAt, time.Second)
}

func TestDeleteInstrumentfilter(t *testing.T) {
	instrumentfilter1 := createRandomInstrumentfilter(t)
	delete_params := DeleteInstrumentfilterParams{
		Product:    instrumentfilter1.Product,
		Exchange:   instrumentfilter1.Exchange,
		Instrument: instrumentfilter1.Instrument,
	}
	err := testQueries.DeleteInstrumentfilter(context.Background(), delete_params)
	require.NoError(t, err)

	params := GetInstrumentfilterParams{
		Product:    pgtype.Text{String: util.RandomProduct(), Valid: true}.String,
		Exchange:   pgtype.Text{String: util.RandomExchange(), Valid: true}.String,
		Instrument: pgtype.Text{String: util.RandomUnderlying(), Valid: true}.String,
	}
	instrumentfilter2, err := testQueries.GetInstrumentfilter(context.Background(), params)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, instrumentfilter2)
}
