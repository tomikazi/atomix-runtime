// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"github.com/spf13/cobra"
)

func getInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "init",
		RunE: runInitCommand,
	}
	return cmd
}

func runInitCommand(cmd *cobra.Command, args []string) error {
	return nil
}
