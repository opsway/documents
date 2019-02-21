package template

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	buf  bytes.Buffer
	data = Context{
		"foo": "bar",
		"bar": "foo",
	}
)

func TestSpec(t *testing.T) {
	Convey("Template", t, func() {
		Convey("error: load", func() {
			actual, err := NewTemplate(".", "quz")
			So(actual, ShouldBeNil)
			So(err, ShouldBeError, "stat quz/index.html: no such file or directory")
		})

		Convey("successful render", func() {
			actual, err := NewTemplate("../../testdata/goodTemplates", "foo")
			So(err, ShouldBeNil)
			err = actual.Render(Context{
				"foo": "bar",
				"bar": "foo",
			}, &buf)
			So(err, ShouldBeNil)
			So(buf.String(), ShouldEqual, "<div>foobar</div>\n")
		})
	})

	Convey("BuildTemplates", t, func() {
		Convey("error: path of templates is not exist", func() {
			err := BuildTemplates("../../testdata/noExistTemplates")
			So(err, ShouldBeError, "failed to read templates dir '../../testdata/noExistTemplates'")
		})

		Convey("error: template is bad", func() {
			err := BuildTemplates("../../testdata/badTemplates")
			So(err, ShouldBeError, "[Error (where: parser) in ../../testdata/badTemplates/foo/index.html | Line 1 Col 8 near 'elseif'] Tag 'elseif' not found (or beginning tag not provided)")
		})

		Convey("successful", func() {
			err := BuildTemplates("../../testdata/goodTemplates")
			So(err, ShouldBeNil)
		})
	})

	Convey("GetTemplate", t, func() {
		Convey("error: is not exist", func() {
			actual, err := GetTemplate("quz")
			So(actual, ShouldBeNil)
			So(err, ShouldBeError, "template 'quz' is not exist")
		})

		Convey("successful", func() {
			err := BuildTemplates("../../testdata/goodTemplates")
			So(err, ShouldBeNil)
			actual, err := GetTemplate("foo")
			So(actual, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
	})
}

func BenchmarkNewTemplateRender(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var tpl, _ = NewTemplate("../../testdata/goodTemplates", "foo")
		_ = tpl.Render(data, &buf)
	}
}
