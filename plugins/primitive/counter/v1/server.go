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
	"github.com/atomix/atomix-runtime/pkg6/primitive"
	"github.com/atomix/atomix-runtime/pkg6/runtime"
)

// ProxyProvider is a counter provider
type ProxyProvider interface {
	GetCounter(ctx context.Context, name string) (counterv1.CounterServer, error)
}

func getCounter(provider ProxyProvider) func(ctx context.Context, name string) (counterv1.CounterServer, error) {
	return func(ctx context.Context, name string) (counterv1.CounterServer, error) {
		return provider.GetCounter(ctx, name)
	}
}

func newServer(runtime runtime.Runtime) counterv1.CounterServer {
	return &counterV1Server{
		proxies: primitive.NewProxyManager[ProxyProvider, counterv1.CounterServer](runtime, getCounter),
	}
}

type counterV1Server struct {
	proxies primitive.ProxyManager[ProxyProvider, counterv1.CounterServer]
}

func (s *counterV1Server) Set(ctx context.Context, request *counterv1.SetRequest) (*counterv1.SetResponse, error) {
	counter, err := s.proxies.GetProxy(ctx, request.Headers)
	if err != nil {
		return nil, err
	}
	return counter.Set(ctx, request)
}

func (s *counterV1Server) Get(ctx context.Context, request *counterv1.GetRequest) (*counterv1.GetResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *counterV1Server) Increment(ctx context.Context, request *counterv1.IncrementRequest) (*counterv1.IncrementResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *counterV1Server) Decrement(ctx context.Context, request *counterv1.DecrementRequest) (*counterv1.DecrementResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ counterv1.CounterServer = (*counterV1Server)(nil)
