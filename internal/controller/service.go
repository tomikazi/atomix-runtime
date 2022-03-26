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
	"github.com/atomix/atomix-runtime/pkg/service"
	"github.com/atomix/atomix-sdk/pkg/logging"
	"google.golang.org/grpc"
	"net"
)

var log = logging.GetLogger("atomix", "runtime", "controller")

// NewService creates a new controller service
func NewService(opts ...Option) service.Service {
	return &Service{
		server:  grpc.NewServer(),
		options: NewOptions(opts...),
	}
}

// Service is a service managing the lifecycle of the controller server
type Service struct {
	server  *grpc.Server
	options Options
}

func (s *Service) Start() error {
	controllerv1.RegisterControllerServer(s.server, NewServer())
	lis, err := net.Listen("tcp", s.options.Address())
	if err != nil {
		return err
	}
	go func() {
		if err := s.server.Serve(lis); err != nil {
			log.Error(err)
		}
	}()
	return nil
}

func (s *Service) Stop() error {
	s.server.Stop()
	return nil
}

var _ service.Service = (*Service)(nil)
