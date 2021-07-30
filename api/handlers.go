package api

import (
        "strings"
	"bytes"
	"encoding/base64"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

// RenderTemplateData refers request of rendering pdf
type RenderTemplateData struct {
	Template        string            `json:"template"`
	Data            template.Context  `json:"data"`
	DocumentOptions document.Document `json:"documentOptions"`
}

// RenderTemplateWithData renders pdf with data
func RenderTemplateWithData(w http.ResponseWriter, r *http.Request) {
	content := r.URL.Query().Get("content")

	if content == "" {
		http.Error(w, "Param 'content' is required", http.StatusBadRequest)
		return
	}
        
        base64Content := strings.Replace(content, " ", "+", -1)
	dec, err := decode([]byte(base64Content))
	if err != nil {
		log.Println(err)
	}

	request := RenderTemplateData{}
        b := bytes.NewBuffer(dec)
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

// RenderTemplatePostData renders pdf with data
func RenderTemplatePostData(w http.ResponseWriter, r *http.Request) {
	//content := string(ioutil.ReadAll(r.Body))
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Param 'content' is required", http.StatusBadRequest)
		return
	}
        
        base64Content := content
	dec, err := decode([]byte(base64Content))
	if err != nil {
		log.Println(err)
	}

	request := RenderTemplateData{}
        b := bytes.NewBuffer(dec)
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func decode(enc []byte) ([]byte, error) {
	// create new buffer from enc
	// you can also use bytes.NewBuffer(enc)
	r := bytes.NewReader(enc)
	// pass it to NewDecoder so that it can read data
	dec := base64.NewDecoder(base64.StdEncoding, r)
	// read decoded data from dec to res
	res, err := ioutil.ReadAll(dec)
	return res, err
}
