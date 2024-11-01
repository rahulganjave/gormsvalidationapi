// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accountfilter.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accountfilter (
    "client_id",
    "trading_acct",
    "trader_id",
    "cash_balance_bf",
    "cash_balance",
    "est_trans_cost",
    "equity_realized_pl",
    "equity_unrealized_pl",
    "margin_realized_pl",
    "margin_unrealized_pl",
    "est_closing_cost",
    "transaction_not_booked",
    "equity_order",
    "account_value",
    "deposit_withdraw",
    "collateral_haircut",
    "equity_position_value",
    "margin_requirement",
    "currency",
    "buy_limit",
    "sell_limit",
    "buy_exposure",
    "sell_exposure",
    "loan_limit",
    "net_loss_limit",
    "net_exposure_limit",
    "margin_factor",
    "authorized_order_type",
    "calc_mode",
    "limit_mode",
    "action_mode",
    "netting_mode",
    "max_ord_qty",
    "max_ord_val",
    "no_of_transaction",
    "transaction_interval"
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36
)
RETURNING client_id, trading_acct, trader_id, cash_balance_bf, cash_balance, est_trans_cost, equity_realized_pl, equity_unrealized_pl, margin_realized_pl, margin_unrealized_pl, est_closing_cost, transaction_not_booked, equity_order, account_value, deposit_withdraw, collateral_haircut, equity_position_value, margin_requirement, currency, buy_limit, sell_limit, buy_exposure, sell_exposure, loan_limit, net_loss_limit, net_exposure_limit, margin_factor, authorized_order_type, calc_mode, limit_mode, action_mode, netting_mode, max_ord_qty, max_ord_val, no_of_transaction, transaction_interval, created_at, updated_at
`

type CreateAccountParams struct {
	ClientID             pgtype.Text    `json:"client_id"`
	TradingAcct          string         `json:"trading_acct"`
	TraderID             pgtype.Text    `json:"trader_id"`
	CashBalanceBf        pgtype.Numeric `json:"cash_balance_bf"`
	CashBalance          pgtype.Numeric `json:"cash_balance"`
	EstTransCost         pgtype.Numeric `json:"est_trans_cost"`
	EquityRealizedPl     pgtype.Numeric `json:"equity_realized_pl"`
	EquityUnrealizedPl   pgtype.Numeric `json:"equity_unrealized_pl"`
	MarginRealizedPl     pgtype.Numeric `json:"margin_realized_pl"`
	MarginUnrealizedPl   pgtype.Numeric `json:"margin_unrealized_pl"`
	EstClosingCost       pgtype.Numeric `json:"est_closing_cost"`
	TransactionNotBooked pgtype.Numeric `json:"transaction_not_booked"`
	EquityOrder          pgtype.Numeric `json:"equity_order"`
	AccountValue         pgtype.Numeric `json:"account_value"`
	DepositWithdraw      pgtype.Numeric `json:"deposit_withdraw"`
	CollateralHaircut    pgtype.Numeric `json:"collateral_haircut"`
	EquityPositionValue  pgtype.Numeric `json:"equity_position_value"`
	MarginRequirement    pgtype.Numeric `json:"margin_requirement"`
	Currency             pgtype.Text    `json:"currency"`
	BuyLimit             pgtype.Numeric `json:"buy_limit"`
	SellLimit            pgtype.Numeric `json:"sell_limit"`
	BuyExposure          pgtype.Numeric `json:"buy_exposure"`
	SellExposure         pgtype.Numeric `json:"sell_exposure"`
	LoanLimit            pgtype.Numeric `json:"loan_limit"`
	NetLossLimit         pgtype.Numeric `json:"net_loss_limit"`
	NetExposureLimit     pgtype.Numeric `json:"net_exposure_limit"`
	MarginFactor         pgtype.Numeric `json:"margin_factor"`
	AuthorizedOrderType  int64          `json:"authorized_order_type"`
	CalcMode             int32          `json:"calc_mode"`
	LimitMode            int32          `json:"limit_mode"`
	ActionMode           int32          `json:"action_mode"`
	NettingMode          int32          `json:"netting_mode"`
	MaxOrdQty            pgtype.Numeric `json:"max_ord_qty"`
	MaxOrdVal            pgtype.Numeric `json:"max_ord_val"`
	NoOfTransaction      int32          `json:"no_of_transaction"`
	TransactionInterval  int16          `json:"transaction_interval"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Accountfilter, error) {
	row := q.db.QueryRow(ctx, createAccount,
		arg.ClientID,
		arg.TradingAcct,
		arg.TraderID,
		arg.CashBalanceBf,
		arg.CashBalance,
		arg.EstTransCost,
		arg.EquityRealizedPl,
		arg.EquityUnrealizedPl,
		arg.MarginRealizedPl,
		arg.MarginUnrealizedPl,
		arg.EstClosingCost,
		arg.TransactionNotBooked,
		arg.EquityOrder,
		arg.AccountValue,
		arg.DepositWithdraw,
		arg.CollateralHaircut,
		arg.EquityPositionValue,
		arg.MarginRequirement,
		arg.Currency,
		arg.BuyLimit,
		arg.SellLimit,
		arg.BuyExposure,
		arg.SellExposure,
		arg.LoanLimit,
		arg.NetLossLimit,
		arg.NetExposureLimit,
		arg.MarginFactor,
		arg.AuthorizedOrderType,
		arg.CalcMode,
		arg.LimitMode,
		arg.ActionMode,
		arg.NettingMode,
		arg.MaxOrdQty,
		arg.MaxOrdVal,
		arg.NoOfTransaction,
		arg.TransactionInterval,
	)
	var i Accountfilter
	err := row.Scan(
		&i.ClientID,
		&i.TradingAcct,
		&i.TraderID,
		&i.CashBalanceBf,
		&i.CashBalance,
		&i.EstTransCost,
		&i.EquityRealizedPl,
		&i.EquityUnrealizedPl,
		&i.MarginRealizedPl,
		&i.MarginUnrealizedPl,
		&i.EstClosingCost,
		&i.TransactionNotBooked,
		&i.EquityOrder,
		&i.AccountValue,
		&i.DepositWithdraw,
		&i.CollateralHaircut,
		&i.EquityPositionValue,
		&i.MarginRequirement,
		&i.Currency,
		&i.BuyLimit,
		&i.SellLimit,
		&i.BuyExposure,
		&i.SellExposure,
		&i.LoanLimit,
		&i.NetLossLimit,
		&i.NetExposureLimit,
		&i.MarginFactor,
		&i.AuthorizedOrderType,
		&i.CalcMode,
		&i.LimitMode,
		&i.ActionMode,
		&i.NettingMode,
		&i.MaxOrdQty,
		&i.MaxOrdVal,
		&i.NoOfTransaction,
		&i.TransactionInterval,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accountfilter
WHERE trading_acct = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, tradingAcct string) error {
	_, err := q.db.Exec(ctx, deleteAccount, tradingAcct)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT client_id, trading_acct, trader_id, cash_balance_bf, cash_balance, est_trans_cost, equity_realized_pl, equity_unrealized_pl, margin_realized_pl, margin_unrealized_pl, est_closing_cost, transaction_not_booked, equity_order, account_value, deposit_withdraw, collateral_haircut, equity_position_value, margin_requirement, currency, buy_limit, sell_limit, buy_exposure, sell_exposure, loan_limit, net_loss_limit, net_exposure_limit, margin_factor, authorized_order_type, calc_mode, limit_mode, action_mode, netting_mode, max_ord_qty, max_ord_val, no_of_transaction, transaction_interval, created_at, updated_at
FROM accountfilter
WHERE client_id = $1
    AND trading_acct = $2
LIMIT 1
`

type GetAccountParams struct {
	ClientID    pgtype.Text `json:"client_id"`
	TradingAcct string      `json:"trading_acct"`
}

func (q *Queries) GetAccount(ctx context.Context, arg GetAccountParams) (Accountfilter, error) {
	row := q.db.QueryRow(ctx, getAccount, arg.ClientID, arg.TradingAcct)
	var i Accountfilter
	err := row.Scan(
		&i.ClientID,
		&i.TradingAcct,
		&i.TraderID,
		&i.CashBalanceBf,
		&i.CashBalance,
		&i.EstTransCost,
		&i.EquityRealizedPl,
		&i.EquityUnrealizedPl,
		&i.MarginRealizedPl,
		&i.MarginUnrealizedPl,
		&i.EstClosingCost,
		&i.TransactionNotBooked,
		&i.EquityOrder,
		&i.AccountValue,
		&i.DepositWithdraw,
		&i.CollateralHaircut,
		&i.EquityPositionValue,
		&i.MarginRequirement,
		&i.Currency,
		&i.BuyLimit,
		&i.SellLimit,
		&i.BuyExposure,
		&i.SellExposure,
		&i.LoanLimit,
		&i.NetLossLimit,
		&i.NetExposureLimit,
		&i.MarginFactor,
		&i.AuthorizedOrderType,
		&i.CalcMode,
		&i.LimitMode,
		&i.ActionMode,
		&i.NettingMode,
		&i.MaxOrdQty,
		&i.MaxOrdVal,
		&i.NoOfTransaction,
		&i.TransactionInterval,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT client_id, trading_acct, trader_id, cash_balance_bf, cash_balance, est_trans_cost, equity_realized_pl, equity_unrealized_pl, margin_realized_pl, margin_unrealized_pl, est_closing_cost, transaction_not_booked, equity_order, account_value, deposit_withdraw, collateral_haircut, equity_position_value, margin_requirement, currency, buy_limit, sell_limit, buy_exposure, sell_exposure, loan_limit, net_loss_limit, net_exposure_limit, margin_factor, authorized_order_type, calc_mode, limit_mode, action_mode, netting_mode, max_ord_qty, max_ord_val, no_of_transaction, transaction_interval, created_at, updated_at
FROM accountfilter
WHERE client_id = $1
AND trading_acct = $2
LIMIT 1
FOR NO KEY UPDATE
`

type GetAccountForUpdateParams struct {
	ClientID    pgtype.Text `json:"client_id"`
	TradingAcct string      `json:"trading_acct"`
}

func (q *Queries) GetAccountForUpdate(ctx context.Context, arg GetAccountForUpdateParams) (Accountfilter, error) {
	row := q.db.QueryRow(ctx, getAccountForUpdate, arg.ClientID, arg.TradingAcct)
	var i Accountfilter
	err := row.Scan(
		&i.ClientID,
		&i.TradingAcct,
		&i.TraderID,
		&i.CashBalanceBf,
		&i.CashBalance,
		&i.EstTransCost,
		&i.EquityRealizedPl,
		&i.EquityUnrealizedPl,
		&i.MarginRealizedPl,
		&i.MarginUnrealizedPl,
		&i.EstClosingCost,
		&i.TransactionNotBooked,
		&i.EquityOrder,
		&i.AccountValue,
		&i.DepositWithdraw,
		&i.CollateralHaircut,
		&i.EquityPositionValue,
		&i.MarginRequirement,
		&i.Currency,
		&i.BuyLimit,
		&i.SellLimit,
		&i.BuyExposure,
		&i.SellExposure,
		&i.LoanLimit,
		&i.NetLossLimit,
		&i.NetExposureLimit,
		&i.MarginFactor,
		&i.AuthorizedOrderType,
		&i.CalcMode,
		&i.LimitMode,
		&i.ActionMode,
		&i.NettingMode,
		&i.MaxOrdQty,
		&i.MaxOrdVal,
		&i.NoOfTransaction,
		&i.TransactionInterval,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAccount = `-- name: ListAccount :many
SELECT client_id, trading_acct, trader_id, cash_balance_bf, cash_balance, est_trans_cost, equity_realized_pl, equity_unrealized_pl, margin_realized_pl, margin_unrealized_pl, est_closing_cost, transaction_not_booked, equity_order, account_value, deposit_withdraw, collateral_haircut, equity_position_value, margin_requirement, currency, buy_limit, sell_limit, buy_exposure, sell_exposure, loan_limit, net_loss_limit, net_exposure_limit, margin_factor, authorized_order_type, calc_mode, limit_mode, action_mode, netting_mode, max_ord_qty, max_ord_val, no_of_transaction, transaction_interval, created_at, updated_at
FROM accountfilter
WHERE client_id = $1
AND trading_acct = $2
LIMIT $2
OFFSET $3
`

type ListAccountParams struct {
	ClientID pgtype.Text `json:"client_id"`
	Limit    int32       `json:"limit"`
	Offset   int32       `json:"offset"`
}

func (q *Queries) ListAccount(ctx context.Context, arg ListAccountParams) ([]Accountfilter, error) {
	rows, err := q.db.Query(ctx, listAccount, arg.ClientID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Accountfilter{}
	for rows.Next() {
		var i Accountfilter
		if err := rows.Scan(
			&i.ClientID,
			&i.TradingAcct,
			&i.TraderID,
			&i.CashBalanceBf,
			&i.CashBalance,
			&i.EstTransCost,
			&i.EquityRealizedPl,
			&i.EquityUnrealizedPl,
			&i.MarginRealizedPl,
			&i.MarginUnrealizedPl,
			&i.EstClosingCost,
			&i.TransactionNotBooked,
			&i.EquityOrder,
			&i.AccountValue,
			&i.DepositWithdraw,
			&i.CollateralHaircut,
			&i.EquityPositionValue,
			&i.MarginRequirement,
			&i.Currency,
			&i.BuyLimit,
			&i.SellLimit,
			&i.BuyExposure,
			&i.SellExposure,
			&i.LoanLimit,
			&i.NetLossLimit,
			&i.NetExposureLimit,
			&i.MarginFactor,
			&i.AuthorizedOrderType,
			&i.CalcMode,
			&i.LimitMode,
			&i.ActionMode,
			&i.NettingMode,
			&i.MaxOrdQty,
			&i.MaxOrdVal,
			&i.NoOfTransaction,
			&i.TransactionInterval,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accountfilter
SET buy_exposure = $2,
sell_exposure = $3
WHERE trading_acct = $1
RETURNING client_id, trading_acct, trader_id, cash_balance_bf, cash_balance, est_trans_cost, equity_realized_pl, equity_unrealized_pl, margin_realized_pl, margin_unrealized_pl, est_closing_cost, transaction_not_booked, equity_order, account_value, deposit_withdraw, collateral_haircut, equity_position_value, margin_requirement, currency, buy_limit, sell_limit, buy_exposure, sell_exposure, loan_limit, net_loss_limit, net_exposure_limit, margin_factor, authorized_order_type, calc_mode, limit_mode, action_mode, netting_mode, max_ord_qty, max_ord_val, no_of_transaction, transaction_interval, created_at, updated_at
`

type UpdateAccountParams struct {
	TradingAcct  string         `json:"trading_acct"`
	BuyExposure  pgtype.Numeric `json:"buy_exposure"`
	SellExposure pgtype.Numeric `json:"sell_exposure"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Accountfilter, error) {
	row := q.db.QueryRow(ctx, updateAccount, arg.TradingAcct, arg.BuyExposure, arg.SellExposure)
	var i Accountfilter
	err := row.Scan(
		&i.ClientID,
		&i.TradingAcct,
		&i.TraderID,
		&i.CashBalanceBf,
		&i.CashBalance,
		&i.EstTransCost,
		&i.EquityRealizedPl,
		&i.EquityUnrealizedPl,
		&i.MarginRealizedPl,
		&i.MarginUnrealizedPl,
		&i.EstClosingCost,
		&i.TransactionNotBooked,
		&i.EquityOrder,
		&i.AccountValue,
		&i.DepositWithdraw,
		&i.CollateralHaircut,
		&i.EquityPositionValue,
		&i.MarginRequirement,
		&i.Currency,
		&i.BuyLimit,
		&i.SellLimit,
		&i.BuyExposure,
		&i.SellExposure,
		&i.LoanLimit,
		&i.NetLossLimit,
		&i.NetExposureLimit,
		&i.MarginFactor,
		&i.AuthorizedOrderType,
		&i.CalcMode,
		&i.LimitMode,
		&i.ActionMode,
		&i.NettingMode,
		&i.MaxOrdQty,
		&i.MaxOrdVal,
		&i.NoOfTransaction,
		&i.TransactionInterval,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
