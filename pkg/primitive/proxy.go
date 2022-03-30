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

package primitive

import (
	"context"
	metav1 "github.com/atomix/atomix-runtime/api/atomix/runtime/meta/v1"
	"github.com/atomix/atomix-runtime/pkg/errors"
	"github.com/atomix/atomix-runtime/pkg/runtime"
)

type ProxyProvider[S any] interface{}

type ProxyGetter[P ProxyProvider[S], S any] func(P) func(ctx context.Context, name string) (S, error)

type ProxyManager[P ProxyProvider[S], S any] interface {
	GetProxy(ctx context.Context, headers metav1.RequestHeaders) (S, error)
}

func NewProxyManager[P ProxyProvider[S], S any](runtime runtime.Runtime, getter ProxyGetter[P, S]) ProxyManager[P, S] {
	return &primitiveClient[P, S]{
		runtime: runtime,
		getter:  getter,
	}
}

type primitiveClient[P ProxyProvider[S], S any] struct {
	runtime runtime.Runtime
	getter  ProxyGetter[P, S]
}

func (c *primitiveClient[P, S]) GetProxy(ctx context.Context, headers metav1.RequestHeaders) (S, error) {
	driver, err := c.runtime.GetDriver(ctx, headers.Store)
	if err != nil {
		return nil, err
	}

	conn, err := driver.Connect(ctx, headers.Store)
	if err != nil {
		return nil, err
	}

	provider, ok := conn.(P)
	if !ok {
		return nil, errors.NewNotSupported("Counter type not supported by driver '%s/%s'", driver.Name(), driver.Version())
	}
	return c.getter(provider)(ctx, headers.Primitive)
}

var _ ProxyManager[ProxyProvider[any], any] = (*primitiveClient[ProxyProvider[any], any])(nil)
