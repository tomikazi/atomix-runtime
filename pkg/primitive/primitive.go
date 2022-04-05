// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package primitive

import (
	"github.com/atomix/atomix-runtime/pkg/runtime"
)

type Primitive interface {
	runtime.Component
}
