// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package cli

import "github.com/spf13/cobra"

func getRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "run",
	}
	return cmd
}
