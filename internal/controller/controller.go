// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
