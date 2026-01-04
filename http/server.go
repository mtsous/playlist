package http

import (
	"log"
	"net/http"
	"time"
)

const (
	port = "8080"
)

type Server struct {
	server *http.Server
	router *http.ServeMux
}

func NewServer() *Server {
	s := &Server{
		router: http.NewServeMux(),
	}

	// Register routes
	s.registerPlaylistRoutes(s.router)

	// Create server with router as handler
	s.server = &http.Server{
		Addr:         ":" + port,
		Handler:      s.router, // ← Assign router here
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s
}

func (s *Server) Open() {
	log.Print("server running on port " + port)

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatal("fatal err")
	}
}
