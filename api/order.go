package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/gormsvalidationapi/db/sqlc"
)

type createOrderRequest struct {
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req createOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateOrderParams{}

	order, err := server.store.CreateOrder(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, order)
}
