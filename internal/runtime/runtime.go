// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"context"
	controllerv1 "github.com/atomix/atomix-runtime/api/atomix/controller/v1"
	"github.com/atomix/atomix-runtime/pkg/driver"
	"github.com/atomix/atomix-runtime/pkg/errors"
	"github.com/atomix/atomix-runtime/pkg/grpc/retry"
	"github.com/atomix/atomix-runtime/pkg/logging"
	"github.com/atomix/atomix-runtime/pkg/primitive"
	"github.com/atomix/atomix-runtime/pkg/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"sync"
)

var log = logging.GetLogger("atomix", "runtime")

func NewRuntime(opts ...Option) *Runtime {
	options := NewOptions(opts...)
	return &Runtime{
		options:           options,
		driverRegistry:    NewPluginRegistry[driver.Plugin, driver.Driver](options.DriverPluginDir),
		primitiveRegistry: NewPluginRegistry[primitive.Plugin, primitive.Primitive](options.PrimitivePluginDir),
	}
}

type Runtime struct {
	options           Options
	controllerConn    *grpc.ClientConn
	server            *grpc.Server
	driverRegistry    PluginRegistry[driver.Plugin, driver.Driver]
	primitiveRegistry PluginRegistry[primitive.Plugin, primitive.Primitive]
	drivers           map[string]driver.Driver
	mu                sync.RWMutex
}

func (r *Runtime) Server() *grpc.Server {
	return r.server
}

func (r *Runtime) Controller() controllerv1.ControllerClient {
	return controllerv1.NewControllerClient(r.controllerConn)
}

func (r *Runtime) GetDriver(ctx context.Context, store string) (driver.Driver, error) {
	r.mu.RLock()
	driver, ok := r.drivers[store]
	r.mu.RUnlock()
	if ok {
		return driver, nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	driver, ok = r.drivers[store]
	if ok {
		return driver, nil
	}

	controllerClient := controllerv1.NewControllerClient(r.controllerConn)
	request := &controllerv1.GetConfigurationRequest{
		Store: store,
	}
	response, err := controllerClient.GetConfiguration(ctx, request)
	if err != nil {
		return nil, errors.From(err)
	}

	driverName, driverVersion := response.Configuration.Type, response.Configuration.Version
	driverPlugin, err := r.driverRegistry.Lookup(driverName, driverVersion)
	if err != nil {
		return nil, err
	}
	driver = driverPlugin.New(r)
	r.drivers[store] = driver
	return driver, nil
}

func (r *Runtime) Start() error {
	controllerConn, err := grpc.Dial(
		r.options.ControllerOptions.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor()),
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor()))
	if err != nil {
		return err
	}
	r.controllerConn = controllerConn

	r.server = grpc.NewServer()
	primitives, err := r.primitiveRegistry.LoadAll()
	if err != nil {
		return err
	}
	for _, primitive := range primitives {
		primitive.New(r)
	}

	lis, err := net.Listen("tcp", r.options.ServerOptions.Address())
	if err != nil {
		return err
	}
	go func() {
		if err := r.server.Serve(lis); err != nil {
			log.Error(err)
		}
	}()
	return nil
}

func (r *Runtime) Stop() error {
	r.server.Stop()
	return r.controllerConn.Close()
}

var _ runtime.Runtime = (*Runtime)(nil)
