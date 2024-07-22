// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package projectconfig

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/server"
	"github.com/daytonaio/daytona/pkg/server/projectconfig/dto"
	"github.com/daytonaio/daytona/pkg/workspace/project/config"
	"github.com/gin-gonic/gin"
)

// GetProjectConfig godoc
//
//	@Tags			project-config
//	@Summary		Get project config data
//	@Description	Get project config data
//	@Accept			json
//	@Param			configName	path		string	true	"Config name"
//	@Success		200			{object}	ProjectConfig
//	@Router			/project-config/{configName} [get]
//
//	@id				GetProjectConfig
func GetProjectConfig(ctx *gin.Context) {
	configName := ctx.Param("configName")

	server := server.GetInstance(nil)

	projectConfig, err := server.ProjectConfigService.Find(configName)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get project config: %s", err.Error()))
		return
	}

	ctx.JSON(200, projectConfig)
}

// ListProjectConfigs godoc
//
//	@Tags			project-config
//	@Summary		List project configs
//	@Description	List project configs
//	@Produce		json
//	@Success		200	{array}	ProjectConfig
//	@Router			/project-config [get]
//
//	@id				ListProjectConfigs
func ListProjectConfigs(ctx *gin.Context) {
	server := server.GetInstance(nil)

	projectConfigs, err := server.ProjectConfigService.List()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to list project configs: %s", err.Error()))
		return
	}

	ctx.JSON(200, projectConfigs)
}

// SetProjectConfig godoc
//
//	@Tags			project-config
//	@Summary		Set project config data
//	@Description	Set project config data
//	@Accept			json
//	@Param			projectConfig	body	dto.CreateProjectConfigDTO	true	"Project config"
//	@Success		201
//	@Router			/project-config [put]
//
//	@id				SetProjectConfig
func SetProjectConfig(ctx *gin.Context) {
	var req dto.CreateProjectConfigDTO
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	projectConfig := config.ProjectConfig{
		Name:       req.Name,
		Build:      req.Build,
		Repository: req.Source.Repository,
		EnvVars:    req.EnvVars,
	}

	if req.Image != nil {
		projectConfig.Image = *req.Image
	}

	if req.User != nil {
		projectConfig.User = *req.User
	}

	fmt.Println(projectConfig)

	s := server.GetInstance(nil)
	err = s.ProjectConfigService.Save(&projectConfig)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to save project config: %s", err.Error()))
		return
	}

	ctx.Status(201)
}

// DeleteProjectConfig godoc
//
//	@Tags			project-config
//	@Summary		Delete project config data
//	@Description	Delete project config data
//	@Param			configName	path	string	true	"Config name"
//	@Success		204
//	@Router			/project-config/{configName} [delete]
//
//	@id				DeleteProjectConfig
func DeleteProjectConfig(ctx *gin.Context) {
	configName := ctx.Param("configName")

	server := server.GetInstance(nil)

	projectConfig, err := server.ProjectConfigService.Find(configName)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to find project config: %s", err.Error()))
		return
	}

	err = server.ProjectConfigService.Delete(projectConfig)
	if err != nil {
		if config.IsProjectConfigNotFound(err) {
			ctx.Status(204)
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get project config: %s", err.Error()))
		return
	}

	ctx.Status(204)
}
