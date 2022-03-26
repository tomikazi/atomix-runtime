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
	context "context"
	controllerv1 "github.com/atomix/atomix-runtime/api/atomix/controller/v1"
	"github.com/atomix/atomix-sdk/pkg/logging"
)

func NewServer() controllerv1.ControllerServer {
	return &Server{}
}

type Server struct {
	log logging.Logger
}

func (s *Server) CreatePrimitive(ctx context.Context, request *controllerv1.CreatePrimitiveRequest) (*controllerv1.CreatePrimitiveResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeletePrimitive(ctx context.Context, request *controllerv1.DeletePrimitiveRequest) (*controllerv1.DeletePrimitiveResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ controllerv1.ControllerServer = (*Server)(nil)
