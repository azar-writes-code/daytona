// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package workspaces

import (
	"context"
	"fmt"
	"io"

	"github.com/daytonaio/daytona/pkg/logs"
	"github.com/daytonaio/daytona/pkg/provider"
	"github.com/daytonaio/daytona/pkg/telemetry"
	"github.com/daytonaio/daytona/pkg/workspace"
	log "github.com/sirupsen/logrus"

	"github.com/daytonaio/daytona/internal/util"
)

func (s *WorkspaceService) StartWorkspace(ctx context.Context, workspaceId string) error {
	w, err := s.workspaceStore.Find(workspaceId)
	if err != nil {
		return ErrWorkspaceNotFound
	}

	target, err := s.targetStore.Find(w.Target)
	if err != nil {
		return err
	}

	workspaceLogger := s.loggerFactory.CreateWorkspaceLogger(w.Id, logs.LogSourceServer)
	defer workspaceLogger.Close()

	wsLogWriter := io.MultiWriter(&util.InfoLogWriter{}, workspaceLogger)

	err = s.startWorkspace(ctx, w, target, wsLogWriter)

	if !telemetry.TelemetryEnabled(ctx) {
		return err
	}

	telemetryProps := telemetry.NewWorkspaceEventProps(w, target)
	event := telemetry.ServerEventWorkspaceStarted
	if err != nil {
		telemetryProps["error"] = err.Error()
		event = telemetry.ServerEventWorkspaceStartError
	}
	telemetryError := s.telemetryService.TrackServerEvent(event, workspaceId, telemetryProps)
	if telemetryError != nil {
		log.Trace(telemetryError)
	}

	return err
}

func (s *WorkspaceService) StartProject(ctx context.Context, workspaceId, projectName string) error {
	w, err := s.workspaceStore.Find(workspaceId)
	if err != nil {
		return ErrWorkspaceNotFound
	}

	project, err := w.GetProject(projectName)
	if err != nil {
		return ErrProjectNotFound
	}

	target, err := s.targetStore.Find(project.Target)
	if err != nil {
		return err
	}

	projectLogger := s.loggerFactory.CreateProjectLogger(w.Id, project.Name, logs.LogSourceServer)
	defer projectLogger.Close()

	return s.startProject(ctx, project, target, projectLogger)
}

func (s *WorkspaceService) startWorkspace(ctx context.Context, workspace *workspace.Workspace, target *provider.ProviderTarget, wsLogWriter io.Writer) error {
	wsLogWriter.Write([]byte("Starting workspace\n"))

	err := s.provisioner.StartWorkspace(workspace, target)
	if err != nil {
		return err
	}

	for _, project := range workspace.Projects {
		projectLogger := s.loggerFactory.CreateProjectLogger(workspace.Id, project.Name, logs.LogSourceServer)
		defer projectLogger.Close()

		err = s.startProject(ctx, project, target, projectLogger)
		if err != nil {
			return err
		}
	}

	wsLogWriter.Write([]byte(fmt.Sprintf("Workspace %s started\n", workspace.Name)))

	return nil
}

func (s *WorkspaceService) startProject(ctx context.Context, project *workspace.Project, target *provider.ProviderTarget, logWriter io.Writer) error {
	logWriter.Write([]byte(fmt.Sprintf("Starting project %s\n", project.Name)))

	projectToStart := *project
	projectToStart.EnvVars = workspace.GetProjectEnvVars(project, s.serverApiUrl, s.serverUrl, telemetry.TelemetryEnabled(ctx))

	err := s.provisioner.StartProject(project, target)
	if err != nil {
		return err
	}

	logWriter.Write([]byte(fmt.Sprintf("Project %s started\n", project.Name)))

	return nil
}
