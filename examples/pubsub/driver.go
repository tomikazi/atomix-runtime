// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pubsub

import (
	"fmt"
	"github.com/atomix/atomix-runtime/pkg/pubsub"
	"github.com/atomix/atomix-runtime/pkg/runtime"
)

const Name = "example"
const Version = "v1"

func init() {
	runtime.Register(Driver{})
}

type Driver struct{}

func (d Driver) Name() string {
	return Name
}

func (d Driver) Version() string {
	return Version
}

func (d Driver) NewTopic(config pubsub.TopicConfig) pubsub.Topic {
	//TODO implement me
	panic("implement me")
}

func (d Driver) String() string {
	return fmt.Sprintf("%s/%s", Name, Version)
}

var _ pubsub.Driver = Driver{}
