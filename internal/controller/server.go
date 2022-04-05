// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"context"
	controllerv1 "github.com/atomix/atomix-runtime/api/atomix/controller/v1"
	"github.com/atomix/atomix-runtime/pkg/logging"
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

func (s *Server) GetPrimitive(ctx context.Context, request *controllerv1.GetPrimitiveRequest) (*controllerv1.GetPrimitiveResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeletePrimitive(ctx context.Context, request *controllerv1.DeletePrimitiveRequest) (*controllerv1.DeletePrimitiveResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetStore(ctx context.Context, request *controllerv1.GetStoreRequest) (*controllerv1.GetStoreResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ controllerv1.ControllerServer = (*Server)(nil)
