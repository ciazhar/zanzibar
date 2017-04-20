// Code generated by zanzibar
// @generated
// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package googlenow

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients/googlenow"
)

// AddCredentialsHandler is the handler for "/googlenow/add-credentials"
type AddCredentialsHandler struct {
	Clients *clients.Clients
}

// NewAddCredentialsEndpoint creates a handler
func NewAddCredentialsEndpoint(
	gateway *zanzibar.Gateway,
) *AddCredentialsHandler {
	return &AddCredentialsHandler{
		Clients: gateway.Clients.(*clients.Clients),
	}
}

// HandleRequest handles "/googlenow/add-credentials".
func (handler *AddCredentialsHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	if !req.CheckHeaders([]string{"x-uuid", "x-token"}) {
		return
	}
	var requestBody AddCredentialsHTTPRequest
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	workflow := AddCredentialsEndpoint{
		Clients: handler.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		req.Logger.Warn("Workflow for endpoint returned error",
			zap.String("error", err.Error()),
		)
		res.SendErrorString(500, "Unexpected server error")
		return
	}
	// TODO(sindelar): implement check headers on response

	res.WriteJSONBytes(202, cliRespHeaders, nil)
}

// AddCredentialsEndpoint calls thrift client GoogleNow.AddCredentials
type AddCredentialsEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w AddCredentialsEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.ServerHeaderInterface,
	r *AddCredentialsHTTPRequest,
) (zanzibar.ServerHeaderInterface, error) {
	clientRequest := convertToAddCredentialsClientRequest(r)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("X-Uuid")
	if ok {
		clientHeaders["X-Uuid"] = h
	}

	cliRespHeaders, err := w.Clients.GoogleNow.AddCredentials(
		ctx, clientHeaders, clientRequest,
	)
	if err != nil {
		w.Logger.Warn("Could not make client request",
			zap.String("error", err.Error()),
		)
		// TODO(sindelar): Consider returning partial headers in error case.
		return nil, err
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	resHeaders.Set("X-Uuid", cliRespHeaders["X-Uuid"])

	return resHeaders, nil
}

func convertToAddCredentialsClientRequest(body *AddCredentialsHTTPRequest) *googlenowClient.AddCredentialsHTTPRequest {
	clientRequest := &googlenowClient.AddCredentialsHTTPRequest{}

	clientRequest.AuthCode = string(body.AuthCode)

	return clientRequest
}
