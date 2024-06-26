package ghttp_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/xrcn/cg/frame/g"
	"github.com/xrcn/cg/net/ghttp"
	"github.com/xrcn/cg/test/gtest"
	"github.com/xrcn/cg/util/guid"
)

type testWrapStdHTTPStruct struct {
	T    *gtest.T
	text string
}

func (t *testWrapStdHTTPStruct) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t.T.Assert(req.Method, "POST")
	t.T.Assert(req.URL.Path, "/api/wraph")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, t.text)
}

func Test_Server_Wrap_Handler(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		str1 := "hello"
		str2 := "hello again"
		s.Group("/api", func(group *ghttp.RouterGroup) {
			group.GET("/wrapf", ghttp.WrapF(func(w http.ResponseWriter, req *http.Request) {
				t.Assert(req.Method, "GET")
				t.Assert(req.URL.Path, "/api/wrapf")
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, str1)
			}))

			group.POST("/wraph", ghttp.WrapH(&testWrapStdHTTPStruct{t, str2}))
		})

		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		time.Sleep(100 * time.Millisecond)
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d/api", s.GetListenedPort()))

		response, err := client.Get(ctx, "/wrapf")
		t.AssertNil(err)
		defer response.Close()
		t.Assert(response.StatusCode, http.StatusBadRequest)
		t.Assert(response.ReadAllString(), str1)

		response2, err := client.Post(ctx, "/wraph")
		t.AssertNil(err)
		defer response2.Close()
		t.Assert(response2.StatusCode, http.StatusInternalServerError)
		t.Assert(response2.ReadAllString(), str2)
	})
}
