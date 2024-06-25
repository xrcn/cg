package g

import (
	"github.com/xrc360/golang/database/gdb"
	"github.com/xrc360/golang/database/gredis"
	"github.com/xrc360/golang/frame/gins"
	"github.com/xrc360/golang/i18n/gi18n"
	"github.com/xrc360/golang/net/gclient"
	"github.com/xrc360/golang/net/ghttp"
	"github.com/xrc360/golang/net/gtcp"
	"github.com/xrc360/golang/net/gudp"
	"github.com/xrc360/golang/os/gcfg"
	"github.com/xrc360/golang/os/glog"
	"github.com/xrc360/golang/os/gres"
	"github.com/xrc360/golang/os/gview"
	"github.com/xrc360/golang/util/gvalid"
)

// Client is a convenience function, which creates and returns a new HTTP client.
func Client() *gclient.Client {
	return gclient.New()
}

// Server returns an instance of http server with specified name.
func Server(name ...interface{}) *ghttp.Server {
	return gins.Server(name...)
}

// TCPServer returns an instance of tcp server with specified name.
func TCPServer(name ...interface{}) *gtcp.Server {
	return gtcp.GetServer(name...)
}

// UDPServer returns an instance of udp server with specified name.
func UDPServer(name ...interface{}) *gudp.Server {
	return gudp.GetServer(name...)
}

// View returns an instance of template engine object with specified name.
func View(name ...string) *gview.View {
	return gins.View(name...)
}

// Config returns an instance of config object with specified name.
func Config(name ...string) *gcfg.Config {
	return gins.Config(name...)
}

// Cfg is alias of Config.
// See Config.
func Cfg(name ...string) *gcfg.Config {
	return Config(name...)
}

// Resource returns an instance of Resource.
// The parameter `name` is the name for the instance.
func Resource(name ...string) *gres.Resource {
	return gins.Resource(name...)
}

// I18n returns an instance of gi18n.Manager.
// The parameter `name` is the name for the instance.
func I18n(name ...string) *gi18n.Manager {
	return gins.I18n(name...)
}

// Res is alias of Resource.
// See Resource.
func Res(name ...string) *gres.Resource {
	return Resource(name...)
}

// Log returns an instance of glog.Logger.
// The parameter `name` is the name for the instance.
func Log(name ...string) *glog.Logger {
	return gins.Log(name...)
}

// DB returns an instance of database ORM object with specified configuration group name.
func DB(name ...string) gdb.DB {
	return gins.Database(name...)
}

// Model creates and returns a model based on configuration of default database group.
func Model(tableNameOrStruct ...interface{}) *gdb.Model {
	return DB().Model(tableNameOrStruct...)
}

// ModelRaw creates and returns a model based on a raw sql not a table.
func ModelRaw(rawSql string, args ...interface{}) *gdb.Model {
	return DB().Raw(rawSql, args...)
}

// Redis returns an instance of redis client with specified configuration group name.
func Redis(name ...string) *gredis.Redis {
	return gins.Redis(name...)
}

// Validator is a convenience function, which creates and returns a new validation manager object.
func Validator() *gvalid.Validator {
	return gvalid.New()
}
