package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const TIMEOUT = 30 * time.Second

type ServerOption func(server *http.Server)

func Start(port string, handler http.Handler, options ...ServerOption) *http.Server {
	server := &http.Server{
		ReadTimeout:  TIMEOUT,
		WriteTimeout: TIMEOUT,
		Addr:         ":" + port,
		Handler:      handler,
	}

	for _, option := range options {
		option(server)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		<-ctx.Done()
		log.Println("Shutting down server...")
		if err := server.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	log.Printf("Starting server on port %s...\n", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %v", err)
	}

	return nil
}
