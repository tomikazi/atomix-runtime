// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package version

var (
	version string
	commit  string
)

func Version() string {
	return version
}

func CommitHash() string {
	return commit
}
