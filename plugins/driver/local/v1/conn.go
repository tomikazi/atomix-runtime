// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
