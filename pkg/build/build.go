// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package build

import "github.com/daytonaio/daytona/pkg/workspace"

type BuildState string

const (
	BuildStatePending BuildState = "pending"
	BuildStateRunning BuildState = "running"
	BuildStateFailure BuildState = "failure"
	BuildStateSuccess BuildState = "success"
)

type Build struct {
	Hash    string            `json:"hash"`
	State   BuildState        `json:"state"`
	Project workspace.Project `json:"project"`
	User    string            `json:"user"`
	Image   string            `json:"image"`
} // @name Build
