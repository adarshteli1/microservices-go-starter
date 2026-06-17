package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/env"
	"syscall"
	"time"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8083")
)

func main() {

	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)

	mux := http.NewServeMux()
	httphandler := h.HttpHandler{Service: svc}

	mux.HandleFunc("POST /preview", httphandler.HandleTripPreview)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	// GRACEFUL SHUTDOWN

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Listening to the Server %s", httpAddr)
		serverErrors <- server.ListenAndServe()

	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Printf("Error while loading the server %s", err)

	case sig := <-shutdown:
		log.Printf("Server is shutting down due to %v signal:", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Server Could not Gracefully Shutdown %s", err)
			server.Close()
		}
	}
}
