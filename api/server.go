package api

import (
	"log"
	"net/http"
	"time"
)

func Server(config Config) {
	log.Printf("Starting the service...")

	handler := NewHandler(config)

	srv := &http.Server{
		Handler:      handler,
		Addr:         config.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("The service is ready to listen on %v", config.Address)
	log.Fatal(srv.ListenAndServe())
}
