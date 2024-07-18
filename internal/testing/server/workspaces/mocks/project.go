//go:build testing

// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/workspace"
)

var MockProject = workspace.Project{
	Build: &workspace.ProjectBuild{
		Devcontainer: &workspace.ProjectBuildDevcontainer{
			DevContainerFilePath: ".devcontainer/devcontainer.json",
		},
	},
	Repository: &gitprovider.GitRepository{
		Url: "url",
	},
}
