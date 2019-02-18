package document

import (
	"bytes"
	"io"
	"strings"

	generator "github.com/SebastiaanKlippert/go-wkhtmltopdf"

	"github.com/opsway/documents/cmd/template"
	"github.com/opsway/documents/util"
)

type Pdf struct {
	generator *generator.PDFGenerator
	option    Document
}

func (pdf *Pdf) SetOptions(option Document) {
	pdf.generator.Orientation.Set(option.Orientation)
	pdf.generator.PageSize.Set(option.PageSize)
	pdf.generator.MarginBottom.Set(option.MarginBottom)
	pdf.generator.MarginTop.Set(option.MarginTop)
	pdf.generator.MarginLeft.Set(option.MarginLeft)
	pdf.generator.MarginRight.Set(option.MarginRight)
}

func (pdf *Pdf) AddPageFromUrl(url string) {
	pdf.generator.AddPage(generator.NewPage(url))
}

func (pdf *Pdf) AddPageFromString(content string) {
	pdf.AddPage(strings.NewReader(content))
}

func (pdf *Pdf) AddPage(input io.Reader) {
	pdf.generator.AddPage(generator.NewPageReader(input))
}

func (pdf *Pdf) Render(writer io.Writer) error {
	pdf.generator.SetOutput(writer)

	return pdf.generator.Create()
}

func (pdf *Pdf) RenderByContent(writer io.Writer, content string) error {
	if util.IsValidUrl(content) {
		pdf.AddPageFromUrl(content)
	} else {
		pdf.AddPageFromString(content)
	}

	return pdf.Render(writer)
}

func (pdf *Pdf) RenderByTemplate(writer io.Writer, templateName string, data template.TemplateData) error {
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

func NewPdf() (*Pdf, error) {
	pdfg, err := generator.NewPDFGenerator()

	if err != nil {
		return nil, err
	}

	return &Pdf{
		generator: pdfg,
	}, nil
}
