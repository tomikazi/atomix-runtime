// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
