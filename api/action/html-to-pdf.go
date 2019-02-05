package action

import (
	"github.com/opsway/documents/cmd"
	"net/http"
)

func HtmlToPdfGet(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	if url == "" {
		http.Error(w, "Param url is required", http.StatusUnprocessableEntity)
		return
	}

	pdf, err := cmd.RenderPdf(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/pdf")
	_, _ = w.Write(pdf)
}
