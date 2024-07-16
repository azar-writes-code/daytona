// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package dto

import "github.com/daytonaio/daytona/pkg/build"

type BuildDTO struct {
	Hash              string `gorm:"primaryKey"`
	State             string `json:"state"`
	User              string `json:"user"`
	Image             string `json:"image"`
	ProjectVolumePath string `json:"projectVolumePath"`
}

func ToBuildDTO(build *build.Build) BuildDTO {
	return BuildDTO{
		Hash:              build.Hash,
		State:             string(build.State),
		User:              build.User,
		Image:             build.Image,
		ProjectVolumePath: build.ProjectVolumePath,
	}
}

func ToBuild(buildDTO BuildDTO) *build.Build {
	return &build.Build{
		Hash:              buildDTO.Hash,
		State:             build.BuildState(buildDTO.State),
		User:              buildDTO.User,
		Image:             buildDTO.Image,
		ProjectVolumePath: buildDTO.ProjectVolumePath,
	}
}
