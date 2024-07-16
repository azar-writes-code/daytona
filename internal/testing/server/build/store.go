//go:build testing

// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package build

import t_build "github.com/daytonaio/daytona/pkg/build"

type InMemoryBuildStore struct {
	builds map[string]*t_build.Build
}

func NewInMemoryBuildStore() t_build.Store {
	return &InMemoryBuildStore{
		builds: make(map[string]*t_build.Build),
	}
}

func (s *InMemoryBuildStore) Find(hash string) (*t_build.Build, error) {
	build, ok := s.builds[hash]
	if !ok {
		return nil, t_build.ErrBuildNotFound
	}

	return build, nil
}

func (s *InMemoryBuildStore) FindAllByState(state t_build.BuildState) ([]*t_build.Build, error) {
	builds := []*t_build.Build{}
	for _, build := range s.builds {
		if build.State == state {
			builds = append(builds, build)
		}
	}

	return builds, nil
}

func (s *InMemoryBuildStore) List() ([]*t_build.Build, error) {
	build := []*t_build.Build{}
	for _, a := range s.builds {
		build = append(build, a)
	}

	return build, nil
}

func (s *InMemoryBuildStore) Save(build *t_build.Build) error {
	s.builds[build.Hash] = build
	return nil
}

func (s *InMemoryBuildStore) Delete(hash string) error {
	delete(s.builds, hash)
	return nil
}
