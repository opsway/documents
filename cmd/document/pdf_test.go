package document

import (
	"bytes"
	"os"
	"testing"

	"github.com/opsway/documents/cmd/template"
	. "github.com/smartystreets/goconvey/convey"
)

type BadWriter struct {
	err error
}

func (w BadWriter) Write(p []byte) (n int, err error) {
	return 0, w.err
}

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
			err = pdf.RenderByContent(&buf, "file://../../testdata/goodTemplates/foo/index.html")
			So(err, ShouldBeNil)
			So(buf, ShouldNotBeNil)
		})
	})

	Convey("RenderByTemplate", t, func() {
		_ = template.BuildTemplates("../../testdata/goodTemplates")
		pdf, _ := NewPdf()
		pdf.SetOptions(Document{})

		Convey("error: template", func() {
			//var actual bytes.Buffer
			actual := bytes.NewBuffer(nil)
			err := pdf.RenderByTemplate(actual, "quz", template.Context{})
			So(err, ShouldBeError, "template 'quz' is not exist")
			So(actual.String(), ShouldBeEmpty)
		})

		Convey("error: PDF render", func() {
			actual := new(BadWriter)
			err := pdf.RenderByTemplate(actual, "foo", template.Context{})
			So(err, ShouldBeError)
		})

		Convey("successful", func() {
			var actual bytes.Buffer
			err := pdf.RenderByTemplate(&actual, "foo", template.Context{})
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
		_ = pdf.RenderByContent(&buf, "<p>hello</p>")
	}
}
