package main

import (
	"context"
	"go-microservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	logger := log.New(os.Stdout, "go-mcrsvc", log.LstdFlags)

	// prepare handler
	productHandler := handlers.NewProduct(logger)

	// create servemux
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productHandler.GetProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productHandler.AddProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.Use(productHandler.MiddlewareProductValidation)
	putRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProduct)

	// sm.Handle("/products", productHandler)

	// create server
	s := &http.Server{
		Addr:        ":9090",
		Handler:     sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
	}

	// start server
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	sig := <-sigChan
	logger.Println("Got signal:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.Shutdown(ctx); err != nil {
		logger.Fatal(err)
	}

	logger.Println("Server gracefully stopped")
}
