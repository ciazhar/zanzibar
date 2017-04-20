package googlenowClient

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
	"github.com/uber/zanzibar/runtime"
)

// GoogleNowClient is the http client for service GoogleNow.
type GoogleNowClient struct {
	ClientID   string
	HTTPClient *zanzibar.HTTPClient
}

// NewClient returns a new http client for service GoogleNow.
func NewClient(
	gateway *zanzibar.Gateway,
) *GoogleNowClient {
	ip := gateway.Config.MustGetString("clients.googleNow.ip")
	port := gateway.Config.MustGetInt("clients.googleNow.port")

	baseURL := "http://" + ip + ":" + strconv.Itoa(int(port))
	return &GoogleNowClient{
		ClientID:   "googlenow",
		HTTPClient: zanzibar.NewHTTPClient(gateway, baseURL),
	}
}

// AddCredentials calls "/add-credentials" endpoint.
func (c *GoogleNowClient) AddCredentials(
	ctx context.Context,
	headers map[string]string,
	r *AddCredentialsHTTPRequest,
) (map[string]string, error) {

	req := zanzibar.NewClientHTTPRequest(
		c.ClientID, "addCredentials", c.HTTPClient,
	)

	// Generate full URL.
	fullURL := c.HTTPClient.BaseURL + "/add-credentials"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{202})

	switch res.StatusCode {
	case 202:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return respHeaders, err
		}
		return respHeaders, nil
	}

	return respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// CheckCredentials calls "/check-credentials" endpoint.
func (c *GoogleNowClient) CheckCredentials(
	ctx context.Context,
	headers map[string]string,
) (map[string]string, error) {

	req := zanzibar.NewClientHTTPRequest(
		c.ClientID, "checkCredentials", c.HTTPClient,
	)

	// Generate full URL.
	fullURL := c.HTTPClient.BaseURL + "/check-credentials"

	err := req.WriteJSON("POST", fullURL, headers, nil)
	if err != nil {
		return nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{202})

	switch res.StatusCode {
	case 202:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return respHeaders, err
		}
		return respHeaders, nil
	}

	return respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}
