package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/opsway/documents/api"
)

func main() {
	var (
		addr       string
		publicPath string
	)

	flag.StringVar(&publicPath, "public", "./public", "-public ./public")
	flag.StringVar(&addr, "addr", "0.0.0.0:8515", "-addr 0.0.0.0:8515")

	flag.Parse()

	log.Printf("Starting the service...")

	router := api.NewRouter()
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(publicPath))))

	srv := &http.Server{
		Handler: router,
		Addr:    addr,
	}

	log.Printf("The service is ready to listen on %v public path is %v", addr, publicPath)
	log.Fatal(srv.ListenAndServe())
}
