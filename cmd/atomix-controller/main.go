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
	"github.com/atomix/atomix-runtime/internal/controller"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cmd := &cobra.Command{
		Use: "atomix-controller",
	}
	cmd.Flags().StringP("host", "h", "", "the host to which to bind the controller server")
	cmd.Flags().IntP("port", "p", 5680, "the port to which to bind the controller server")

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

	c := controller.NewController(
		controller.WithHost(host),
		controller.WithPort(port))
	if err := c.Start(); err != nil {
		panic(err)
	}

	// Wait for an interrupt signal
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	if err := c.Stop(); err != nil {
		panic(err)
	}
}
