// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package projectconfig

import (
	"github.com/daytonaio/daytona/pkg/workspace/project/config"
)

type IProjectConfigService interface {
	Delete(projectConfig *config.ProjectConfig) error
	Find(projectConfigName string) (*config.ProjectConfig, error)
	List() ([]*config.ProjectConfig, error)
	Map() (map[string]*config.ProjectConfig, error)
	Save(projectConfig *config.ProjectConfig) error
}

type ProjectConfigServiceConfig struct {
	ConfigStore config.Store
}

type ProjectConfigService struct {
	configStore config.Store
}

func NewProjectConfigService(config ProjectConfigServiceConfig) IProjectConfigService {
	return &ProjectConfigService{
		configStore: config.ConfigStore,
	}
}

func (s *ProjectConfigService) List() ([]*config.ProjectConfig, error) {
	return s.configStore.List()
}

func (s *ProjectConfigService) Map() (map[string]*config.ProjectConfig, error) {
	list, err := s.configStore.List()
	if err != nil {
		return nil, err
	}

	projectConfigs := make(map[string]*config.ProjectConfig)
	for _, projectConfig := range list {
		projectConfigs[projectConfig.Name] = projectConfig
	}

	return projectConfigs, nil
}

func (s *ProjectConfigService) Find(projectConfigName string) (*config.ProjectConfig, error) {
	return s.configStore.Find(projectConfigName)
}

func (s *ProjectConfigService) Save(projectConfig *config.ProjectConfig) error {
	return s.configStore.Save(projectConfig)
}

func (s *ProjectConfigService) Delete(projectConfig *config.ProjectConfig) error {
	return s.configStore.Delete(projectConfig)
}
