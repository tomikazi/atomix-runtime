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

package main

import (
	_ "github.com/atomix/atomix-runtime/examples/pubsub"
	_ "github.com/atomix/atomix-runtime/examples/storage"
	"github.com/atomix/atomix-runtime/internal/runtime"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Configure and start the runtime
	rt := runtime.New(runtime.GlobalRegistry)
	if err := rt.Run(); err != nil {
		panic(err)
	}

	// Wait for an interrupt signal
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
}
