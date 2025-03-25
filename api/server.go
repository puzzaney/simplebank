package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/puzzaney/simplebank/db/sqlc"
)

// NOTE: Server serves http request to our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NOTE: NewServer Creates a HTTP server
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
