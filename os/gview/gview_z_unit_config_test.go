package gview_test

import (
	"context"
	"testing"

	"github.com/xrcn/cg/frame/g"
	"github.com/xrcn/cg/i18n/gi18n"
	"github.com/xrcn/cg/internal/command"
	"github.com/xrcn/cg/os/gview"
	"github.com/xrcn/cg/test/gtest"
)

func Test_Config(t *testing.T) {
	// show error print
	command.Init("-cg.gview.errorprint=true")
	gtest.C(t, func(t *gtest.T) {
		config := gview.Config{
			Paths: []string{gtest.DataPath("config")},
			Data: g.Map{
				"name": "cg",
			},
			DefaultFile: "test.html",
			Delimiters:  []string{"${", "}"},
		}

		view := gview.New()
		err := view.SetConfig(config)
		t.AssertNil(err)

		view.SetI18n(gi18n.New())

		str := `hello ${.name},version:${.version}`
		view.Assigns(g.Map{"version": "1.7.0"})
		result, err := view.ParseContent(context.TODO(), str, nil)
		t.AssertNil(err)
		t.Assert(result, "hello cg,version:1.7.0")

		result, err = view.ParseDefault(context.TODO())
		t.AssertNil(err)
		t.Assert(result, "name:cg")

		t.Assert(view.GetDefaultFile(), "test.html")
	})
	// SetConfig path fail: notexist
	gtest.C(t, func(t *gtest.T) {
		config := gview.Config{
			Paths: []string{"notexist", gtest.DataPath("config/test.html")},
			Data: g.Map{
				"name": "cg",
			},
			DefaultFile: "test.html",
			Delimiters:  []string{"${", "}"},
		}

		view := gview.New()
		err := view.SetConfig(config)
		t.AssertNE(err, nil)
	})
	// SetConfig path fail: set file path
	gtest.C(t, func(t *gtest.T) {
		config := gview.Config{
			Paths: []string{gtest.DataPath("config/test.html")},
			Data: g.Map{
				"name": "cg",
			},
			DefaultFile: "test.html",
			Delimiters:  []string{"${", "}"},
		}

		view := gview.New()
		err := view.SetConfig(config)
		t.AssertNE(err, nil)
	})
}

func Test_ConfigWithMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		view := gview.New()
		err := view.SetConfigWithMap(g.Map{
			"Paths":       []string{gtest.DataPath("config")},
			"DefaultFile": "test.html",
			"Delimiters":  []string{"${", "}"},
			"Data": g.Map{
				"name": "cg",
			},
		})
		t.AssertNil(err)

		str := `hello ${.name},version:${.version}`
		view.Assigns(g.Map{"version": "1.7.0"})
		result, err := view.ParseContent(context.TODO(), str, nil)
		t.AssertNil(err)
		t.Assert(result, "hello cg,version:1.7.0")

		result, err = view.ParseDefault(context.TODO())
		t.AssertNil(err)
		t.Assert(result, "name:cg")
	})
	// path as paths
	gtest.C(t, func(t *gtest.T) {
		view := gview.New()
		err := view.SetConfigWithMap(g.Map{
			"Path":        gtest.DataPath("config"),
			"DefaultFile": "test.html",
			"Delimiters":  []string{"${", "}"},
			"Data": g.Map{
				"name": "cg",
			},
		})
		t.AssertNil(err)

		str := `hello ${.name},version:${.version}`
		view.Assigns(g.Map{"version": "1.7.0"})
		result, err := view.ParseContent(context.TODO(), str, nil)
		t.AssertNil(err)
		t.Assert(result, "hello cg,version:1.7.0")

		result, err = view.ParseDefault(context.TODO())
		t.AssertNil(err)
		t.Assert(result, "name:cg")
	})
	// path as paths
	gtest.C(t, func(t *gtest.T) {
		view := gview.New()
		err := view.SetConfigWithMap(g.Map{
			"Path":        []string{gtest.DataPath("config")},
			"DefaultFile": "test.html",
			"Delimiters":  []string{"${", "}"},
			"Data": g.Map{
				"name": "cg",
			},
		})
		t.AssertNil(err)

		str := `hello ${.name},version:${.version}`
		view.Assigns(g.Map{"version": "1.7.0"})
		result, err := view.ParseContent(context.TODO(), str, nil)
		t.AssertNil(err)
		t.Assert(result, "hello cg,version:1.7.0")

		result, err = view.ParseDefault(context.TODO())
		t.AssertNil(err)
		t.Assert(result, "name:cg")
	})
	// map is nil
	gtest.C(t, func(t *gtest.T) {
		view := gview.New()
		err := view.SetConfigWithMap(nil)
		t.AssertNE(err, nil)
	})
}
