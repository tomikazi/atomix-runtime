// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"fmt"
	"github.com/atomix/atomix-runtime/pkg/logging"
	"github.com/spf13/cobra"
	"os"
)

func GetRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "atomix",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			level, _ := cmd.PersistentFlags().GetString("log-level")
			switch level {
			case "debug":
				logging.SetLevel(logging.DebugLevel)
			case "info":
				logging.SetLevel(logging.InfoLevel)
			case "warn":
				logging.SetLevel(logging.WarnLevel)
			case "error":
				logging.SetLevel(logging.ErrorLevel)
			case "fatal":
				logging.SetLevel(logging.FatalLevel)
			case "panic":
				logging.SetLevel(logging.PanicLevel)
			default:
				fmt.Println("unknown log level", level)
				os.Exit(1)
			}
		},
	}
	cmd.PersistentFlags().StringP("log-level", "l", "info", "the log level")
	cmd.AddCommand(getInitCommand())
	cmd.AddCommand(getRunCommand())
	cmd.AddCommand(getCreateCommand())
	cmd.AddCommand(getDeleteCommand())
	cmd.AddCommand(getVersionCommand())
	return cmd
}
