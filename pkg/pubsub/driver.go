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

package pubsub

import (
	"github.com/atomix/atomix-runtime/pkg/component"
	"github.com/atomix/atomix-runtime/pkg/service"
)

// Driver is an interface for implementing publish-subscribe drivers
type Driver interface {
	component.Component
	NewTopic(config TopicConfig) Topic
}

// TopicConfig is a topic configuration
type TopicConfig struct {
}

// Topic is a publish-subscribe topic
type Topic interface {
	service.Service
}
