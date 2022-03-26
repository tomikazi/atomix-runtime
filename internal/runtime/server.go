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
	"github.com/atomix/atomix-runtime/pkg/service"
	"google.golang.org/grpc"
)

// NewServer creates a new runtime server
func NewServer() *Server {
	return &Server{
		server: grpc.NewServer(),
	}
}

// Server is the runtime server
type Server struct {
	server *grpc.Server
}

func (s *Server) Start() error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Stop() error {
	//TODO implement me
	panic("implement me")
}

var _ service.Service = (*Server)(nil)
