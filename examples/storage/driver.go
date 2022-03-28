// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package storage

import (
	"fmt"
	"github.com/atomix/atomix-runtime/pkg/runtime"
	"github.com/atomix/atomix-runtime/pkg/storage"
)

const Name = "example"
const Version = "v1"

func init() {
	runtime.Register(Driver{})
}

type Driver struct{}

func (d Driver) Name() string {
	return Name
}

func (d Driver) Version() string {
	return Version
}

func (d Driver) NewProxy(config storage.ProxyConfig) storage.Proxy {
	//TODO implement me
	panic("implement me")
}

func (d Driver) String() string {
	return fmt.Sprintf("%s/%s", Name, Version)
}

var _ storage.Driver = Driver{}
