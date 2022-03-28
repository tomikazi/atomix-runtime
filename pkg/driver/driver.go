// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"github.com/atomix/atomix-runtime/pkg/service"
)

// Config is a driver configuration
type Config interface{}

// Driver is a storage driver
type Driver[C Config] interface {
	NewConfig() C
	NewAgent() Agent[C]
}

// Agent is a storage agent
type Agent[C Config] interface {
	service.Service
	Configure(config C) error
}
