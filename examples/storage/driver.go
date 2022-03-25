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
