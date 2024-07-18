// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package build

import (
	"fmt"

	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/poller"
	"github.com/daytonaio/daytona/pkg/scheduler"
	"github.com/daytonaio/daytona/pkg/server/gitproviders"
	log "github.com/sirupsen/logrus"
)

type PollerConfig struct {
	Scheduler          scheduler.IScheduler
	Interval           string
	BuilderFactory     IBuilderFactory
	BuildStore         Store
	GitProviderService gitproviders.IGitProviderService
}

type BuildPoller struct {
	poller.AbstractPoller
	builderFactory     IBuilderFactory
	buildStore         Store
	gitProviderService gitproviders.IGitProviderService
}

func NewPoller(config PollerConfig) *BuildPoller {
	poller := &BuildPoller{
		AbstractPoller:     *poller.NewPoller(config.Interval, config.Scheduler),
		builderFactory:     config.BuilderFactory,
		buildStore:         config.BuildStore,
		gitProviderService: config.GitProviderService,
	}
	poller.AbstractPoller.IPoller = poller

	return poller
}

func (p *BuildPoller) Poll() {
	builds, err := p.buildStore.FindAllByState(BuildStatePending)
	if err != nil {
		log.Error(err)
	}

	for _, build := range builds {
		go p.runBuildProcess(build)
	}
}

func (p *BuildPoller) runBuildProcess(build *Build) {
	if build.Project.Build == nil {
		return
	}

	gc, err := p.gitProviderService.GetConfigForUrl(build.Project.Repository.Url)
	if err != nil && !gitprovider.IsGitProviderNotFound(err) {
		log.Error(err)
		return
	}

	builder, err := p.builderFactory.Create(build.Project, gc)
	if err != nil {
		log.Error(err)
		return
	}

	build.State = BuildStateRunning
	err = p.buildStore.Save(build)
	if err != nil {
		log.Error(err)
	}

	result, err := builder.Build()
	if err != nil {
		p.handleBuildError(*build, builder, err)
		return
	}

	build = result

	err = builder.Publish()
	if err != nil {
		p.handleBuildError(*build, builder, err)
		return
	}

	err = p.buildStore.Save(build)
	if err != nil {
		log.Error(err)
	}

	err = builder.CleanUp()
	if err != nil {
		log.Error(fmt.Sprintf("Error cleaning up build: %s\n", err.Error()))
		return
	}
}

func (p *BuildPoller) handleBuildError(build Build, builder IBuilder, err error) {
	var errMsg string
	errMsg += "################################################\n"
	errMsg += fmt.Sprintf("#### BUILD FAILED FOR PROJECT %s: %s\n", build.Project.Name, err.Error())
	errMsg += "################################################\n"

	build.State = BuildStateFailure
	err = p.buildStore.Save(&build)
	if err != nil {
		errMsg += fmt.Sprintf("Error saving build: %s\n", err.Error())
	}

	cleanupErr := builder.CleanUp()
	if cleanupErr != nil {
		errMsg += fmt.Sprintf("Error cleaning up build: %s\n", cleanupErr.Error())
	}

	log.Error(errMsg)
}
