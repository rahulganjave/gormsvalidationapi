package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/gormsvalidationapi/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type createInstrumentRequest struct {
	ProductCode       string              `json:"product_code" binding:"required"`
	ExchangeCode      string              `json:"exchange_code" binding:"required"`
	InstrumentCode    string              `json:"instrument_code" binding:"required"`
	UnderlyingCode    *string             `json:"underlying_code"`    // Pointer to allow null values
	InstrumentDesc    *string             `json:"instrument_desc"`    // Pointer to allow null values
	MaturityMonthYear *string             `json:"maturity_monthyear"` // Pointer to allow null values
	MaturityDay       *string             `json:"maturity_day"`       // Pointer to allow null values
	PutOrCall         *string             `json:"put_or_call"`        // Pointer to allow null values
	StrikePrice       pgtype.Numeric      `json:"strike_price" binding:"default=0"`
	Symbol            *string             `json:"symbol"`                    // Pointer to allow null values
	SymbolSfx         *string             `json:"symbol_sfx"`                // Pointer to allow null values
	Cusip             *string             `json:"cusip"`                     // Pointer to allow null values
	Isin              *string             `json:"isin"`                      // Pointer to allow null values
	Ric               *string             `json:"ric"`                       // Pointer to allow null values
	Sedol             *string             `json:"sedol"`                     // Pointer to allow null values
	Blpsyn            *string             `json:"blpsyn"`                    // Pointer to allow null values
	CfiCode           *string             `json:"cficode"`                   // Pointer to allow null values
	PriceCode         *string             `json:"price_code"`                // Pointer to allow null values
	EffectiveDate     *pgtype.Timestamptz `json:"effective_date"`            // Pointer to allow null values
	MinQty            int32               `json:"minqty" binding:"required"` // Changed to int32
	Sector            *string             `json:"sector"`                    // Pointer to allow null values
	ContractSize      pgtype.Numeric      `json:"contract_size" binding:"default=1"`
	CorporateAction   *string             `json:"corporate_action"`   // Pointer to allow null values
	TradingSessionID  *string             `json:"trading_session_id"` // Pointer to allow null values
	TradingCurr       *string             `json:"trading_curr"`       // Pointer to allow null values
	PriceDec          pgtype.Numeric      `json:"price_dec" binding:"default=0"`
	HandlInst         *string             `json:"handl_inst"`                 // Pointer to allow null values
	TickID            int32               `json:"tick_id" binding:"required"` // Changed to int32
	MarketSlippage    pgtype.Numeric      `json:"market_slippage" binding:"default=0"`
	RefPx             pgtype.Numeric      `json:"ref_px" binding:"default=0"`
	RefPxType         int16               `json:"ref_px_type" binding:"default=0"` // Changed to int32
	TPlusOne          int32               `json:"tplusone"`                        // Changed to int32
	//CreatedAt         pgtype.Timestamptz  `json:"created_at" binding:"required"`   // Assuming you want to set this on creation
	//UpdatedAt         *pgtype.Timestamptz `json:"updated_at"`
}

func (server *Server) createInstrument(ctx *gin.Context) {
	var req createInstrumentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateInstrumentParams{
		ProductCode:       req.ProductCode,
		ExchangeCode:      req.ExchangeCode,
		InstrumentCode:    req.InstrumentCode,
		UnderlyingCode:    pgtype.Text{String: *req.UnderlyingCode, Valid: req.UnderlyingCode != nil},
		InstrumentDesc:    pgtype.Text{String: *req.InstrumentDesc, Valid: req.InstrumentDesc != nil},
		MaturityMonthyear: pgtype.Text{String: *req.MaturityMonthYear, Valid: req.MaturityMonthYear != nil},
		MaturityDay:       pgtype.Text{String: *req.MaturityDay, Valid: req.MaturityDay != nil},
		PutOrCall:         pgtype.Text{String: *req.PutOrCall, Valid: req.PutOrCall != nil},
		StrikePrice:       req.StrikePrice,
		Symbol:            pgtype.Text{String: *req.Symbol, Valid: req.Symbol != nil},
		SymbolSfx:         pgtype.Text{String: *req.SymbolSfx, Valid: req.SymbolSfx != nil},
		Cusip:             pgtype.Text{String: *req.Cusip, Valid: req.Cusip != nil},
		Isin:              pgtype.Text{String: *req.Isin, Valid: req.Isin != nil},
		Ric:               pgtype.Text{String: *req.Ric, Valid: req.Ric != nil},
		Sedol:             pgtype.Text{String: *req.Sedol, Valid: req.Sedol != nil},
		Blpsyn:            pgtype.Text{String: *req.Blpsyn, Valid: req.Blpsyn != nil},
		Cficode:           pgtype.Text{String: *req.CfiCode, Valid: req.CfiCode != nil},
		PriceCode:         pgtype.Text{String: *req.PriceCode, Valid: req.PriceCode != nil},
		EffectiveDate:     pgtype.Timestamp{Time: req.EffectiveDate.Time, Valid: req.EffectiveDate != nil},
		Minqty:            int32(req.MinQty),
		Sector:            pgtype.Text{String: *req.Sector, Valid: req.Sector != nil},
		ContractSize:      req.ContractSize,
		CorporateAction:   pgtype.Text{String: *req.CorporateAction, Valid: req.CorporateAction != nil},
		TradingSessionID:  pgtype.Text{String: *req.TradingSessionID, Valid: req.TradingSessionID != nil},
		TradingCurr:       pgtype.Text{String: *req.TradingCurr, Valid: req.TradingCurr != nil},
		PriceDec:          req.PriceDec,
		HandlInst:         pgtype.Text{String: *req.HandlInst, Valid: req.HandlInst != nil},
		TickID:            int32(req.TickID),
		MarketSlippage:    req.MarketSlippage,
		RefPx:             req.RefPx,
		RefPxType:         int16(req.RefPxType),
		Tplusone:          pgtype.Int4{Int32: req.TPlusOne, Valid: true},
		//CreatedAt:         req.CreatedAt,
		//UpdatedAt:         req.UpdatedAt,
	}

	account, err := server.store.CreateInstrument(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
