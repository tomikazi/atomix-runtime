// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package primitive

import (
	"google.golang.org/grpc"
)

// Config is an identifier interface for primitive configuration
type Config interface{}

type PrimitiveType interface {
	RegisterServer(server *grpc.Server)
}
