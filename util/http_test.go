package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtil(t *testing.T) {

	Convey("GetURL", t, func() {
		Convey("empty", func() {
			_, err := GetURL("")
			So(err, ShouldNotBeNil)
		})
		Convey("error status", func() {
			_, err := GetURL("https://github.com/opsway/documents/foo")
			So(err, ShouldNotBeNil)
		})
		Convey("content", func() {
			actual, err := GetURL("https://github.com/opsway")
			So(err, ShouldBeNil)
			So(actual, ShouldNotBeNil)
		})
	})

	Convey("IsValidURL", t, func() {
		Convey("empty", func() {
			So(IsValidURL(""), ShouldBeFalse)
		})
		Convey("valid", func() {
			So(IsValidURL("https://opsway.com"), ShouldBeTrue)
		})
	})
}
