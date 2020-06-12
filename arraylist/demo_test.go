package arraylist

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestArrayList(t *testing.T) {
	arr := NewArrayList()
	Convey("Test insert", t, func() {
		err := arr.Insert("-1", -1)
		So(err, ShouldNotBeNil)
		err = arr.Insert("4", 4)
		So(err, ShouldBeNil)
		Convey("Test append", func() {
			arr.Append("5")
			Convey("Test size", func() {
				So(arr.Size(), ShouldEqual, 6)
			})
		})
		Convey("Test GetValue", func() {
			value, err := arr.GetValue(4)
			So(err, ShouldBeNil)
			So(value, ShouldEqual, "4")
			value, err = arr.GetValue(100)
			So(err, ShouldNotBeNil)
			So(value, ShouldBeNil)
		})
		Convey("Test SetValue", func() {
			err := arr.SetValue("44", 4)
			So(err, ShouldBeNil)
			value, _ := arr.GetValue(4)
			So(value, ShouldEqual, "44")
			err = arr.SetValue(100, 100)
			So(err, ShouldNotBeNil)
		})
		Convey("Test remove", func() {
			err := arr.Remove(-1)
			So(err, ShouldNotBeNil)
			err = arr.Remove(100)
			So(err, ShouldNotBeNil)
			err = arr.Remove(1)
			So(err, ShouldBeNil)
			Convey("Test size", func() {
				So(arr.Size(), ShouldEqual, 5)
			})
		})

	})
}
