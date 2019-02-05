// Package classification Documents generation API
//
//     BasePath: /
//     Version: 1.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//     - application/pdf
//
// swagger:meta
package api

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	. "github.com/opsway/documents/api/action"
)

type Config struct {
	PublicPath string
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{

	// swagger:operation GET /html-to-pdf HtmlToPdfGet
	//
	// Render URL to PDF
	//
	// ---
	// produces:
	// - application/pdf
	// parameters:
	// - name: url
	//   in: query
	//   description: Content URL for render
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: A PDF file
	//   '422':
	//     description: Validation error
	Route{
		"HtmlToPdfGet",
		strings.ToUpper("Get"),
		"/html-to-pdf",
		HtmlToPdfGet,
	},
}
