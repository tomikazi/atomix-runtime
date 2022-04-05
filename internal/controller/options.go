// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import "fmt"

const DefaultHost = "127.0.0.1"
const DefaultPort = 5680

// NewOptions creates new Options from the given set of Option
func NewOptions(opts ...Option) Options {
	var options Options
	options.apply(opts...)
	return options
}

type Options struct {
	ServerOptions
}

func (o Options) apply(opts ...Option) {
	var serverOpts []ServerOption
	for _, opt := range opts {
		if serviceOpt, ok := opt.(ServerOption); ok {
			serverOpts = append(serverOpts, serviceOpt)
		}
	}
	o.ServerOptions.apply(serverOpts...)
}

type Option interface {
	ServerOption
}

// ServerOptions is a set of service options
type ServerOptions struct {
	Host string
	Port int
}

func (o ServerOptions) Address() string {
	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}

func (o ServerOptions) apply(opts ...ServerOption) {
	o.Port = DefaultPort
	for _, opt := range opts {
		opt.applyServerOption(&o)
	}
}

// ServerOption is a service option
type ServerOption interface {
	applyServerOption(*ServerOptions)
}

// WithHost overrides the default service host
func WithHost(host string) ServerOption {
	return newFuncServerOption(func(options *ServerOptions) {
		options.Host = host
	})
}

// WithPort overrides the default service port
func WithPort(port int) ServerOption {
	return newFuncServerOption(func(options *ServerOptions) {
		options.Port = port
	})
}

func newFuncServerOption(f func(*ServerOptions)) ServerOption {
	return &funcServerOption{f}
}

type funcServerOption struct {
	f func(*ServerOptions)
}

func (o *funcServerOption) applyServerOption(options *ServerOptions) {
	o.f(options)
}
