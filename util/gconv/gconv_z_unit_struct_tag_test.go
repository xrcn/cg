package gconv_test

import (
	"testing"

	"github.com/xrcn/cg/frame/g"
	"github.com/xrcn/cg/test/gtest"
	"github.com/xrcn/cg/util/gconv"
)

func Test_StructTag(t *testing.T) {
	type User struct {
		Uid   int
		Name  string
		Pass1 string `orm:"password1"`
		Pass2 string `orm:"password2"`
	}
	gtest.C(t, func(t *gtest.T) {
		user := new(User)
		params1 := g.Map{
			"uid":       1,
			"Name":      "john",
			"password1": "123",
			"password2": "456",
		}
		if err := gconv.Struct(params1, user); err != nil {
			t.Error(err)
		}
		t.Assert(user, &User{
			Uid:   1,
			Name:  "john",
			Pass1: "",
			Pass2: "",
		})
	})
	gtest.C(t, func(t *gtest.T) {
		user := new(User)
		params1 := g.Map{
			"uid":       1,
			"Name":      "john",
			"password1": "123",
			"password2": "456",
		}
		if err := gconv.StructTag(params1, user, "orm"); err != nil {
			t.Error(err)
		}
		t.Assert(user, &User{
			Uid:   1,
			Name:  "john",
			Pass1: "123",
			Pass2: "456",
		})
	})
	gtest.C(t, func(t *gtest.T) {
		user := new(User)
		params2 := g.Map{
			"uid":       2,
			"name":      "smith",
			"password1": "111",
			"password2": "222",
		}
		if err := gconv.StructTag(params2, user, "orm"); err != nil {
			t.Error(err)
		}
		t.Assert(user, &User{
			Uid:   2,
			Name:  "smith",
			Pass1: "111",
			Pass2: "222",
		})
	})
}
