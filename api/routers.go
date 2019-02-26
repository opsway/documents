// Documents generation API
//
//     BasePath: /
//     Version: 1.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/pdf
//
// swagger:meta
package api

import (
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/opsway/documents/api/action"
	"github.com/opsway/documents/api/middleware"
	"github.com/opsway/documents/cmd/template"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewHandler(config Config) http.Handler {
	_ = template.BuildTemplates(config.TemplatesPath) // TODO parse error

	router := NewRouter()
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(config.PublicPath))))

	return handlers.RecoveryHandler()(router)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = middleware.Logger(handler, route.Name)

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
	// - name: content
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
		action.HtmlToPdfGet,
	},

	// swagger:operation POST /render-template RenderTemplatePost
	//
	// Render template to PDF
	//
	// ---
	// produces:
	// - application/pdf
	// parameters:
	// - name: template
	//   description: Name template
	//   required: true
	//   type: string
	// - name: data
	//   description: Name template
	//   required: false
	//   type: object
	// - name: options
	//   description: Options for create PDF
	//   required: false
	//   type: object
	// responses:
	//   '200':
	//     description: A PDF file
	//   '422':
	//     description: Validation error
	Route{
		"RenderTemplatePost",
		strings.ToUpper("Post"),
		"/render-template",
		action.RenderTemplate,
	},
}
