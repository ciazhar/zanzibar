// Code generated by mockery v1.0.0
package googlenowClient

import context "context"
import googlenow "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/googlenow/googlenow"
import mock "github.com/stretchr/testify/mock"
import zanzibar "github.com/uber/zanzibar/runtime"

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

// AddCredentials provides a mock function with given fields: ctx, reqHeaders, args
func (_m *MockClient) AddCredentials(ctx context.Context, reqHeaders map[string]string, args *googlenow.GoogleNowService_AddCredentials_Args) (map[string]string, error) {
	ret := _m.Called(ctx, reqHeaders, args)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string, *googlenow.GoogleNowService_AddCredentials_Args) map[string]string); ok {
		r0 = rf(ctx, reqHeaders, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, map[string]string, *googlenow.GoogleNowService_AddCredentials_Args) error); ok {
		r1 = rf(ctx, reqHeaders, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckCredentials provides a mock function with given fields: ctx, reqHeaders
func (_m *MockClient) CheckCredentials(ctx context.Context, reqHeaders map[string]string) (map[string]string, error) {
	ret := _m.Called(ctx, reqHeaders)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string) map[string]string); ok {
		r0 = rf(ctx, reqHeaders)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, map[string]string) error); ok {
		r1 = rf(ctx, reqHeaders)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HTTPClient provides a mock function with given fields:
func (_m *MockClient) HTTPClient() *zanzibar.HTTPClient {
	ret := _m.Called()

	var r0 *zanzibar.HTTPClient
	if rf, ok := ret.Get(0).(func() *zanzibar.HTTPClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zanzibar.HTTPClient)
		}
	}

	return r0
}
