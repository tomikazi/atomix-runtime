// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
	cmd.Flags().IntP("port", "p", controller.DefaultPort, "the port to which to bind the controller server")

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
