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

func createRandomAccount(t *testing.T) Accountfilter {
	// Create a new account
	params := CreateAccountParams{
		ClientID:             pgtype.Text{String: "PSPL", Valid: true},
		Account:              pgtype.Text{String: util.RandomAccountNo(), Valid: true}.String,
		Trader:               pgtype.Text{String: "trader-1", Valid: true},
		CashBalanceBf:        pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		CashBalance:          pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		EstTransCost:         pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		EquityRealizedPl:     pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		EquityUnrealizedPl:   pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		MarginRealizedPl:     pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		MarginUnrealizedPl:   pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		EstClosingCost:       pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		TransactionNotBooked: pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		EquityOrder:          pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		AccountValue:         pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		DepositWithdraw:      pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		CollateralHaircut:    pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		EquityPositionValue:  pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		MarginRequirement:    pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		Currency:             pgtype.Text{String: util.RandomCurrency(), Valid: true},
		BuyLimit:             pgtype.Numeric{Int: big.NewInt(util.RandomMoney()), Valid: true},
		SellLimit:            pgtype.Numeric{Int: big.NewInt(util.RandomMoney()), Valid: true},
		BuyExposure:          pgtype.Numeric{Int: big.NewInt(util.RandomMoney()), Valid: true},
		SellExposure:         pgtype.Numeric{Int: big.NewInt(util.RandomMoney()), Valid: true},
		LoanLimit:            pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		MarginFactor:         pgtype.Numeric{Int: big.NewInt(0), Valid: true},
		AuthorizedOrderType:  pgtype.Int8{1, true},
		CalcMode:             pgtype.Text{String: "mode-1", Valid: true},
		LimitMode:            pgtype.Int4{1, true},
		Action:               pgtype.Text{String: "action-1", Valid: true},
		NettingMode:          "netting-mode-1",
		MaxOrdQty:            pgtype.Numeric{Int: big.NewInt(10), Valid: true},
		MaxOrdVal:            pgtype.Numeric{Int: big.NewInt(100), Valid: true},
		NoOfTransaction:      0,
		TransactionInterval:  0,
	}

	account, err := testQueries.CreateAccount(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, params.ClientID, account.ClientID)
	require.Equal(t, params.Account, account.Account)
	require.Equal(t, params.Trader, account.Trader)

	//require.NotZero(t, account.ID)
	//require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	params := GetAccountParams{
		ClientID: pgtype.Text{String: "PSPL", Valid: true},
		Account:  pgtype.Text{String: util.RandomAccountNo(), Valid: true}.String,
	}
	account2, err := testQueries.GetAccount(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.Account, account2.Account)
	require.Equal(t, account1.ClientID, account2.ClientID)
	require.Equal(t, account1.Trader, account2.Trader)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	params := UpdateAccountParams{
		BuyExposure:  account1.BuyExposure,
		SellExposure: account1.SellExposure,
	}

	account2, err := testQueries.UpdateAccount(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.Account, account2.Account)
	require.Equal(t, account1.ClientID, account2.ClientID)
	//require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.Account)
	require.NoError(t, err)

	params := GetAccountParams{
		ClientID: pgtype.Text{String: "PSPL", Valid: true},
		Account:  pgtype.Text{String: util.RandomAccountNo(), Valid: true}.String,
	}
	account2, err := testQueries.GetAccount(context.Background(), params)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, account2)
}
