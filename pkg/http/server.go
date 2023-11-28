package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Server represents the HTTP server.
type Server struct {
	port   int
	router *httprouter.Router
}

// NewServer creates a new instance of Server.
func NewServer(port int) *Server {
	server := &Server{
		port:   port,
		router: httprouter.New(),
	}

	server.setupRoutes()
	return server
}

// Start starts the HTTP server.
func (s *Server) Start() {
	addr := fmt.Sprintf(":%d", s.port)
	http.ListenAndServe(addr, s.router)
}

func (s *Server) setupRoutes() {
	h := NewHandler()
	s.router.POST("/calculate", h.CalculateHandler)
}
