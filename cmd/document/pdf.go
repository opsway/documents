package document

import (
	"bytes"
	"io"
	"strings"

	generator "github.com/SebastiaanKlippert/go-wkhtmltopdf"

	"github.com/opsway/documents/cmd/template"
	"github.com/opsway/documents/util"
)

// PDF is structure of PDF generator
type PDF struct {
	generator *generator.PDFGenerator
	option    Document
}

// SetOptions settle margin and page size of PDF
func (pdf *PDF) SetOptions(option Document) {
	pdf.generator.Orientation.Set(option.Orientation)
	pdf.generator.PageSize.Set(option.PageSize)
	pdf.generator.MarginBottom.Set(option.MarginBottom)
	pdf.generator.MarginTop.Set(option.MarginTop)
	pdf.generator.MarginLeft.Set(option.MarginLeft)
	pdf.generator.MarginRight.Set(option.MarginRight)
}

// AddPageFromURL generates PDF pages from URL
func (pdf *PDF) AddPageFromURL(url string) {
	pdf.generator.AddPage(generator.NewPage(url))
}

// AddPageFromString generates PDF pages including string
func (pdf *PDF) AddPageFromString(content string) {
	pdf.AddPage(strings.NewReader(content))
}

// AddPage generates PDF pages from reader
func (pdf *PDF) AddPage(input io.Reader) {
	pdf.generator.AddPage(generator.NewPageReader(input))
}

// Render creates PDF to writer
func (pdf *PDF) Render(writer io.Writer) error {
	pdf.generator.SetOutput(writer)

	return pdf.generator.Create()
}

// RenderByContent creates PDF from content to writer
func (pdf *PDF) RenderByContent(writer io.Writer, content string) error {
	if util.IsValidURL(content) {
		pdf.AddPageFromURL(content)
	} else {
		pdf.AddPageFromString(content)
	}

	return pdf.Render(writer)
}

// RenderByTemplate creates PDF from template and data to writer
func (pdf *PDF) RenderByTemplate(writer io.Writer, templateName string, data template.Context) error {
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

// RenderByVirtualTemplate creates PDF from specified template content and data to writer
func (pdf *PDF) RenderByVirtualTemplate(writer io.Writer, templateContent string, data template.Context) error {
	tmpl, err := template.NewVirtualTemplate(templateContent)

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

// NewPDF return PDF generator
func NewPDF() (*PDF, error) {
	PDFGen, err := generator.NewPDFGenerator()

	if err != nil {
		return nil, err
	}

	return &PDF{
		generator: PDFGen,
	}, nil
}
