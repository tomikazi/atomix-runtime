// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package primitive

import "github.com/atomix/atomix-runtime/pkg/runtime"

func NewPlugin(factory runtime.ComponentFactory[Primitive]) Plugin {
	return newPrimitivePlugin(factory)
}

type Plugin interface {
	runtime.Plugin[Primitive]
}

func newPrimitivePlugin(factory runtime.ComponentFactory[Primitive]) Plugin {
	return &primitivePlugin{
		factory: factory,
	}
}

type primitivePlugin struct {
	factory runtime.ComponentFactory[Primitive]
}

func (p *primitivePlugin) New(runtime runtime.Runtime) Primitive {
	return p.factory(runtime)
}

var _ Plugin = (*primitivePlugin)(nil)
