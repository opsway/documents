package template

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	buf  bytes.Buffer
	data = TemplateData{
		"foo": "bar",
		"bar": "foo",
	}
)

func TestSpec(t *testing.T) {
	Convey("Template", t, func() {
		Convey("Error loadTemplate", func() {
			_, err := NewTemplate(".", "tests")
			So(err, ShouldBeError)
		})

		Convey("Render", func() {
			var tmpl, err = NewTemplate("../../tests", "foo")

			So(err, ShouldBeNil)

			err = tmpl.Render(TemplateData{
				"foo": "bar",
				"bar": "foo",
			}, &buf)

			So(err, ShouldBeNil)
			So(buf.String(), ShouldEqual, "<div>foobar</div>\n")
		})
	})

	Convey("BuildTemplates", t, func() {
		Convey("Error Build templates", func() {
			err := BuildTemplates("../../tests/quz")
			So(err, ShouldBeError, "failed to read templates dir '../../tests/quz'")
		})

		Convey("Bailed templates", func() {
			err := BuildTemplates("../../tests")
			So(err, ShouldBeNil)
		})
	})

}

func BenchmarkNewTemplateRender(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var tpl, _ = NewTemplate("../../tests", "foo")
		_ = tpl.Render(data, &buf)
	}
}
