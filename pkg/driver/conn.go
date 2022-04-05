// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	controllerv1 "github.com/atomix/atomix-runtime/api/atomix/controller/v1"
	"github.com/atomix/atomix-runtime/pkg/errors"
	"github.com/atomix/atomix-runtime/pkg/runtime"
	"github.com/golang/protobuf/proto"
	"io"
	"sync"
)

// ConnManager is a driver connection manager
type ConnManager interface {
	GetConn(ctx context.Context, store string) (Conn, error)
}

func NewConnManager[C proto.Message](runtime runtime.Runtime, connector Connector[C], configProvider ConfigurationProvider[C]) ConnManager {
	return &connManager[C]{
		runtime:        runtime,
		connector:      connector,
		configProvider: configProvider,
	}
}

type connManager[C proto.Message] struct {
	runtime        runtime.Runtime
	connector      Connector[C]
	configProvider ConfigurationProvider[C]
	conns          map[string]Conn
	mu             sync.RWMutex
}

func (m *connManager[C]) GetConn(ctx context.Context, store string) (Conn, error) {
	m.mu.RLock()
	conn, ok := m.conns[store]
	m.mu.RUnlock()
	if ok {
		return conn, nil
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	conn, ok = m.conns[store]
	if ok {
		return conn, nil
	}

	conn, err := m.getConn(ctx, store)
	if err != nil {
		return nil, err
	}
	m.conns[store] = conn
	return conn, nil
}

func (m *connManager[C]) getConn(ctx context.Context, store string) (Conn, error) {
	config, err := m.getConfiguration(ctx, store)
	if err != nil {
		return nil, err
	}

	conn, err := m.connector(ctx, config)
	if err != nil {
		return nil, err
	}

	if configurable, ok := conn.(Configurable[C]); ok {
		ch := make(chan C)
		err := m.watchConfiguration(context.Background(), store, ch)
		if err != nil {
			return nil, err
		}
		go func() {
			for config := range ch {
				err := configurable.Configure(context.Background(), config)
				if err != nil {
					log.Error(err)
				}
			}
		}()
	}
	return conn, nil
}

func (m *connManager[C]) getConfiguration(ctx context.Context, store string) (C, error) {
	request := &controllerv1.GetConfigurationRequest{
		Store: store,
	}
	response, err := m.runtime.Controller().GetConfiguration(ctx, request)
	if err != nil {
		return nil, errors.From(err)
	}
	config := m.configProvider()
	err = proto.Unmarshal(response.Configuration.Data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (m *connManager[C]) watchConfiguration(ctx context.Context, store string, ch chan<- C) error {
	request := &controllerv1.WatchConfigurationRequest{
		Store: store,
	}
	stream, err := m.runtime.Controller().WatchConfiguration(ctx, request)
	if err != nil {
		return errors.From(err)
	}
	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Error(err)
			} else {
				config := m.configProvider()
				err = proto.Unmarshal(response.Configuration.Data, config)
				if err != nil {
					log.Error(err)
				} else {
					ch <- config
				}
			}
		}
	}()
	return nil
}

var _ ConnManager = (*connManager[proto.Message])(nil)

// Conn is a driver connection
type Conn interface {
	io.Closer
}

type Connector[C proto.Message] func(ctx context.Context, configuration C) (Conn, error)

type ConfigurationProvider[C proto.Message] func() C

type Configurable[C proto.Message] interface {
	Configure(ctx context.Context, configuration C) error
}
