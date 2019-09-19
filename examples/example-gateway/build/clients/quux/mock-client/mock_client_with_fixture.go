// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
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

package clientmock

import (
	"github.com/golang/mock/gomock"
)

// MockIClientWithFixture is a mock of Client interface with preset fixture
type MockIClientWithFixture struct {
	*MockIClient
	fixture *ClientFixture

	echoMessageMock *EchoMessageMock
	echoStringMock  *EchoStringMock
}

// Call is a thin wrapper around gomock.Call for exposing the methods that do not mutate the fixture related information
// like Return().
type Call struct {
	call *gomock.Call
}

// MaxTimes marks a fixture as callable up to a maximum number of times.
func (c Call) MaxTimes(max int) {
	c.call.MaxTimes(max)
}

// MinTimes marks a fixture as must be called a minimum number of times.
func (c Call) MinTimes(max int) {
	c.call.MinTimes(max)
}

// New creates a new mock instance
func New(ctrl *gomock.Controller, fixture *ClientFixture) *MockIClientWithFixture {
	return &MockIClientWithFixture{
		MockIClient: NewMockIClient(ctrl),
		fixture:     fixture,
	}
}

// EXPECT shadows the EXPECT method on the underlying mock client.
// It should not be called directly.
func (m *MockIClientWithFixture) EXPECT() {
	panic("should not call EXPECT directly.")
}

// EchoMessageMock mocks the EchoMessage method
type EchoMessageMock struct {
	scenarios  *EchoMessageScenarios
	mockClient *MockIClient
}

// ExpectEchoMessage returns an object that allows the caller to choose expected scenario for EchoMessage
func (m *MockIClientWithFixture) ExpectEchoMessage() *EchoMessageMock {
	if m.echoMessageMock == nil {
		m.echoMessageMock = &EchoMessageMock{
			scenarios:  m.fixture.EchoMessage,
			mockClient: m.MockIClient,
		}
	}
	return m.echoMessageMock
}

// Success sets the expected scenario as defined in the concrete fixture package
// github.com/uber/zanzibar/examples/example-gateway/clients/quux/fixture
func (s *EchoMessageMock) Success() Call {
	f := s.scenarios.Success

	var arg0 interface{}
	arg0 = f.Arg0
	if f.Arg0Any {
		arg0 = gomock.Any()
	}

	ret0 := f.Ret0

	return Call{call: s.mockClient.EXPECT().EchoMessage(arg0).Return(ret0)}
}

// EchoStringMock mocks the EchoString method
type EchoStringMock struct {
	scenarios  *EchoStringScenarios
	mockClient *MockIClient
}

// ExpectEchoString returns an object that allows the caller to choose expected scenario for EchoString
func (m *MockIClientWithFixture) ExpectEchoString() *EchoStringMock {
	if m.echoStringMock == nil {
		m.echoStringMock = &EchoStringMock{
			scenarios:  m.fixture.EchoString,
			mockClient: m.MockIClient,
		}
	}
	return m.echoStringMock
}

// Success sets the expected scenario as defined in the concrete fixture package
// github.com/uber/zanzibar/examples/example-gateway/clients/quux/fixture
func (s *EchoStringMock) Success() Call {
	f := s.scenarios.Success

	var arg0 interface{}
	arg0 = f.Arg0
	if f.Arg0Any {
		arg0 = gomock.Any()
	}

	ret0 := f.Ret0

	return Call{call: s.mockClient.EXPECT().EchoString(arg0).Return(ret0)}
}
