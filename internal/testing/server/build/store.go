//go:build testing

// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package build

import "github.com/daytonaio/daytona/pkg/build"

type InMemoryBuildStore struct {
	builds map[string]*build.Build
}

func NewInMemoryBuildStore() build.Store {
	return &InMemoryBuildStore{
		builds: make(map[string]*build.Build),
	}
}

func (s *InMemoryBuildStore) Find(hash string) (*build.Build, error) {
	result, ok := s.builds[hash]
	if !ok {
		return nil, build.ErrBuildNotFound
	}

	return result, nil
}

func (s *InMemoryBuildStore) FindAllByState(state build.BuildState) ([]*build.Build, error) {
	builds := []*build.Build{}
	for _, build := range s.builds {
		if build.State == state {
			builds = append(builds, build)
		}
	}

	return builds, nil
}

func (s *InMemoryBuildStore) List() ([]*build.Build, error) {
	builds := []*build.Build{}
	for _, a := range s.builds {
		builds = append(builds, a)
	}

	return builds, nil
}

func (s *InMemoryBuildStore) Save(result *build.Build) error {
	s.builds[result.Hash] = result
	return nil
}

func (s *InMemoryBuildStore) Delete(hash string) error {
	delete(s.builds, hash)
	return nil
}
