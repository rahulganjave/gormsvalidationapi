package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/gormsvalidationapi/db/sqlc"
	"github.com/jackc/pgx/pgtype"
)

type createAccountRequest struct {
	ClientID             pgtype.Text    `json:"client_id" binding:"required"`
	Account              string         `json:"account" binding:"required"`
	Trader               pgtype.Text    `json:"trader" binding:"required"`
	CashBalanceBf        pgtype.Numeric `json:"cash_balance_bf" binding:"required"`
	CashBalance          pgtype.Numeric `json:"cash_balance" binding:"required"`
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
	Currency             pgtype.Text    `json:"currency" binding:"required,oneof=USD EUR CAD"`
	BuyLimit             pgtype.Numeric `json:"buy_limit" binding:"required"`
	SellLimit            pgtype.Numeric `json:"sell_limit" binding:"required"`
	BuyExposure          pgtype.Numeric `json:"buy_exposure"`
	SellExposure         pgtype.Numeric `json:"sell_exposure"`
	LoanLimit            pgtype.Numeric `json:"loan_limit"`
	MarginFactor         pgtype.Numeric `json:"margin_factor"`
	AuthorizedOrderType  pgtype.Int8    `json:"authorized_order_type"`
	CalcMode             pgtype.Text    `json:"calc_mode"`
	LimitMode            pgtype.Int4    `json:"limit_mode"`
	Action               pgtype.Text    `json:"action"`
	NettingMode          string         `json:"netting_mode"`
	MaxOrdQty            pgtype.Numeric `json:"max_ord_qty"`
	MaxOrdVal            pgtype.Numeric `json:"max_ord_val"`
	NoOfTransaction      int32          `json:"no_of_transaction"`
	TransactionInterval  int16          `json:"transaction_interval"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		ClientID:             req.ClientID,
		Account:              req.Account,
		Trader:               req.Trader,
		CashBalanceBf:        req.CashBalanceBf,
		CashBalance:          req.CashBalance,
		EstTransCost:         req.EstTransCost,
		EquityRealizedPl:     req.EquityRealizedPl,
		EquityUnrealizedPl:   req.EquityUnrealizedPl,
		MarginRealizedPl:     req.MarginRealizedPl,
		MarginUnrealizedPl:   req.MarginUnrealizedPl,
		EstClosingCost:       req.EstClosingCost,
		TransactionNotBooked: req.TransactionNotBooked,
		EquityOrder:          req.EquityOrder,
		AccountValue:         req.AccountValue,
		DepositWithdraw:      req.DepositWithdraw,
		CollateralHaircut:    req.CollateralHaircut,
		EquityPositionValue:  req.EquityPositionValue,
		MarginRequirement:    req.MarginRequirement,
		Currency:             req.Currency,
		BuyLimit:             req.BuyLimit,
		SellLimit:            req.SellLimit,
		BuyExposure:          req.BuyExposure,
		SellExposure:         req.SellExposure,
		LoanLimit:            req.LoanLimit,
		MarginFactor:         req.MarginFactor,
		AuthorizedOrderType:  req.AuthorizedOrderType,
		CalcMode:             req.CalcMode,
		LimitMode:            req.LimitMode,
		Action:               req.Action,
		NettingMode:          req.NettingMode,
		MaxOrdQty:            req.MaxOrdQty,
		MaxOrdVal:            req.MaxOrdVal,
		NoOfTransaction:      req.NoOfTransaction,
		TransactionInterval:  req.TransactionInterval,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
