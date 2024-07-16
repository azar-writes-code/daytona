// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package build_test

import (
	"io"
	"testing"

	"github.com/daytonaio/daytona/internal/testing/git/mocks"
	t_build "github.com/daytonaio/daytona/internal/testing/server/build"
	"github.com/daytonaio/daytona/pkg/build"
	"github.com/daytonaio/daytona/pkg/git"
	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/workspace"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var project workspace.Project = workspace.Project{
	Repository: &gitprovider.GitRepository{},
	Build: &workspace.ProjectBuild{
		Devcontainer: &workspace.ProjectBuildDevcontainer{
			DevContainerFilePath: ".devcontainer/devcontainer.json",
		},
	},
}

var predefBuild build.Build = build.Build{
	Hash:              "test-predef",
	User:              "test-predef",
	Image:             "test-predef",
	ProjectVolumePath: "test-predef",
}

var expectedBuilds []*build.Build

type BuilderTestSuite struct {
	suite.Suite
	mockGitService *mocks.MockGitService
	builder        build.IBuilder
	buildStore     build.Store
}

func NewBuilderTestSuite() *BuilderTestSuite {
	return &BuilderTestSuite{}
}

func (s *BuilderTestSuite) SetupTest() {
	s.buildStore = t_build.NewInMemoryBuildStore()
	s.mockGitService = mocks.NewMockGitService()
	factory := build.NewBuilderFactory(build.BuilderFactoryConfig{
		BuilderConfig: build.BuilderConfig{
			BuildStore: s.buildStore,
		},
		CreateGitService: func(projectDir string, w io.Writer) git.IGitService {
			return s.mockGitService
		},
	})
	s.mockGitService.On("CloneRepository", mock.Anything, mock.Anything).Return(nil)
	s.builder, _ = factory.Create(project, nil)
	err := s.buildStore.Save(&predefBuild)
	if err != nil {
		panic(err)
	}
}

func TestBuilder(t *testing.T) {
	suite.Run(t, NewBuilderTestSuite())
}

func (s *BuilderTestSuite) TestSaveBuilds() {
	expectedBuilds = append(expectedBuilds, &predefBuild)

	require := s.Require()

	err := s.builder.SaveBuilds(predefBuild)
	require.NoError(err)

	savedBuilds, err := s.buildStore.List()
	require.NoError(err)
	require.ElementsMatch(expectedBuilds, savedBuilds)
}
