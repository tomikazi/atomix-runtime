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

package controller

import (
	controllerv1 "github.com/atomix/atomix-runtime/api/atomix/controller/v1"
	"github.com/atomix/atomix-runtime/pkg/logging"
	"github.com/atomix/atomix-runtime/pkg/service"
	"google.golang.org/grpc"
	"net"
)

var log = logging.GetLogger("atomix", "controller")

// NewController creates a new controller service
func NewController(opts ...Option) Controller {
	return &atomixController{
		server:  grpc.NewServer(),
		options: NewOptions(opts...),
	}
}

// Controller is a runtime controller
type Controller interface {
	service.Service
}

type atomixController struct {
	server  *grpc.Server
	options Options
}

func (c *atomixController) Start() error {
	controllerv1.RegisterControllerServer(c.server, NewServer())
	lis, err := net.Listen("tcp", c.options.Address())
	if err != nil {
		return err
	}
	go func() {
		if err := c.server.Serve(lis); err != nil {
			log.Error(err)
		}
	}()
	return nil
}

func (c *atomixController) Stop() error {
	c.server.Stop()
	return nil
}

var _ Controller = (*atomixController)(nil)
