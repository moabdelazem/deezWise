package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/moabdelazem/deezWise/internals/server"
)

func gracefulShutdown(server *http.Server) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("error shutting down server: %v", err)
	}

	log.Println("server gracefully stopped")
}

func main() {
	server := server.NewServer()

	go gracefulShutdown(server)

	err := http.ListenAndServe(server.Addr, server.Handler)

	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("error starting server: %v", err))
	}
}
