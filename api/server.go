package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	db "github.com/gormsvalidationapi/db/sqlc"
	"github.com/gormsvalidationapi/util"
)

type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	//router.POST("/orders", server.createOrder)
	//router.GET("/positions", server.getPositions)
	//router.POST("/instrumentfilters", server.createInstrumentfilter)

	server.router = router

}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.route.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
