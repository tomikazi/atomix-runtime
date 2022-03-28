// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package cli

import "github.com/spf13/cobra"

func getDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete",
	}
	cmd.AddCommand(getDeleteStoreCommand())
	cmd.AddCommand(getDeleteBrokerCommand())
	cmd.AddCommand(getDeleteTopicCommand())
	cmd.AddCommand(getDeleteCounterCommand())
	cmd.AddCommand(getDeleteLockCommand())
	cmd.AddCommand(getDeleteMapCommand())
	cmd.AddCommand(getDeleteSetCommand())
	return cmd
}

func getDeleteStoreCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "store",
	}
	return cmd
}

func getDeleteBrokerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "broker",
	}
	return cmd
}

func getDeleteTopicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "topic",
	}
	return cmd
}

func getDeleteCounterCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "counter",
	}
	return cmd
}

func getDeleteLockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "lock",
	}
	return cmd
}

func getDeleteMapCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "map",
	}
	return cmd
}

func getDeleteSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "set",
	}
	return cmd
}
