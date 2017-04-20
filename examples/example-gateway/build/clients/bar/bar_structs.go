package barClient

import (
	"runtime"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	clientsFooFoo "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/foo/foo"
	"github.com/uber/zanzibar/runtime"
)

func getDirName() string {
	_, file, _, _ := runtime.Caller(0)
	return zanzibar.GetDirnameFromRuntimeCaller(file)
}

// ArgNotStructHTTPRequest is the http body type for endpoint argNotStruct.
type ArgNotStructHTTPRequest struct {
	Request string `json:"request"`
}

// NormalHTTPRequest is the http body type for endpoint normal.
type NormalHTTPRequest struct {
	Request *clientsBarBar.BarRequest `json:"request"`
}

// TooManyArgsHTTPRequest is the http body type for endpoint tooManyArgs.
type TooManyArgsHTTPRequest struct {
	Request *clientsBarBar.BarRequest `json:"request"`
	Foo     *clientsFooFoo.FooStruct  `json:"foo"`
}
