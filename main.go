package main

import (
	"github.com/opsway/documents/api"
	"github.com/opsway/documents/util"
	"log"
)

func main() {
	log.Printf("Starting the service...")
	config := api.Config{
		Address:       util.GetEnv("DOCUMENTS_ADDRESS", ":8515"),
		TemplatesPath: util.GetEnv("DOCUMENTS_TEMPLATES", "./templates"),
		PublicPath:    util.GetEnv("DOCUMENTS_PUBLIC_PATH", "./public"),
	}
	srv := api.Server(config)
	log.Printf("The service is ready to listen on %v", config.Address)
	log.Fatal(srv)
}
