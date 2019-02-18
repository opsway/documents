package main

import (
	"github.com/opsway/documents/api"
	"github.com/opsway/documents/util"
)

func main() {
	api.Server(api.Config{
		Address:       util.Getenv("DOCUMENTS_ADDRESS", ":8515"),
		TemplatesPath: util.Getenv("DOCUMENTS_TEMPLATES", "./templates"),
		PublicPath:    util.Getenv("DOCUMENTS_PUBLIC_PATH", "./public"),
	})
}
