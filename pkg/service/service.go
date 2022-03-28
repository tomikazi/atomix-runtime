// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package service

// Service is the base interface for all services
type Service interface {
	// Start starts the service
	Start() error
	// Stop stops the service
	Stop() error
}
