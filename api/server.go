package api

import (
	"net/http"
	"time"
)

// Server provides new api from config
func Server(config Config) error {
	handler := NewHandler(config)

	srv := &http.Server{
		Handler:      handler,
		Addr:         config.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
