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

import "github.com/atomix/atomix-runtime/pkg/runtime"

func NewPlugin(factory runtime.ComponentFactory[Driver]) Plugin {
	return newDriverPlugin(factory)
}

type Plugin interface {
	runtime.Plugin[Driver]
}

func newDriverPlugin(factory runtime.ComponentFactory[Driver]) Plugin {
	return &driverPlugin{
		factory: factory,
	}
}

type driverPlugin struct {
	factory runtime.ComponentFactory[Driver]
}

func (p *driverPlugin) New(runtime runtime.Runtime) Driver {
	return p.factory(runtime)
}

var _ Plugin = (*driverPlugin)(nil)
