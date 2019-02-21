package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestOs(t *testing.T) {
	Convey("Getenv", t, func() {
		So(Getenv("NO_SET_VAR", "foo"), ShouldEqual, "foo")
		os.Setenv("SET_VAR", "bar")
		So(Getenv("SET_VAR", "foo"), ShouldEqual, "bar")
	})

	Convey("IsValidDir", t, func() {
		So(IsValidDir("../testdata"), ShouldBeTrue)
		So(IsValidDir("tests"), ShouldBeFalse)
	})
}
