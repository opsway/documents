package document

import (
	"bytes"
	"os"
	"testing"

	"github.com/opsway/documents/cmd/template"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPdf(t *testing.T) {

	Convey("NewPdf", t, func() {
		Convey("empty", func() {
			path := os.Getenv("PATH")
			os.Setenv("PATH", "")
			_, err := NewPdf()
			os.Setenv("PATH", path)
			So(err, ShouldBeError, "wkhtmltopdf not found")
		})
	})

	Convey("RenderByContent", t, func() {
		pdf, err := NewPdf()
		So(err, ShouldBeNil)

		Convey("empty", func() {
			var buf bytes.Buffer
			err = pdf.RenderByContent(&buf, "")
			So(err, ShouldBeNil)
			So(buf, ShouldNotBeNil)
		})

		Convey("content by url", func() {
			var buf bytes.Buffer
			err = pdf.RenderByContent(&buf, "https://github.com/opsway")
			So(err, ShouldBeNil)
			So(buf, ShouldNotBeNil)
		})
	})

	Convey("RenderByTemplate", t, func() {
		_ = template.BuildTemplates("../../tests")
		pdf, _ := NewPdf()
		pdf.SetOptions(Document{})

		Convey("error template", func() {
			var actual bytes.Buffer

			err := pdf.RenderByTemplate(&actual, "quz", template.TemplateData{})
			So(err, ShouldBeError, "template 'quz' is not exist")
			So(actual.String(), ShouldBeEmpty)
		})

		Convey("template", func() {
			var actual bytes.Buffer

			err := pdf.RenderByTemplate(&actual, "foo", template.TemplateData{})
			So(err, ShouldBeNil)
			So(actual, ShouldNotBeNil)
		})
	})
}

func BenchmarkPdfRenderByContent(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pdf, _ := NewPdf()
		var buf bytes.Buffer
		_ = pdf.RenderByContent(&buf, "https://github.com/opsway")
	}
}
