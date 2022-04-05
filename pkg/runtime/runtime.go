// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	controllerv1 "github.com/atomix/atomix-runtime/api/atomix/controller/v1"
	"github.com/atomix/atomix-runtime/pkg/driver"
	"google.golang.org/grpc"
)

type Runtime interface {
	driver.Provider
	Server() *grpc.Server
	Controller() controllerv1.ControllerClient
}

type Component interface {
	Name() string
	Version() string
}

type ComponentFactory[T Component] func(runtime Runtime) T

type Plugin[T Component] interface {
	New(runtime Runtime) T
}
