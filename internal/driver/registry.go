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
	"path/filepath"
	"plugin"
	"sync"
)

// Registry is a driver registry
type Registry interface {
	Lookup(name, version string) (driver.Driver, error)
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

	s, err := p.Lookup("Driver")
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
