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
