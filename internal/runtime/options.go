// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"fmt"
	"github.com/atomix/atomix-runtime/internal/controller"
)

const (
	DefaultHost               = "127.0.0.1"
	DefaultPort               = 5678
	DefaultControllerHost     = controller.DefaultHost
	DefaultControllerPort     = controller.DefaultPort
	DefaultDriverPluginDir    = "./plugins/driver"
	DefaultPrimitivePluginDir = "./plugins/primitive"
)

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
	Host               string
	Port               int
	DriverPluginDir    string
	PrimitivePluginDir string
}

func (o ServerOptions) Address() string {
	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}

func (o ServerOptions) apply(opts ...ServerOption) {
	o.Host = DefaultHost
	o.Port = DefaultPort
	o.DriverPluginDir = DefaultDriverPluginDir
	o.PrimitivePluginDir = DefaultPrimitivePluginDir
	for _, opt := range opts {
		opt.applyServerOptions(&o)
	}
}

// ServerOption is a service option
type ServerOption interface {
	applyServerOptions(*ServerOptions)
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

// WithDriverPluginDir overrides the default driver plugin directory
func WithDriverPluginDir(driverPluginDir string) ServerOption {
	return newFuncOption(func(options *ServerOptions) {
		options.DriverPluginDir = driverPluginDir
	})
}

// WithPrimitivePluginDir overrides the default primitive plugin directory
func WithPrimitivePluginDir(primitivePluginDir string) ServerOption {
	return newFuncOption(func(options *ServerOptions) {
		options.PrimitivePluginDir = primitivePluginDir
	})
}

func newFuncOption(f func(*ServerOptions)) ServerOption {
	return &funcServiceOption{f}
}

type funcServiceOption struct {
	f func(*ServerOptions)
}

func (o *funcServiceOption) applyServerOptions(options *ServerOptions) {
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
	o.Host = DefaultControllerHost
	o.Port = DefaultControllerPort
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
