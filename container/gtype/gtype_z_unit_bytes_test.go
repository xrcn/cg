package gtype_test

import (
	"testing"

	"github.com/xrcn/cg/container/gtype"
	"github.com/xrcn/cg/internal/json"
	"github.com/xrcn/cg/test/gtest"
	"github.com/xrcn/cg/util/gconv"
)

func Test_Bytes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		i := gtype.NewBytes([]byte("abc"))
		iClone := i.Clone()
		t.AssertEQ(iClone.Set([]byte("123")), []byte("abc"))
		t.AssertEQ(iClone.Val(), []byte("123"))

		// empty param test
		i1 := gtype.NewBytes()
		t.AssertEQ(i1.Val(), nil)

		i2 := gtype.NewBytes([]byte("abc"))
		t.Assert(i2.String(), "abc")

		copyVal := i2.DeepCopy()
		i2.Set([]byte("def"))
		t.AssertNE(copyVal, iClone.Val())
		i2 = nil
		copyVal = i2.DeepCopy()
		t.AssertNil(copyVal)
	})
}

func Test_Bytes_JSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		b := []byte("i love cg")
		i := gtype.NewBytes(b)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := gtype.NewBytes()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.Val(), b)
	})
}

func Test_Bytes_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *gtype.Bytes
	}
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "123",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), "123")
	})
}
