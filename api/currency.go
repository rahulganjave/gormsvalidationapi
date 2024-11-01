package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/gormsvalidationapi/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type createCurrencyRequest struct {
	CurrencyPair string         `json:"currency_pair" binding:"required"`
	Bid          pgtype.Numeric `json:"bid" binding:"default=0"`
	Ask          pgtype.Numeric `json:"ask" binding:"default=0"`
}

func (server *Server) createCurrency(ctx *gin.Context) {
	var req createCurrencyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCurrencyParams{
		CurrencyPair: req.CurrencyPair,
		Bid:          req.Bid,
		Ask:          req.Ask,
	}

	currency, err := server.store.CreateCurrency(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, currency)
}
