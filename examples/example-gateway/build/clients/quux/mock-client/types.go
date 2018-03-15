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
	base "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/foo/base/base"
)

// ClientFixture defines the client fixture type
type ClientFixture struct {
	EchoMessage *EchoMessageScenarios
	EchoString  *EchoStringScenarios
}

// EchoMessageScenarios defines all fixture scenarios for EchoMessage
type EchoMessageScenarios struct {
	Success *EchoMessageFixture `scenario:"success"`
}

// EchoStringScenarios defines all fixture scenarios for EchoString
type EchoStringScenarios struct {
	Success *EchoStringFixture `scenario:"success"`
}

// EchoMessageFixture defines the fixture type for EchoMessage
type EchoMessageFixture struct {
	Arg0 *base.Message

	// Arg{n}Any indicates the nth argument could be gomock.Any
	Arg0Any bool

	Ret0 *base.Message
}

// EchoStringFixture defines the fixture type for EchoString
type EchoStringFixture struct {
	Arg0 string

	// Arg{n}Any indicates the nth argument could be gomock.Any
	Arg0Any bool

	Ret0 string
}
