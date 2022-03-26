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

package controller

import "fmt"

const defaultHost = "127.0.0.1"
const defaultPort = 5679

// NewOptions creates new Options from the given set of Option
func NewOptions(opts ...Option) Options {
	var options Options
	options.apply(opts...)
	return options
}

// Options is a set of controller options
type Options struct {
	Host string
	Port int
}

func (o Options) Address() string {
	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}

func (o Options) apply(opts ...Option) {
	o.Host = defaultHost
	o.Port = defaultPort
	for _, opt := range opts {
		opt.apply(&o)
	}
}

// Option is a controller option
type Option interface {
	apply(*Options)
}

// WithHost overrides the default controller host
func WithHost(host string) Option {
	return newFuncOption(func(options *Options) {
		options.Host = host
	})
}

// WithPort overrides the default controller port
func WithPort(port int) Option {
	return newFuncOption(func(options *Options) {
		options.Port = port
	})
}

func newFuncOption(f func(*Options)) Option {
	return &funcOption{f}
}

type funcOption struct {
	f func(*Options)
}

func (o *funcOption) apply(options *Options) {
	o.f(options)
}
