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
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cmd := &cobra.Command{
		Use: "atomix-runtime",
	}
	cmd.Flags().StringP("host", "h", "127.0.0.1", "the host to which to bind the runtime server")
	cmd.Flags().IntP("port", "p", 5678, "the port to which to bind the runtime server")
	cmd.Flags().String("controller-host", "127.0.0.1", "the host at which to connect to the controller server")
	cmd.Flags().Int("controller-port", 5680, "the port at which to connect to the controller server")

	if err := cmd.Execute(); err != nil {
		panic(err)
	}

	host, err := cmd.Flags().GetString("host")
	if err != nil {
		panic(err)
	}
	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		panic(err)
	}
	controllerHost, err := cmd.Flags().GetString("controller-host")
	if err != nil {
		panic(err)
	}
	controllerPort, err := cmd.Flags().GetInt("controller-port")
	if err != nil {
		panic(err)
	}

	r := runtime.New(
		runtime.WithHost(host),
		runtime.WithPort(port),
		runtime.WithControllerHost(controllerHost),
		runtime.WithControllerPort(controllerPort))
	if err := r.Start(); err != nil {
		panic(err)
	}

	// Wait for an interrupt signal
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	if err := r.Stop(); err != nil {
		panic(err)
	}
}
