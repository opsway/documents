package cmd

import (
	. "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func RenderPdf(content string) ([]byte, error) {
	// Create new PDF generator
	pdfg, err := NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	// TODO move to options
	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(OrientationLandscape)

	// Create a new input page from an URL
	page := NewPage(content)
	pdfg.AddPage(page)
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}
