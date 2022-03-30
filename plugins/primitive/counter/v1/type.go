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
