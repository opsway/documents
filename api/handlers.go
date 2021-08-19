package api

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"github.com/opsway/documents/cmd/document"
	"github.com/opsway/documents/cmd/template"
	"io/ioutil"
	"net/http"
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

// PrepareTemplateData refers request of rendering pdf
type PrepareTemplateData struct {
	Template        string            `json:"template"`
	Data            template.Context  `json:"data"`
	DocumentOptions document.Document `json:"documentOptions"`
}

// RenderPdfWithPostTemplateData renders pdf with posted template & data
func RenderPdfWithPostTemplateData(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Param 'content' is required", http.StatusBadRequest)
		return
	}

	request := PrepareTemplateData{}
	b := bytes.NewBuffer(content)
	json.NewDecoder(b).Decode(&request)

	enableCors(&w)
	pdf, err := document.NewPDF()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pdf.SetOptions(request.DocumentOptions)
	err = pdf.RenderByVirtualTemplate(w, request.Template, request.Data)

	if err != nil {
		panic(err)
	}
}

// RenderPdfWithPostTemplateOptions return options response
func RenderPdfWithPostTemplateOptions(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}
