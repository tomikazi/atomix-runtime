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

package driver

import (
	"github.com/atomix/atomix-runtime/pkg/codec"
	"github.com/atomix/atomix-runtime/pkg/service"
)

// Config is a driver configuration
type Config interface{}

// Driver is a storage driver
type Driver[C Config] interface {
	Codec() codec.Codec[C]
	NewAgent(id AgentID) Agent[C]
}

type AgentID string

// Agent is a storage agent
type Agent[C Config] interface {
	service.Service
	Configure(config C) error
}
