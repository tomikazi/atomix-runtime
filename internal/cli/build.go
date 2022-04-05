// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package cli

import "github.com/spf13/cobra"

func getBuildCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "build",
	}
	cmd.AddCommand(getBuildControllerCommand())
	cmd.AddCommand(getBuildRuntimeCommand())
	return cmd
}

func getBuildControllerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "controller",
	}
	return cmd
}

func getBuildRuntimeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "runtime",
	}
	cmd.Flags().StringP("config", "c", "", "the configuration file to use")
	cmd.Flags().StringToStringP("primitive", "p", map[string]string{}, "a mapping of additional primitive plugins to build")
	cmd.Flags().StringToStringP("driver", "d", map[string]string{}, "a mapping of additional driver plugins to build")
	return cmd
}
