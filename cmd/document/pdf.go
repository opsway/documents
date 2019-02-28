package document

import (
	"bytes"
	"io"
	"strings"

	generator "github.com/SebastiaanKlippert/go-wkhtmltopdf"

	"github.com/opsway/documents/cmd/template"
	"github.com/opsway/documents/util"
)

//Pdf is structure of pdf generator
type Pdf struct {
	generator *generator.PDFGenerator
	option    Document
}

// SetOptions settle margin and page size of pdf
func (pdf *Pdf) SetOptions(option Document) {
	pdf.generator.Orientation.Set(option.Orientation)
	pdf.generator.PageSize.Set(option.PageSize)
	pdf.generator.MarginBottom.Set(option.MarginBottom)
	pdf.generator.MarginTop.Set(option.MarginTop)
	pdf.generator.MarginLeft.Set(option.MarginLeft)
	pdf.generator.MarginRight.Set(option.MarginRight)
}

// AddPageFromURL generates pdf pages from url
func (pdf *Pdf) AddPageFromURL(url string) {
	pdf.generator.AddPage(generator.NewPage(url))
}

// AddPageFromString generates pdf pages including string
func (pdf *Pdf) AddPageFromString(content string) {
	pdf.AddPage(strings.NewReader(content))
}

// AddPage generates pdf pages inputed content
func (pdf *Pdf) AddPage(input io.Reader) {
	pdf.generator.AddPage(generator.NewPageReader(input))
}

// Render creates pdf to writer
func (pdf *Pdf) Render(writer io.Writer) error {
	pdf.generator.SetOutput(writer)

	return pdf.generator.Create()
}

// RenderByContent creates pdf from content
func (pdf *Pdf) RenderByContent(writer io.Writer, content string) error {
	if util.IsValidURL(content) {
		pdf.AddPageFromURL(content)
	} else {
		pdf.AddPageFromString(content)
	}

	return pdf.Render(writer)
}

// RenderByTemplate creates pdf from template and data
func (pdf *Pdf) RenderByTemplate(writer io.Writer, templateName string, data template.Context) error {
	tmpl, err := template.GetTemplate(templateName)

	if err != nil {
		return err
	}

	var input io.Reader
	var buf bytes.Buffer

	err = tmpl.Render(data, &buf)

	if err != nil {
		return err
	}

	input = &buf

	pdf.AddPage(input)

	return pdf.Render(writer)
}

// NewPdf return pdf generator
func NewPdf() (*Pdf, error) {
	pdfg, err := generator.NewPDFGenerator()

	if err != nil {
		return nil, err
	}

	return &Pdf{
		generator: pdfg,
	}, nil
}
