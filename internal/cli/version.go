// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"fmt"
	"github.com/atomix/atomix-runtime/pkg/version"
	"github.com/spf13/cobra"
)

func getVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "version",
		RunE: runVersionCommand,
	}
	return cmd
}

func runVersionCommand(cmd *cobra.Command, args []string) error {
	fmt.Println(version.Version())
	return nil
}
