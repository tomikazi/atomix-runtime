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

package runtime

import (
	controllerv1 "github.com/atomix/atomix-runtime/api/atomix/controller/v1"
	runtimev1 "github.com/atomix/atomix-runtime/api/atomix/runtime/v1"
	"github.com/atomix/atomix-runtime/internal/driver"
	"github.com/atomix/atomix-runtime/internal/primitive"
	"github.com/atomix/atomix-runtime/pkg/logging"
	"github.com/atomix/atomix-runtime/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

var log = logging.GetLogger("atomix", "runtime")

func New(opts ...Option) Runtime {
	return &atomixRuntime{
		server:  grpc.NewServer(),
		options: NewOptions(opts...),
	}
}

type Runtime interface {
	service.Service
}

type atomixRuntime struct {
	server     *grpc.Server
	controller controllerv1.ControllerClient
	drivers    driver.Registry
	primitives primitive.Registry
	options    Options
}

func (r *atomixRuntime) Start() error {
	controllerConn, err := grpc.Dial(
		r.options.ControllerOptions.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	r.controller = controllerv1.NewControllerClient(controllerConn)

	runtimev1.RegisterRuntimeServer(r.server, NewServer(r.controller))
	primtiives, err := r.primitives.List()
	if err != nil {
		return err
	}
	for _, primitive := range primtiives {
		primitive.RegisterServer(r.server)
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

func (r *atomixRuntime) Stop() error {
	r.server.Stop()
	return nil
}

var _ Runtime = (*atomixRuntime)(nil)
