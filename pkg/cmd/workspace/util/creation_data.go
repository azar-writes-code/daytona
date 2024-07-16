// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/daytonaio/daytona/pkg/apiclient"
	"github.com/daytonaio/daytona/pkg/views/workspace/create"
	"github.com/daytonaio/daytona/pkg/views/workspace/selection"
	log "github.com/sirupsen/logrus"
)

type ProjectsDataPromptConfig struct {
	UserGitProviders []apiclient.GitProvider
	ProjectConfigs   []apiclient.ProjectConfig
	Manual           bool
	MultiProject     bool
	ApiClient        *apiclient.APIClient
	Defaults         *create.ProjectDefaults
}

func GetProjectsCreationDataFromPrompt(config ProjectsDataPromptConfig) ([]apiclient.CreateProjectDTO, error) {
	var projectList []apiclient.CreateProjectDTO
	var err error
	// The following three maps keep track of visited repos and their respective namespaces and Git providers
	// During workspace creation, lookups will help in avoiding duplicated repo entries and disabling specific
	// namespaces and git providers list options from further selection under which all repos are visited.
	// The hierarchy is :
	// Git Provider ----> Namespaces (organizations/projects) ----> Repositories
	// Git Provider ----> Repositories (if no orgs available)
	selectedRepos := make(map[string]bool)
	disabledNamespaces := make(map[string]bool)
	disabledGitProviders := make(map[string]bool)

	addMore := config.MultiProject

	for i := 1; addMore || i == 1; i++ {

		if len(config.ProjectConfigs) > 0 {
			projectConfig := selection.GetProjectConfigFromPrompt(config.ProjectConfigs, "Use")
			if projectConfig == nil {
				return nil, fmt.Errorf("must select a project config")
			}

			if *projectConfig.Name != selection.BlankProjectIdentifier {
				projectList = append(projectList, apiclient.CreateProjectDTO{
					ExistingProjectConfigName: projectConfig.Name,
				})
				continue
			}
		}

		var providerRepo *apiclient.GitRepository

		providerRepo, err = getRepositoryFromWizard(RepositoryWizardConfig{
			ApiClient:            config.ApiClient,
			UserGitProviders:     config.UserGitProviders,
			MultiProject:         config.MultiProject,
			ProjectOrder:         i,
			DisabledGitProviders: disabledGitProviders,
			DisabledNamespaces:   disabledNamespaces,
			SelectedRepos:        selectedRepos,
		})
		if err != nil {
			return nil, err
		}

		if i > 1 {
			addMore, err = create.RunAddMoreProjectsForm()
			if err != nil {
				return nil, err
			}
		}

		providerRepoName, err := GetSanitizedProjectName(*providerRepo.Name)
		if err != nil {
			return nil, err
		}

		projectList = append(projectList, newCreateProjectDTO(config, providerRepo, providerRepoName))
	}

	return projectList, nil
}

func newCreateProjectDTO(config ProjectsDataPromptConfig, providerRepo *apiclient.GitRepository, providerRepoName string) apiclient.CreateProjectDTO {
	project := apiclient.CreateProjectDTO{
		Name: &providerRepoName,
		Source: &apiclient.CreateProjectConfigSourceDTO{
			Repository: providerRepo,
		},
		Build:   &apiclient.ProjectBuild{},
		Image:   config.Defaults.Image,
		User:    config.Defaults.ImageUser,
		EnvVars: &map[string]string{},
	}

	return project
}

func GetProjectNameFromRepo(repoUrl string) string {
	projectNameSlugRegex := regexp.MustCompile(`[^a-zA-Z0-9-]`)
	return projectNameSlugRegex.ReplaceAllString(strings.TrimSuffix(strings.ToLower(filepath.Base(repoUrl)), ".git"), "-")
}

func GetSuggestedWorkspaceName(firstProjectName string, existingWorkspaceNames []string) string {
	suggestion := firstProjectName

	if !slices.Contains(existingWorkspaceNames, suggestion) {
		return suggestion
	} else {
		i := 2
		for {
			newSuggestion := fmt.Sprintf("%s%d", suggestion, i)
			if !slices.Contains(existingWorkspaceNames, newSuggestion) {
				return newSuggestion
			}
			i++
		}
	}
}

func GetSanitizedProjectName(projectName string) (string, error) {
	projectName, err := url.QueryUnescape(projectName)
	if err != nil {
		return "", err
	}
	projectName = strings.ReplaceAll(projectName, " ", "-")

	return projectName, nil
}

func GetEnvVariables(project *apiclient.CreateProjectDTO, profileData *apiclient.ProfileData) *map[string]string {
	envVars := map[string]string{}

	if profileData.EnvVars != nil {
		for k, v := range *profileData.EnvVars {
			if strings.HasPrefix(v, "$") {
				env, ok := os.LookupEnv(v[1:])
				if ok {
					envVars[k] = env
				} else {
					log.Warnf("Environment variable %s not found", v[1:])
				}
			} else {
				envVars[k] = v
			}
		}
	}

	if project.EnvVars != nil {
		for k, v := range *project.EnvVars {
			if strings.HasPrefix(v, "$") {
				env, ok := os.LookupEnv(v[1:])
				if ok {
					envVars[k] = env
				} else {
					log.Warnf("Environment variable %s not found", v[1:])
				}
			} else {
				envVars[k] = v
			}
		}
	}

	return &envVars
}
