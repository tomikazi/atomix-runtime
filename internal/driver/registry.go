// Copyright 2022-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package driver

import (
	"fmt"
	"github.com/atomix/atomix-runtime/pkg/driver"
	"io/fs"
	"path/filepath"
	"plugin"
	"strings"
	"sync"
)

// Registry is a driver registry
type Registry interface {
	Lookup(name, version string) (driver.Driver, error)
	List() ([]driver.Driver, error)
}

func NewPluginRegistry(path string) Registry {
	return &pluginRegistry{
		path:    path,
		drivers: make(map[string]map[string]driver.Driver),
	}
}

type pluginRegistry struct {
	path    string
	drivers map[string]map[string]driver.Driver
	mu      sync.Mutex
}

func (r *pluginRegistry) Lookup(name, version string) (driver.Driver, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	versions, ok := r.drivers[name]
	if ok {
		d, ok := versions[version]
		if ok {
			return d, nil
		}
	}

	p, err := plugin.Open(filepath.Join(r.path, fmt.Sprintf("%s.%s.so", name, version)))
	if err != nil {
		return nil, err
	}

	s, err := p.Lookup(fmt.Sprintf("%s%s", strings.ToTitle(name), strings.ToTitle(version)))
	if err != nil {
		return nil, err
	}

	d := s.(driver.Driver)

	versions, ok = r.drivers[name]
	if !ok {
		versions = make(map[string]driver.Driver)
		r.drivers[name] = versions
	}
	versions[version] = d
	return d, nil
}

func (r *pluginRegistry) List() ([]driver.Driver, error) {
	var drivers []driver.Driver
	err := filepath.Walk(r.path, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		parts := strings.Split(info.Name(), ".")
		if len(parts) != 3 {
			return nil
		}
		if parts[2] != "so" {
			return nil
		}
		name, version := parts[0], parts[1]
		driver, err := r.Lookup(name, version)
		if err != nil {
			return err
		}
		drivers = append(drivers, driver)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return drivers, nil
}
