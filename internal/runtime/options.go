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
	"fmt"
	"github.com/atomix/atomix-runtime/internal/controller"
)

const DefaultHost = "127.0.0.1"
const DefaultPort = 5678

// NewOptions creates new Options from the given set of Option
func NewOptions(opts ...Option) Options {
	var options Options
	options.apply(opts...)
	return options
}

type Options struct {
	ServerOptions
	ControllerOptions
}

func (o Options) apply(opts ...Option) {
	var serverOpts []ServerOption
	var controllerOpts []ControllerOption
	for _, opt := range opts {
		if serviceOpt, ok := opt.(ServerOption); ok {
			serverOpts = append(serverOpts, serviceOpt)
		} else if controllerOpt, ok := opt.(ControllerOption); ok {
			controllerOpts = append(controllerOpts, controllerOpt)
		}
	}
	o.ServerOptions.apply(serverOpts...)
	o.ControllerOptions.apply(controllerOpts...)
}

type Option interface {
	ServerOption | ControllerOption
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
	o.Host = DefaultHost
	o.Port = DefaultPort
	for _, opt := range opts {
		opt.applyServiceOptions(&o)
	}
}

// ServerOption is a service option
type ServerOption interface {
	applyServiceOptions(*ServerOptions)
}

// WithHost overrides the default service host
func WithHost(host string) ServerOption {
	return newFuncOption(func(options *ServerOptions) {
		options.Host = host
	})
}

// WithPort overrides the default service port
func WithPort(port int) ServerOption {
	return newFuncOption(func(options *ServerOptions) {
		options.Port = port
	})
}

func newFuncOption(f func(*ServerOptions)) ServerOption {
	return &funcServiceOption{f}
}

type funcServiceOption struct {
	f func(*ServerOptions)
}

func (o *funcServiceOption) applyServiceOptions(options *ServerOptions) {
	o.f(options)
}

// NewControllerOptions creates new ControllerOptions from the given set of ControllerOption
func NewControllerOptions(opts ...ControllerOption) ControllerOptions {
	var options ControllerOptions
	options.apply(opts...)
	return options
}

// ControllerOptions is a set of service options
type ControllerOptions struct {
	Host string
	Port int
}

func (o ControllerOptions) Address() string {
	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}

func (o ControllerOptions) apply(opts ...ControllerOption) {
	o.Host = controller.DefaultHost
	o.Port = controller.DefaultPort
	for _, opt := range opts {
		opt.applyControllerOptions(&o)
	}
}

// ControllerOption is a service option
type ControllerOption interface {
	applyControllerOptions(*ControllerOptions)
}

// WithControllerHost overrides the default service host
func WithControllerHost(host string) ControllerOption {
	return newFuncControllerOption(func(options *ControllerOptions) {
		options.Host = host
	})
}

// WithControllerPort overrides the default service port
func WithControllerPort(port int) ControllerOption {
	return newFuncControllerOption(func(options *ControllerOptions) {
		options.Port = port
	})
}

func newFuncControllerOption(f func(*ControllerOptions)) ControllerOption {
	return &funcControllerOption{f}
}

type funcControllerOption struct {
	f func(*ControllerOptions)
}

func (o *funcControllerOption) applyControllerOptions(options *ControllerOptions) {
	o.f(options)
}