package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/moabdelazem/deezWise/internals/database"
)

// Server is the main server struct
type Server struct {
	port int

	db database.Service
}

// NewServer creates a new server instance
func NewServer() *http.Server {
	devPort, _ := strconv.Atoi(os.Getenv("PORT"))

	NewServer := &Server{
		port: devPort,
		db:   database.New(),
	}

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
