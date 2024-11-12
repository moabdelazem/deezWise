package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// RegisterRoutes registers the routes for the server
func (s *Server) RegisterRoutes() http.Handler {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	router.Get("/health", s.Health)
	return router
}

// Health returns the health of the server
func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())
	if err != nil {
		log.Fatal("error marshalling health response", err)
	}
	w.Write(jsonResp)
}
