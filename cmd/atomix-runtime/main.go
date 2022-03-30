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
	cmd.Flags().StringP("host", "h", runtime.DefaultHost, "the host to which to bind the runtime server")
	cmd.Flags().IntP("port", "p", runtime.DefaultPort, "the port to which to bind the runtime server")
	cmd.Flags().String("controller-host", runtime.DefaultControllerHost, "the host at which to connect to the controller server")
	cmd.Flags().Int("controller-port", runtime.DefaultControllerPort, "the port at which to connect to the controller server")
	cmd.Flags().String("driver-plugin-dir", runtime.DefaultDriverPluginDir, "the directory from which to load driver plugins")
	cmd.Flags().String("primitive-plugin-dir", runtime.DefaultPrimitivePluginDir, "the directory from which to load primitive plugins")

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
	driverPluginDir, err := cmd.Flags().GetString("driver-plugin-dir")
	if err != nil {
		panic(err)
	}
	primitivePluginDir, err := cmd.Flags().GetString("primitive-plugin-dir")
	if err != nil {
		panic(err)
	}

	r := runtime.NewRuntime(
		runtime.WithHost(host),
		runtime.WithPort(port),
		runtime.WithControllerHost(controllerHost),
		runtime.WithControllerPort(controllerPort),
		runtime.WithDriverPluginDir(driverPluginDir),
		runtime.WithPrimitivePluginDir(primitivePluginDir))
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
