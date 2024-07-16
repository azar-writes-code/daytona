// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package projectconfig

import (
	"context"

	apiclient_util "github.com/daytonaio/daytona/internal/util/apiclient"
	"github.com/daytonaio/daytona/pkg/apiclient"
	workspace_util "github.com/daytonaio/daytona/pkg/cmd/workspace/util"
	"github.com/daytonaio/daytona/pkg/views"
	"github.com/daytonaio/daytona/pkg/views/workspace/create"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var projectConfigAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add project config",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var projects []apiclient.CreateProjectDTO
		var projectConfigName string
		var existingProjectConfigNames []string
		ctx := context.Background()

		apiClient, err := apiclient_util.GetApiClient(nil)
		if err != nil {
			log.Fatal(err)
		}

		gitProviders, res, err := apiClient.GitProviderAPI.ListGitProviders(ctx).Execute()
		if err != nil {
			log.Fatal(apiclient_util.HandleErrorResponse(res, err))
		}

		apiServerConfig, res, err := apiClient.ServerAPI.GetConfig(context.Background()).Execute()
		if err != nil {
			log.Fatal(apiclient_util.HandleErrorResponse(res, err))
		}

		profileData, res, err := apiClient.ProfileAPI.GetProfileData(ctx).Execute()
		if err != nil {
			log.Fatal(apiclient_util.HandleErrorResponse(res, err))
		}

		existingProjectConfigs, res, err := apiClient.ProjectConfigAPI.ListProjectConfigs(context.Background()).Execute()
		if err != nil {
			log.Fatal(apiclient_util.HandleErrorResponse(res, err))
		}
		for _, pc := range existingProjectConfigs {
			existingProjectConfigNames = append(existingProjectConfigNames, *pc.Name)
		}

		projectDefaults := &create.ProjectDefaults{
			BuildChoice:          create.AUTOMATIC,
			Image:                apiServerConfig.DefaultProjectImage,
			ImageUser:            apiServerConfig.DefaultProjectUser,
			DevcontainerFilePath: create.DEVCONTAINER_FILEPATH,
		}

		projects, err = workspace_util.GetProjectsCreationDataFromPrompt(workspace_util.ProjectsDataPromptConfig{
			UserGitProviders: gitProviders,
			Manual:           false,
			MultiProject:     false,
			ApiClient:        apiClient,
			Defaults:         projectDefaults,
		},
		)
		if err != nil {
			log.Fatal(err)
		}

		create.ProjectsConfigurationChanged, err = create.ConfigureProjects(&projects, *projectDefaults)
		if err != nil {
			log.Fatal(err)
		}

		if projects[0].Name == nil {
			log.Fatal("project config name is required")
		}

		err = create.RunSubmissionForm(&projectConfigName, *projects[0].Name, existingProjectConfigNames, &projects, projectDefaults)
		if err != nil {
			log.Fatal(err)
		}

		newProjectConfig := apiclient.CreateProjectConfigDTO{
			Name:  projects[0].Name,
			Build: projects[0].Build,
			Image: projects[0].Image,
			User:  projects[0].User,
		}

		projects[0].EnvVars = workspace_util.GetEnvVariables(&projects[0], profileData)

		res, err = apiClient.ProjectConfigAPI.SetProjectConfig(ctx).ProjectConfig(newProjectConfig).Execute()
		if err != nil {
			log.Fatal(apiclient_util.HandleErrorResponse(res, err))
		}

		views.RenderInfoMessage("Project config set successfully")
	},
}

func init() {
}
