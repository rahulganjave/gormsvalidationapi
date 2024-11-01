package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/gormsvalidationapi/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type createPositionRequest struct {
	PositionID             string              `json:"position_id" binding:"required"`
	RefPositionID          *string             `json:"ref_position_id"` // Pointer to allow null values
	TradingAcct            string              `json:"trading_acct" binding:"required"`
	ClientID               string              `json:"client_id" binding:"required"`
	ProductCode            string              `json:"product_code" binding:"required"`
	ExchangeCode           string              `json:"exchange_code" binding:"required"`
	InstrumentCode         string              `json:"instrument_code" binding:"required"`
	PositionSide           int                 `json:"position_side" binding:"default=1"`
	PositionStatus         int                 `json:"position_status" binding:"default=1"`
	PositionType           int                 `json:"position_type" binding:"default=0"`
	SettlementStatus       int                 `json:"settlement_status" binding:"default=0"`
	CalcType               int                 `json:"calc_type" binding:"default=0"`
	OpenQty                pgtype.Numeric      `json:"open_qty" binding:"default=0"`
	OpenPx                 pgtype.Numeric      `json:"open_px" binding:"default=0"`
	ClosedQty              pgtype.Numeric      `json:"closed_qty" binding:"default=0"`
	ClosedPx               pgtype.Numeric      `json:"closed_px" binding:"default=0"`
	TradingCurr            *string             `json:"trading_curr"` // Pointer to allow null values
	SettleCurr             *string             `json:"settle_curr"`  // Pointer to allow null values
	MarkPrice              pgtype.Numeric      `json:"mark_price" binding:"default=0"`
	MarkValue              pgtype.Numeric      `json:"mark_value" binding:"default=0"`
	ExchangeRate           pgtype.Numeric      `json:"exchange_rate" binding:"default=0"`
	UnrealizedPL           pgtype.Numeric      `json:"unrealized_pl" binding:"default=0"`
	RealizedPL             pgtype.Numeric      `json:"realized_pl" binding:"default=0"`
	MmAndCash              pgtype.Numeric      `json:"mm_and_cash" binding:"default=0"`
	CollateralHaircutRatio pgtype.Numeric      `json:"collateral_haircut_ratio" binding:"default=0"`
	CollateralHaircut      pgtype.Numeric      `json:"collateral_haircut" binding:"default=0"`
	OpenDate               *pgtype.Timestamptz `json:"open_date"`   // Pointer to allow null values
	ClosedDate             *pgtype.Timestamptz `json:"closed_date"` // Pointer to allow null values
	CostToClose            pgtype.Numeric      `json:"cost_to_close" binding:"default=0"`
	TradeDate              *time.Time          `json:"trade_date"`         // Pointer to allow null values
	TradingSessionID       *string             `json:"trading_session_id"` // Pointer to allow null values
	UnderlyingCode         *string             `json:"underlying_code"`    // Pointer to allow null values
	OpenClose              int                 `json:"open_close" binding:"default=79"`
	PutOrCall              *int                `json:"put_or_call"` // Pointer to allow null values
	ContractSize           pgtype.Numeric      `json:"contract_size" binding:"default=1"`
}

func (server *Server) createPosition(ctx *gin.Context) {
	var req createPositionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePositionParams{
		PositionID:             req.PositionID,
		RefPositionID:          pgtype.Text{String: *req.RefPositionID, Valid: req.RefPositionID != nil},
		TradingAcct:            req.TradingAcct,
		ClientID:               req.ClientID,
		ProductCode:            req.ProductCode,
		ExchangeCode:           req.ExchangeCode,
		InstrumentCode:         req.InstrumentCode,
		PositionSide:           int32(req.PositionSide),
		PositionStatus:         int32(req.PositionStatus),
		PositionType:           int32(req.PositionType),
		SettlementStatus:       int32(req.SettlementStatus),
		CalcType:               int32(req.CalcType),
		OpenQty:                req.OpenQty,
		OpenPx:                 req.OpenPx,
		ClosedQty:              req.ClosedQty,
		ClosedPx:               req.ClosedPx,
		TradingCurr:            (*req.TradingCurr),
		SettleCurr:             (*req.SettleCurr),
		MarkPrice:              req.MarkPrice,
		MarkValue:              req.MarkValue,
		ExchangeRate:           req.ExchangeRate,
		UnrealizedPl:           req.UnrealizedPL,
		RealizedPl:             req.RealizedPL,
		MmAndCash:              req.MmAndCash,
		CollateralHaircut:      req.CollateralHaircut,
		CollateralHaircutRatio: req.CollateralHaircutRatio,
		OpenDate:               pgtype.Timestamptz{Time: req.OpenDate.Time, Valid: req.OpenDate.Valid},
		ClosedDate:             pgtype.Timestamptz{Time: req.ClosedDate.Time, Valid: req.OpenDate.Valid},
		CostToClose:            req.CostToClose,
		TradeDate:              pgtype.Date{Time: *req.TradeDate, Valid: req.TradeDate != nil},
		TradingSessionID:       pgtype.Text{String: *req.TradingSessionID, Valid: req.TradingSessionID != nil},
		UnderlyingCode:         pgtype.Text{String: *req.UnderlyingCode, Valid: req.UnderlyingCode != nil},
		OpenClose:              int32(req.OpenClose),
		PutOrCall:              pgtype.Int4{Int32: int32(*req.PutOrCall), Valid: req.PutOrCall != nil},
		ContractSize:           req.ContractSize,
	}

	position, err := server.store.CreatePosition(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, position)
}
