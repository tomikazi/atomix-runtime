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
