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
	counterv1 "github.com/atomix/atomix-runtime/api/atomix/runtime/counter/v1"
	"github.com/atomix/atomix-runtime/pkg6/driver"
	"github.com/golang/protobuf/proto"
)

func newConfiguration() Configuration {
	return Configuration{}
}

type Configuration struct {
	proto.Message
}

func newConn(ctx context.Context, configuration Configuration) (driver.Conn, error) {
	return &localV1Conn{}, nil
}

type localV1Conn struct{}

func (c *localV1Conn) GetCounter(ctx context.Context, name string) (counterv1.CounterServer, error) {
	return newCounterV1(c), nil
}

func (c *localV1Conn) Close() error {
	return nil
}

var _ driver.Conn = (*localV1Conn)(nil)
