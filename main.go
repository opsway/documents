package main

import (
	"github.com/opsway/documents/api"
	"github.com/opsway/documents/util"
)

func main() {
	api.Server(api.Config{
		Address:       util.GetEnv("DOCUMENTS_ADDRESS", ":8515"),
		TemplatesPath: util.GetEnv("DOCUMENTS_TEMPLATES", "./templates"),
		PublicPath:    util.GetEnv("DOCUMENTS_PUBLIC_PATH", "./public"),
	})
}
