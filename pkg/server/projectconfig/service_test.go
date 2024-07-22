// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package projectconfig_test

import (
	"testing"

	projectconfig_internal "github.com/daytonaio/daytona/internal/testing/server/projectconfig"
	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/server/projectconfig"
	"github.com/daytonaio/daytona/pkg/workspace/project/config"
	"github.com/stretchr/testify/suite"
)

var projectConfig1 *config.ProjectConfig = &config.ProjectConfig{
	Name:  "pc1",
	Image: "image1",
	User:  "user1",
	Build: nil,
	Repository: &gitprovider.GitRepository{
		Url: "url1",
	},
}

var projectConfig2 *config.ProjectConfig = &config.ProjectConfig{
	Name:  "pc2",
	Image: "image2",
	User:  "user2",
	Build: nil,
	Repository: &gitprovider.GitRepository{
		Url: "url2",
	},
}

var projectConfig3 *config.ProjectConfig = &config.ProjectConfig{
	Name:  "pc3",
	Image: "image3",
	User:  "user3",
	Build: nil,
	Repository: &gitprovider.GitRepository{
		Url: "url3",
	},
}

var projectConfig4 *config.ProjectConfig = &config.ProjectConfig{
	Name:  "pc4",
	Image: "image4",
	User:  "user4",
	Build: nil,
	Repository: &gitprovider.GitRepository{
		Url: "url4",
	},
}

var expectedProjectConfigs []*config.ProjectConfig
var expectedProjectConfigsMap map[string]*config.ProjectConfig

type ProjectConfigServiceTestSuite struct {
	suite.Suite
	projectConfigService projectconfig.IProjectConfigService
	projectConfigStore   config.Store
}

func NewProjectConfigServiceTestSuite() *ProjectConfigServiceTestSuite {
	return &ProjectConfigServiceTestSuite{}
}

func (s *ProjectConfigServiceTestSuite) SetupTest() {
	expectedProjectConfigs = []*config.ProjectConfig{
		projectConfig1, projectConfig2, projectConfig3,
	}

	expectedProjectConfigsMap = map[string]*config.ProjectConfig{
		projectConfig1.Name: projectConfig1,
		projectConfig2.Name: projectConfig2,
		projectConfig3.Name: projectConfig3,
	}

	s.projectConfigStore = projectconfig_internal.NewInMemoryProjectConfigStore()
	s.projectConfigService = projectconfig.NewProjectConfigService(projectconfig.ProjectConfigServiceConfig{
		ConfigStore: s.projectConfigStore,
	})

	for _, pc := range expectedProjectConfigs {
		_ = s.projectConfigService.Save(pc)
	}
}

func TestProjectConfigService(t *testing.T) {
	suite.Run(t, NewProjectConfigServiceTestSuite())
}

func (s *ProjectConfigServiceTestSuite) TestList() {
	require := s.Require()

	projectConfigs, err := s.projectConfigService.List()
	require.Nil(err)
	require.ElementsMatch(expectedProjectConfigs, projectConfigs)
}

func (s *ProjectConfigServiceTestSuite) TestMap() {
	require := s.Require()

	projectConfigsMap, err := s.projectConfigService.Map()
	require.Nil(err)
	require.Equal(expectedProjectConfigsMap, projectConfigsMap)
}

func (s *ProjectConfigServiceTestSuite) TestFind() {
	require := s.Require()

	projectConfig, err := s.projectConfigService.Find(projectConfig1.Name)
	require.Nil(err)
	require.Equal(projectConfig1, projectConfig)
}

func (s *ProjectConfigServiceTestSuite) TestSave() {
	expectedProjectConfigs = append(expectedProjectConfigs, projectConfig4)

	require := s.Require()

	err := s.projectConfigService.Save(projectConfig4)
	require.Nil(err)

	projectConfigs, err := s.projectConfigService.List()
	require.Nil(err)
	require.ElementsMatch(expectedProjectConfigs, projectConfigs)
}

func (s *ProjectConfigServiceTestSuite) TestDelete() {
	expectedProjectConfigs = expectedProjectConfigs[:2]

	require := s.Require()

	err := s.projectConfigService.Delete(projectConfig3)
	require.Nil(err)

	projectConfigs, err := s.projectConfigService.List()
	require.Nil(err)
	require.ElementsMatch(expectedProjectConfigs, projectConfigs)
}
