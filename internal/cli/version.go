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
