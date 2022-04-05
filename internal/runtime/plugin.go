// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"fmt"
	"github.com/atomix/atomix-runtime/pkg/errors"
	"github.com/atomix/atomix-runtime/pkg6/runtime"
	"io/fs"
	"path/filepath"
	"plugin"
	"strings"
	"sync"
)

const pluginSymbol = "Plugin"

// NewPluginRegistry creates a new plugin registry of the given type
func NewPluginRegistry[P runtime.Plugin[T], T runtime.Component](path string) PluginRegistry[P, T] {
	return &pluginRegistry[P, T]{
		path:    path,
		plugins: make(map[string]map[string]P),
	}
}

// PluginRegistry is a plugin registry
type PluginRegistry[P runtime.Plugin[T], T runtime.Component] interface {
	Lookup(name, version string) (P, error)
	LoadAll() ([]P, error)
}

type pluginRegistry[P runtime.Plugin[T], T runtime.Component] struct {
	path    string
	plugins map[string]map[string]P
	mu      sync.RWMutex
}

func (r *pluginRegistry[P, T]) Lookup(name, version string) (P, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	versions, ok := r.plugins[name]
	if ok {
		d, ok := versions[version]
		if ok {
			return d, nil
		}
	}

	p, err := plugin.Open(filepath.Join(r.path, fmt.Sprintf("%s.%s.so", name, version)))
	if err != nil {
		return nil, errors.NewNotFound(err.Error())
	}

	symbol, err := p.Lookup(pluginSymbol)
	if err != nil {
		return nil, errors.NewInvalid(err.Error())
	}

	versions, ok = r.plugins[name]
	if !ok {
		versions = make(map[string]P)
		r.plugins[name] = versions
	}
	versions[version] = symbol.(P)
	return symbol.(P), nil
}

func (r *pluginRegistry[P, T]) LoadAll() ([]P, error) {
	var plugins []P
	err := filepath.Walk(r.path, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		parts := strings.Split(info.Name(), ".")
		if len(parts) != 3 {
			return nil
		}
		if parts[2] != "so" {
			return nil
		}
		name, version := parts[0], parts[1]
		driver, err := r.Lookup(name, version)
		if err != nil {
			return err
		}
		plugins = append(plugins, driver)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return plugins, nil
}

var _ PluginRegistry[runtime.Plugin[runtime.Component], runtime.Component] = (*pluginRegistry[runtime.Plugin[runtime.Component], runtime.Component])(nil)
