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
	cmd.AddCommand(getBuildCommand())
	cmd.AddCommand(getRunCommand())
	cmd.AddCommand(getCreateCommand())
	cmd.AddCommand(getDeleteCommand())
	cmd.AddCommand(getVersionCommand())
	return cmd
}
