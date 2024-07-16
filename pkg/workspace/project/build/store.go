// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package build

import "errors"

type Store interface {
	List() ([]*ProjectBuild, error)
	Find(name string) (*ProjectBuild, error)
	Save(projectConfig *ProjectBuild) error
	Delete(projectConfig *ProjectBuild) error
}

var (
	ErrProjectConfigNotFound = errors.New("project config not found")
)

func IsProjectConfigNotFound(err error) bool {
	return err.Error() == ErrProjectConfigNotFound.Error()
}
