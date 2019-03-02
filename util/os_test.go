package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestOs(t *testing.T) {
	Convey("GetEnv", t, func() {
		So(GetEnv("NO_SET_VAR", "foo"), ShouldEqual, "foo")
		os.Setenv("SET_VAR", "bar")
		So(GetEnv("SET_VAR", "foo"), ShouldEqual, "bar")
	})

	Convey("IsValidDir", t, func() {
		So(IsValidDir("../testdata"), ShouldBeTrue)
		So(IsValidDir("tests"), ShouldBeFalse)
	})
}
