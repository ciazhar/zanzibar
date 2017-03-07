// Code generated by zanzibar
// @generated

package bar_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/zanzibar/test/lib/test_gateway"
)

func TestNoRequestSuccessfulRequestOKResponse(t *testing.T) {
	var counter int = 0

	gateway, err := testGateway.CreateGateway(t, nil, &testGateway.Options{
		KnownBackends: []string{"bar"},
	})
	if !assert.NoError(t, err, "got bootstrap err") {
		return
	}
	defer gateway.Close()

	res, err := gateway.MakeRequest(
		"GET", "/bar/no-request-path", bytes.NewReader([]byte("{}")),
	)
	if !assert.NoError(t, err, "got http error") {
		return
	}

	assert.Equal(t, "200 OK", res.Status)
	assert.Equal(t, 1, counter)
}