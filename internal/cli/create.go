// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package cli

import "github.com/spf13/cobra"

func getCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create",
	}
	cmd.AddCommand(getCreateStoreCommand())
	cmd.AddCommand(getCreateBrokerCommand())
	cmd.AddCommand(getCreateTopicCommand())
	cmd.AddCommand(getCreateCounterCommand())
	cmd.AddCommand(getCreateLockCommand())
	cmd.AddCommand(getCreateMapCommand())
	cmd.AddCommand(getCreateSetCommand())
	return cmd
}

func getCreateStoreCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "store",
	}
	return cmd
}

func getCreateBrokerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "broker",
	}
	return cmd
}

func getCreateTopicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "topic",
	}
	return cmd
}

func getCreateCounterCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "counter",
	}
	return cmd
}

func getCreateLockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "lock",
	}
	return cmd
}

func getCreateMapCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "map",
	}
	return cmd
}

func getCreateSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "set",
	}
	return cmd
}
