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

package runtime

import (
	"github.com/atomix/atomix-runtime/pkg/component"
	"sync"
)

var GlobalRegistry Registry

func init() {
	GlobalRegistry = NewRegistry()
}

type Registry interface {
	RegisterComponent(component component.Component)
	GetComponent(name, version string) (component.Component, bool)
}

func NewRegistry() Registry {
	return &registry{
		components: make(map[string]map[string]component.Component),
	}
}

type registry struct {
	components map[string]map[string]component.Component
	mu         sync.RWMutex
}

func (r *registry) RegisterComponent(comp component.Component) {
	r.mu.Lock()
	defer r.mu.Unlock()
	versions, ok := r.components[comp.Name()]
	if !ok {
		versions = make(map[string]component.Component)
		r.components[comp.Name()] = versions
	}
	if _, ok := versions[comp.Version()]; !ok {
		versions[comp.Version()] = comp
	}
}

func (r *registry) GetComponent(name, version string) (component.Component, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	versions, ok := r.components[name]
	if !ok {
		return nil, false
	}
	comp, ok := versions[version]
	if !ok {
		return nil, false
	}
	return comp, true
}
