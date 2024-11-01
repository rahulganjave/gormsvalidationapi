package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/gormsvalidationapi/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type createInstrumentFilterRequest struct {
	ProductCode         string              `json:"product_code" binding:"required"`
	ExchangeCode        string              `json:"exchange_code" binding:"required"`
	InstrumentCode      string              `json:"instrument_code" binding:"required"`
	UnderlyingCode      *string             `json:"underlying_code"` // Pointer to allow null values
	FilterOption        int                 `json:"filter_option" binding:"default=0"`
	MaxOrderQty         pgtype.Numeric      `json:"max_order_qty" binding:"default=0"`
	MaxOrderPrice       pgtype.Numeric      `json:"max_order_price" binding:"default=0"`
	MaxOrderValue       pgtype.Numeric      `json:"max_order_value" binding:"default=0"`
	BuyLimit            pgtype.Numeric      `json:"buy_lmt" binding:"default=0"`
	SellLimit           pgtype.Numeric      `json:"sell_lmt" binding:"default=0"`
	ExposureFactor      pgtype.Numeric      `json:"exposure_factor" binding:"default=0"`
	BuyExposure         pgtype.Numeric      `json:"buy_exposure" binding:"default=0"`
	SellExposure        pgtype.Numeric      `json:"sell_exposure" binding:"default=0"`
	BuyQtyLimit         pgtype.Numeric      `json:"buy_qty_lmt" binding:"default=0"`
	BuyQtyExposure      pgtype.Numeric      `json:"buy_qty_exposure" binding:"default=0"`
	SellQtyLimit        pgtype.Numeric      `json:"sell_qty_lmt" binding:"default=0"`
	SellQtyExposure     pgtype.Numeric      `json:"sell_qty_exposure" binding:"default=0"`
	NetQtyLimit         pgtype.Numeric      `json:"net_qty_lmt" binding:"default=0"`
	NetQtyExposure      pgtype.Numeric      `json:"net_qty_exposure" binding:"default=0"`
	MarginType          int                 `json:"margin_type" binding:"default=0"`
	InitialMargin       pgtype.Numeric      `json:"initial_margin" binding:"default=0"`
	MaintenanceMargin   pgtype.Numeric      `json:"maintenance_margin" binding:"default=0"`
	MarginableRatio     pgtype.Numeric      `json:"marginable_ratio" binding:"default=0"`
	MarginCurr          *string             `json:"margin_curr"` // Pointer to allow null values
	CommissionType      int                 `json:"commission_type" binding:"default=0"`
	Commission          pgtype.Numeric      `json:"commission" binding:"default=0"`
	MinCommission       pgtype.Numeric      `json:"min_commission" binding:"default=0"`
	CommissionCurr      *string             `json:"commission_curr"` // Pointer to allow null values
	AuthorizedOrderType uint64              `json:"authorized_order_type" binding:"default=0"`
	TradingCurr         *string             `json:"trading_curr"` // Pointer to allow null values
	PriceCode           *string             `json:"price_code"`   // Pointer to allow null values
	RefPx               pgtype.Numeric      `json:"ref_px" binding:"default=0"`
	RefPxType           int16               `json:"ref_px_type" binding:"default=0"` // Pointer to allow null values
	ContractSize        pgtype.Numeric      `json:"contract_size" binding:"default=1"`
	CreatedAt           pgtype.Timestamptz  `json:"created_at" binding:"required"` // Assuming you want to set this on creation
	UpdatedAt           *pgtype.Timestamptz `json:"updated_at"`                    // Pointer to allow null values
}

func (server *Server) createInstrumentFilter(ctx *gin.Context) {
	var req createInstrumentFilterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateInstrumentfilterParams{
		ProductCode:         req.ProductCode,
		ExchangeCode:        req.ExchangeCode,
		InstrumentCode:      req.InstrumentCode,
		UnderlyingCode:      pgtype.Text{String: *req.UnderlyingCode, Valid: req.UnderlyingCode != nil},
		FilterOption:        int32(req.FilterOption),
		MaxOrderQty:         req.MaxOrderQty,
		MaxOrderPrice:       req.MaxOrderPrice,
		MaxOrderValue:       req.MaxOrderValue,
		BuyLimit:            req.BuyLimit,
		SellLimit:           req.SellLimit,
		ExposureFactor:      req.ExposureFactor,
		BuyExposure:         req.BuyExposure,
		SellExposure:        req.SellExposure,
		BuyQtyLimit:         req.BuyQtyLimit,
		BuyQtyExposure:      req.BuyQtyExposure,
		SellQtyLimit:        req.SellQtyLimit,
		SellQtyExposure:     req.SellQtyExposure,
		NetQtyLimit:         req.NetQtyLimit,
		NetQtyExposure:      req.NetQtyExposure,
		MarginType:          int32(req.MarginType),
		InitialMargin:       req.InitialMargin,
		MaintenanceMargin:   req.MaintenanceMargin,
		MarginableRatio:     req.MarginableRatio,
		MarginCurr:          pgtype.Text{String: *req.MarginCurr, Valid: req.MarginCurr != nil},
		CommissionType:      int32(req.CommissionType),
		Commission:          req.Commission,
		MinCommission:       req.MinCommission,
		CommissionCurr:      pgtype.Text{String: *req.CommissionCurr, Valid: req.MarginCurr != nil},
		AuthorizedOrderType: int64(req.AuthorizedOrderType),
		TradingCurr:         pgtype.Text{String: *req.TradingCurr, Valid: req.TradingCurr != nil},
		PriceCode:           pgtype.Text{String: *req.PriceCode, Valid: req.PriceCode != nil},
		RefPx:               req.RefPx,
		RefPxType:           int16(req.RefPxType),
		ContractSize:        req.ContractSize,
		//CreatedAt:           req.CreatedAt,
		//UpdatedAt:           req.UpdatedAt,
	}

	account, err := server.store.CreateInstrumentfilter(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
