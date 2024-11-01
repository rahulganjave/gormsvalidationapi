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
		ProductCode:            pgtype.Text{String: util.RandomProduct(), Valid: true}.String,
		ExchangeCode:           pgtype.Text{String: util.RandomExchange(), Valid: true}.String,
		InstrumentCode:         pgtype.Text{String: util.RandomUnderlying(), Valid: true}.String,
		PositionSide:           int32(util.RandomSide()),
		PositionStatus:         int32(util.RandomPositionStatus()),
		PositionType:           int32(util.RandomPositionType()),
		SettlementStatus:       int32(util.RandomSettlementStatus()),
		CalcType:               int32(1),
		OpenQty:                pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		OpenPx:                 pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		ClosedQty:              pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		ClosedPx:               pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		TradingCurr:            pgtype.Text{String: util.RandomCurrency(), Valid: true}.String,
		SettleCurr:             pgtype.Text{String: util.RandomCurrency(), Valid: true}.String,
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
		UnderlyingCode:         pgtype.Text{String: util.RandomUnderlying(), Valid: true},
		OpenClose:              int32(util.RandomOpenOrClose()),
		PutOrCall:              pgtype.Int4{Int32: int32(util.RandomPutOrCall()), Valid: true},
		ContractSize:           pgtype.Numeric{Int: big.NewInt(util.RandomMoney()), Valid: true},
	}
	position, err := testQueries.CreatePosition(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, position)

	require.Equal(t, params.InstrumentCode, position.InstrumentCode)
	require.Equal(t, params.ExchangeCode, position.ExchangeCode)
	require.Equal(t, params.ProductCode, position.ProductCode)

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
	require.Equal(t, position1.ProductCode, position2.ProductCode)
	require.Equal(t, position1.ExchangeCode, position2.ExchangeCode)
	require.Equal(t, position1.InstrumentCode, position2.InstrumentCode)
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
