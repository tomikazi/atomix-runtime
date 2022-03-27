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
	context "context"
	controllerv1 "github.com/atomix/atomix-runtime/api/atomix/controller/v1"
	runtimev1 "github.com/atomix/atomix-runtime/api/atomix/runtime/v1"
)

func NewServer(controllerClient controllerv1.ControllerClient) runtimev1.RuntimeServer {
	return &Server{
		controllerClient: controllerClient,
	}
}

type Server struct {
	controllerClient controllerv1.ControllerClient
}

func (s *Server) CreatePrimitive(ctx context.Context, request *runtimev1.CreatePrimitiveRequest) (*runtimev1.CreatePrimitiveResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeletePrimitive(ctx context.Context, request *runtimev1.DeletePrimitiveRequest) (*runtimev1.DeletePrimitiveResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) CreateProxy(ctx context.Context, request *runtimev1.CreateProxyRequest) (*runtimev1.CreateProxyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) CloseProxy(ctx context.Context, request *runtimev1.CloseProxyRequest) (*runtimev1.CloseProxyResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ runtimev1.RuntimeServer = (*Server)(nil)
