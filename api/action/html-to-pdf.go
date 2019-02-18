package action

import (
	"net/http"

	"github.com/opsway/documents/cmd/document"
)

func HtmlToPdfGet(w http.ResponseWriter, r *http.Request) {
	content := r.URL.Query().Get("content")

	if content == "" {
		http.Error(w, "Param 'content' is required", http.StatusBadRequest)
		return
	}

	pdf, err := document.NewPdf()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = pdf.RenderByContent(w, content)

	if err != nil {
		panic(err)
	}
}
