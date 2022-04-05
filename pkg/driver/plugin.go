// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
