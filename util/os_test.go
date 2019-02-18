package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOs(t *testing.T) {
	Convey("IsValidDir", t, func() {
		So(IsValidDir("../tests"), ShouldBeTrue)
		So(IsValidDir("tests"), ShouldBeFalse)
	})
}
