package api

import (
	"encoding/json"
	"net/http"
	"io/ioutil"

	"github.com/opsway/documents/cmd/document"
	"github.com/opsway/documents/cmd/template"
)

// HTMLToPDFGet renders pdf from html
func HTMLToPDFGet(w http.ResponseWriter, r *http.Request) {
	content := r.URL.Query().Get("content")

	if content == "" {
		http.Error(w, "Param 'content' is required", http.StatusBadRequest)
		return
	}

	pdf, err := document.NewPDF()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = pdf.RenderByContent(w, content)

	if err != nil {
		panic(err)
	}
}

// BodyToPDFPost renders pdf from request's HTTP body
func BodyToPDFPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read a request body", http.StatusBadRequest)
		return
	}

	if body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	content := string(body)

	pdf, err := document.NewPDF()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = pdf.RenderByContent(w, content)

	if err != nil {
		panic(err)
	}
}

// RenderTemplateRequest refers request of rendering pdf
type RenderTemplateRequest struct {
	TemplateName    string            `json:"templateName"`
	Data            template.Context  `json:"data"`
	DocumentOptions document.Document `json:"documentOptions"`
}

// RenderTemplate render pdf by template from request
func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	request := RenderTemplateRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pdf, err := document.NewPDF()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pdf.SetOptions(request.DocumentOptions)
	err = pdf.RenderByTemplate(w, request.TemplateName, request.Data)

	if err != nil {
		panic(err)
	}
}
