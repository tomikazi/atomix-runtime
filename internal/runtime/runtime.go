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
	"github.com/atomix/atomix-runtime/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

var log = logging.GetLogger("atomix", "runtime")

func New(opts ...Option) Runtime {
	return &atomixRuntime{
		options: NewOptions(opts...),
	}
}

type Runtime interface {
	Run() error
}

type atomixRuntime struct {
	runtimeServer    *grpc.Server
	controllerClient controllerv1.ControllerClient
	options          Options
}

func (r *atomixRuntime) Run() error {
	controllerConn, err := grpc.Dial(
		r.options.ControllerOptions.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	r.controllerClient = controllerv1.NewControllerClient(controllerConn)
	runtimev1.RegisterRuntimeServer(r.runtimeServer, NewServer(r.controllerClient))
	lis, err := net.Listen("tcp", r.options.ServerOptions.Address())
	if err != nil {
		return err
	}
	go func() {
		if err := r.runtimeServer.Serve(lis); err != nil {
			log.Error(err)
		}
	}()
	return nil
}
