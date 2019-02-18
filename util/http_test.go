package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUtil(t *testing.T) {

	Convey("GetUrl", t, func() {
		Convey("empty", func() {
			_, err := GetUrl("")
			So(err, ShouldNotBeNil)
		})
		Convey("error status", func() {
			_, err := GetUrl("https://github.com/opsway/documents/foo")
			So(err, ShouldNotBeNil)
		})
		Convey("content", func() {
			actual, err := GetUrl("https://github.com/opsway")
			So(err, ShouldBeNil)
			So(actual, ShouldNotBeNil)
		})
	})

	Convey("IsValidUrl", t, func() {
		Convey("empty", func() {
			So(IsValidUrl(""), ShouldBeFalse)
		})
		Convey("valid", func() {
			So(IsValidUrl("https://opsway.com"), ShouldBeTrue)
		})
	})
}
