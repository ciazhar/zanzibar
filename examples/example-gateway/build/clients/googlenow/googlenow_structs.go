package googlenowClient

import (
	"runtime"

	"github.com/uber/zanzibar/runtime"
)

func getDirName() string {
	_, file, _, _ := runtime.Caller(0)
	return zanzibar.GetDirnameFromRuntimeCaller(file)
}

// AddCredentialsHTTPRequest is the http body type for endpoint addCredentials.
type AddCredentialsHTTPRequest struct {
	AuthCode string `json:"authCode"`
}
