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

func createRandomPosition(t *testing.T) Position {
	// Create a new Position
	params := CreatePositionParams{
		PositionID:             pgtype.Text{String: util.GenerateRandomID(), Valid: true}.String,
		RefPositionID:          pgtype.Text{},
		TradingAcct:            pgtype.Text{String: util.RandomAccountNo(), Valid: true}.String,
		ClientID:               pgtype.Text{String: "PSPL", Valid: true}.String,
		Product:                pgtype.Text{String: util.RandomProduct(), Valid: true}.String,
		Exchange:               pgtype.Text{String: util.RandomExchange(), Valid: true}.String,
		Instrument:             pgtype.Text{String: util.RandomUnderlying(), Valid: true}.String,
		Side:                   pgtype.Text{String: util.RandomSide(), Valid: true},
		Status:                 pgtype.Text{String: util.RandomPositionStatus(), Valid: true},
		PositionType:           pgtype.Text{String: util.RandomPositionType(), Valid: true},
		SettlementStatus:       pgtype.Text{String: util.RandomSettlementStatus(), Valid: true},
		CalcType:               pgtype.Text{},
		OpenQty:                pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		OpenPx:                 pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		ClosedQty:              pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		ClosedPx:               pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		TradingCurr:            pgtype.Text{String: util.RandomCurrency(), Valid: true},
		SettleCurr:             pgtype.Text{String: util.RandomCurrency(), Valid: true},
		MarkPrice:              pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		MarkValue:              pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		ExchangeRate:           pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		UnrealizedPl:           pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		RealizedPl:             pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		MmAndCash:              pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		CollateralHaircutRatio: pgtype.Numeric{Int: big.NewInt(int64(util.RandomFloat(0, 1) * 100)), Valid: true},
		CollateralHaircut:      pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		OpenDate:               pgtype.Timestamptz{Time: time.Now(), Valid: true},
		ClosedDate:             pgtype.Timestamptz{Time: time.Now(), Valid: true},
		CostToClose:            pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		TradeDate:              pgtype.Date{Time: util.RandomDate(2020, 2030), Valid: true},
		TradingSessionID:       pgtype.Text{String: util.RandomExchange() + util.RandomProduct(), Valid: true},
		Underlying:             pgtype.Text{String: util.RandomUnderlying(), Valid: true},
		OpenClose:              pgtype.Text{String: util.RandomOpenOrClose(), Valid: true},
		PutOrCall:              pgtype.Text{String: util.RandomPutOrCall(), Valid: true},
		ContractSize:           pgtype.Numeric{Int: big.NewInt(util.RandomMoney()), Valid: true},
	}
	position, err := testQueries.CreatePosition(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, position)

	require.Equal(t, params.Instrument, position.Instrument)
	require.Equal(t, params.Exchange, position.Exchange)
	require.Equal(t, params.Product, position.Product)

	//require.NotZero(t, account.ID)
	//require.NotZero(t, account.CreatedAt)

	return position
}

func TestCreatePosition(t *testing.T) {
	createRandomPosition(t)
}

func TestGetPosition(t *testing.T) {
	position1 := createRandomPosition(t)

	position2, err := testQueries.GetPosition(context.Background(), position1.PositionID)
	require.NoError(t, err)
	require.NotEmpty(t, position2)

	require.Equal(t, position1.PositionID, position2.PositionID)
	require.WithinDuration(t, position1.CreatedAt, position2.CreatedAt, time.Second)
}

func TestUpdatePosition(t *testing.T) {
	position1 := createRandomPosition(t)
	update_params := UpdatePositionParams{
		PositionID: position1.PositionID,
		OpenQty:    position1.OpenPx,
		OpenPx:     position1.OpenPx,
	}

	position2, err := testQueries.UpdatePosition(context.Background(), update_params)
	require.NoError(t, err)
	require.NotEmpty(t, position1)

	require.Equal(t, position1.PositionID, position2.PositionID)
	require.Equal(t, position1.Product, position2.Product)
	require.Equal(t, position1.Exchange, position2.Exchange)
	require.Equal(t, position1.Instrument, position2.Instrument)
	require.Equal(t, position1.TradingCurr, position2.TradingCurr)
	require.WithinDuration(t, position1.CreatedAt, position2.CreatedAt, time.Second)
}

func TestDeletePosition(t *testing.T) {
	position1 := createRandomPosition(t)
	err := testQueries.DeletePosition(context.Background(), position1.PositionID)
	require.NoError(t, err)

	position2, err := testQueries.GetPosition(context.Background(), position1.PositionID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, position2)
}
