// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package primitive

import (
	"fmt"
	"github.com/atomix/atomix-runtime/pkg/primitive"
	"io/fs"
	"path/filepath"
	"plugin"
	"strings"
	"sync"
)

// Registry is a primitive registry
type Registry interface {
	Lookup(name, version string) (primitive.PrimitiveType, error)
	List() ([]primitive.PrimitiveType, error)
}

func NewPluginRegistry(path string) Registry {
	return &pluginRegistry{
		path:       path,
		primitives: make(map[string]map[string]primitive.PrimitiveType),
	}
}

type pluginRegistry struct {
	path       string
	primitives map[string]map[string]primitive.PrimitiveType
	mu         sync.Mutex
}

func (r *pluginRegistry) Lookup(name, version string) (primitive.PrimitiveType, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	versions, ok := r.primitives[name]
	if ok {
		t, ok := versions[version]
		if ok {
			return t, nil
		}
	}

	p, err := plugin.Open(filepath.Join(r.path, fmt.Sprintf("%s.%s.so", name, version)))
	if err != nil {
		return nil, err
	}

	s, err := p.Lookup(fmt.Sprintf("%s%s", strings.ToTitle(name), strings.ToTitle(version)))
	if err != nil {
		return nil, err
	}

	t := s.(primitive.PrimitiveType)

	versions, ok = r.primitives[name]
	if !ok {
		versions = make(map[string]primitive.PrimitiveType)
		r.primitives[name] = versions
	}
	versions[version] = t
	return t, nil
}

func (r *pluginRegistry) List() ([]primitive.PrimitiveType, error) {
	var primitives []primitive.PrimitiveType
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
		primitive, err := r.Lookup(name, version)
		if err != nil {
			return err
		}
		primitives = append(primitives, primitive)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return primitives, nil
}
