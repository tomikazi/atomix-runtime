// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	counterv1 "github.com/atomix/atomix-runtime/api/atomix/runtime/counter/v1"
)

func newCounterV1(conn *localV1Conn) counterv1.CounterServer {
	return &counterV1{
		conn: conn,
	}
}

type counterV1 struct {
	conn *localV1Conn
}

func (s *counterV1) Set(ctx context.Context, request *counterv1.SetRequest) (*counterv1.SetResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *counterV1) Get(ctx context.Context, request *counterv1.GetRequest) (*counterv1.GetResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *counterV1) Increment(ctx context.Context, request *counterv1.IncrementRequest) (*counterv1.IncrementResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *counterV1) Decrement(ctx context.Context, request *counterv1.DecrementRequest) (*counterv1.DecrementResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ counterv1.CounterServer = (*counterV1)(nil)
