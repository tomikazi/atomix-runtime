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
