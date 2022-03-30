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
	"context"
	"fmt"
	"github.com/atomix/atomix-runtime/pkg6/driver"
	"github.com/atomix/atomix-runtime/pkg6/runtime"
)

const (
	Name    = "Local"
	Version = "v1"
)

var Plugin = driver.NewPlugin(func(runtime runtime.Runtime) driver.Driver {
	return newLocalV1Driver(runtime)
})

func newLocalV1Driver(runtime runtime.Runtime) driver.Driver {
	return &localV1Driver{
		conns: driver.NewConnManager[Configuration](runtime, newConn, newConfiguration),
	}
}

type localV1Driver struct {
	conns driver.ConnManager
}

func (d *localV1Driver) Name() string {
	return Name
}

func (d *localV1Driver) Version() string {
	return Version
}

func (d *localV1Driver) Connect(ctx context.Context, store string) (driver.Conn, error) {
	return d.conns.GetConn(ctx, store)
}

func (d *localV1Driver) String() string {
	return fmt.Sprintf("%s/%s", d.Name(), d.Version())
}

var _ driver.Driver = (*localV1Driver)(nil)
