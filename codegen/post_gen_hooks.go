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

package codegen

import (
	"fmt"
	"path"

	"github.com/pkg/errors"
	"path/filepath"
)

// ClientMockGenHook returns a PostGenHook to generate client mocks
func ClientMockGenHook(h *PackageHelper) (PostGenHook, error) {
	bin, err := NewMockgenBin(h.PackageRoot())
	if err != nil {
		return nil, errors.Wrap(err, "error building mockgen binary")
	}

	return func(instances map[string][]*ModuleInstance) error {
		fmt.Println("Generating client mocks:")
		mockCount := len(instances["client"])
		for i, instance := range instances["client"] {
			if err := bin.GenMock(
				instance,
				"mock-client/mock_client.go",
				"clientmock",
				"Client",
			); err != nil {
				return errors.Wrapf(
					err,
					"error generating mocks for client %q",
					instance.InstanceName,
				)
			}
			fmt.Printf(
				"Generating %12s %12s %-30s in %-50s %d/%d\n",
				"mock",
				instance.ClassName,
				instance.InstanceName,
				path.Join(path.Base(h.CodeGenTargetPath()), instance.Directory, "mock-client"),
				i+1, mockCount,
			)
		}
		return nil
	}, nil
}

// ServiceMockGenHook returns a PostGenHook to generate server mocks
func ServiceMockGenHook(h *PackageHelper, t *Template) PostGenHook {
	return func(instances map[string][]*ModuleInstance) error {
		fmt.Println("Generating service mocks:")
		mockCount := len(instances["service"])
		for i, instance := range instances["service"] {
			mockInit, err := generateMockInitializer(instance, h, t)
			if err != nil {
				return errors.Wrapf(
					err,
					"Error generating service mock_init.go for %s",
					instance.InstanceName,
				)
			}
			mockService, err := t.ExecTemplate("service_mock.tmpl", instance, h)
			if err != nil {
				return errors.Wrapf(
					err,
					"Error generating service mock_service.go for %s",
					instance.InstanceName,
				)
			}

			buildPath := filepath.Join(h.CodeGenTargetPath(), instance.Directory)
			mockInitPath := filepath.Join(buildPath, "mock-service/mock_init.go")
			mockServicePath := filepath.Join(buildPath, "mock-service/mock_Service.go")

			for _, s := range []struct {
				path    string
				content []byte
			}{
				{mockInitPath, mockInit},
				{mockServicePath, mockService},
			} {
				if err := writeFile(s.path, s.content); err != nil {
					return errors.Wrapf(
						err,
						"Error writing to file %q",
						s.path,
					)
				}
				if err := FormatGoFile(s.path); err != nil {
					return err
				}

			}

			fmt.Printf(
				"Generating %12s %12s %-30s in %-50s %d/%d\n",
				"mock",
				instance.ClassName,
				instance.InstanceName,
				path.Join(path.Base(h.CodeGenTargetPath()), instance.Directory, "mock-service"),
				i+1, mockCount,
			)
		}
		return nil
	}
}

// generateMockInitializer generates code to initialize modules with leaf nodes being mocks
func generateMockInitializer(instance *ModuleInstance, h *PackageHelper, t *Template) ([]byte, error) {
	leafWithFixture := map[string]string{}
	for _, leaf := range instance.RecursiveDependencies["client"] {
		spec, ok := leaf.genSpec.(*ClientSpec)
		if ok && spec.Fixture != nil {
			leafWithFixture[leaf.InstanceName] = spec.Fixture.ImportPath
		}
	}
	data := map[string]interface{}{
		"Instance":        instance,
		"LeafWithFixture": leafWithFixture,
	}
	return t.ExecTemplate("module_mock_initializer.tmpl", data, h)
}
