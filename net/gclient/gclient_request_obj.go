package gclient

import (
	"context"
	"net/http"
	"reflect"

	"github.com/xrc360/golang/errors/gcode"
	"github.com/xrc360/golang/errors/gerror"
	"github.com/xrc360/golang/text/gregex"
	"github.com/xrc360/golang/text/gstr"
	"github.com/xrc360/golang/util/gconv"
	"github.com/xrc360/golang/util/gmeta"
	"github.com/xrc360/golang/util/gtag"
	"github.com/xrc360/golang/util/gutil"
)

// DoRequestObj does HTTP request using standard request/response object.
// The request object `req` is defined like:
//
//	type UseCreateReq struct {
//	    g.Meta `path:"/user" method:"put"`
//	    // other fields....
//	}
//
// The response object `res` should be a pointer type. It automatically converts result
// to given object `res` is success.
//
// Example:
// var (
//
//	req = UseCreateReq{}
//	res *UseCreateRes
//
// )
//
// err := DoRequestObj(ctx, req, &res)
func (c *Client) DoRequestObj(ctx context.Context, req, res interface{}) error {
	var (
		method = gmeta.Get(req, gtag.Method).String()
		path   = gmeta.Get(req, gtag.Path).String()
	)
	if method == "" {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`no "%s" tag found in request object: %s`,
			gtag.Method, reflect.TypeOf(req).String(),
		)
	}
	if path == "" {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`no "%s" tag found in request object: %s`,
			gtag.Path, reflect.TypeOf(req).String(),
		)
	}
	path = c.handlePathForObjRequest(path, req)
	switch gstr.ToUpper(method) {
	case
		http.MethodGet,
		http.MethodPut,
		http.MethodPost,
		http.MethodDelete,
		http.MethodHead,
		http.MethodPatch,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace:
		if result := c.RequestVar(ctx, method, path, req); res != nil && !result.IsEmpty() {
			return result.Scan(res)
		}
		return nil

	default:
		return gerror.Newf(`invalid HTTP method "%s"`, method)
	}
}

// handlePathForObjRequest replaces parameters in `path` with parameters from request object.
// Eg:
// /order/{id}  -> /order/1
// /user/{name} -> /order/john
func (c *Client) handlePathForObjRequest(path string, req interface{}) string {
	if gstr.Contains(path, "{") {
		requestParamsMap := gconv.Map(req)
		if len(requestParamsMap) > 0 {
			path, _ = gregex.ReplaceStringFuncMatch(`\{(\w+)\}`, path, func(match []string) string {
				foundKey, foundValue := gutil.MapPossibleItemByKey(requestParamsMap, match[1])
				if foundKey != "" {
					return gconv.String(foundValue)
				}
				return match[0]
			})
		}
	}
	return path
}
