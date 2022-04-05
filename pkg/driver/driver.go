// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"github.com/atomix/atomix-runtime/pkg/logging"
	"github.com/atomix/atomix-runtime/pkg/runtime"
)

var log = logging.GetLogger("atomix", "driver")

// Driver is a driver plugin
type Driver interface {
	runtime.Component
	Connect(ctx context.Context, store string) (Conn, error)
}

// Provider is an interface for providers of drivers for a store
type Provider interface {
	GetDriver(ctx context.Context, store string) (Driver, error)
}
