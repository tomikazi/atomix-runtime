// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"fmt"
	counterv1 "github.com/atomix/atomix-runtime/api/atomix/runtime/counter/v1"
	"github.com/atomix/atomix-runtime/pkg6/primitive"
	"github.com/atomix/atomix-runtime/pkg6/runtime"
)

const (
	Name    = "Counter"
	Version = "v1"
)

var Plugin = primitive.NewPlugin(func(runtime runtime.Runtime) primitive.Primitive {
	return newCounterV1(runtime)
})

func newCounterV1(runtime runtime.Runtime) primitive.Primitive {
	counterv1.RegisterCounterServer(runtime.Server(), newServer(runtime))
	return &counterV1{}
}

type counterV1 struct{}

func (t *counterV1) Name() string {
	return Name
}

func (t *counterV1) Version() string {
	return Version
}

func (t *counterV1) Init(runtime runtime.Runtime) error {
	counterv1.RegisterCounterServer(runtime.Server(), newServer(runtime))
	return nil
}

func (t *counterV1) String() string {
	return fmt.Sprintf("%s/%s", t.Name(), t.Version())
}

var _ primitive.Primitive = (*counterV1)(nil)
