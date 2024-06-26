package gtime_test

import (
	"testing"

	"github.com/xrcn/cg/os/gtime"
	"github.com/xrcn/cg/test/gtest"
)

// 1681
func Test_Issue1681(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gtime.New("2022-03-08T03:01:14-07:00").Local().Time, gtime.New("2022-03-08T10:01:14Z").Local().Time)
		t.Assert(gtime.New("2022-03-08T03:01:14-08:00").Local().Time, gtime.New("2022-03-08T11:01:14Z").Local().Time)
		t.Assert(gtime.New("2022-03-08T03:01:14-09:00").Local().Time, gtime.New("2022-03-08T12:01:14Z").Local().Time)
		t.Assert(gtime.New("2022-03-08T03:01:14+08:00").Local().Time, gtime.New("2022-03-07T19:01:14Z").Local().Time)
	})
}

// 2803
func Test_Issue2803(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		newTime := gtime.New("2023-07-26").LayoutTo("2006-01")
		t.Assert(newTime.Year(), 2023)
		t.Assert(newTime.Month(), 7)
		t.Assert(newTime.Day(), 1)
		t.Assert(newTime.Hour(), 0)
		t.Assert(newTime.Minute(), 0)
		t.Assert(newTime.Second(), 0)
	})
}
