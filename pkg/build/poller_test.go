// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package build_test

import (
	"testing"
	"time"

	t_build "github.com/daytonaio/daytona/internal/testing/server/build"
	"github.com/daytonaio/daytona/internal/testing/server/workspaces/mocks"
	"github.com/daytonaio/daytona/pkg/build"
	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/poller"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type PollerTestSuite struct {
	suite.Suite
	mockGitProviderService *mocks.MockGitProviderService
	mockBuilderFactory     *mocks.MockBuilderFactory
	mockBuilder            *mocks.MockBuilderPlugin
	mockScheduler          *mocks.MockSchedulerPlugin
	buildStore             build.Store
	Poller                 poller.IPoller
}

func NewPollerTestSuite() *PollerTestSuite {
	return &PollerTestSuite{}
}

func TestPoller(t *testing.T) {
	suite.Run(t, NewPollerTestSuite())
}

func (s *PollerTestSuite) SetupTest() {
	s.buildStore = t_build.NewInMemoryBuildStore()
	s.mockGitProviderService = mocks.NewMockGitProviderService()
	s.mockBuilderFactory = &mocks.MockBuilderFactory{}
	s.mockBuilder = &mocks.MockBuilderPlugin{}
	s.mockScheduler = &mocks.MockSchedulerPlugin{}
	s.Poller = build.NewPoller(build.PollerConfig{
		Scheduler:          s.mockScheduler,
		Interval:           "0 */5 * * * *",
		BuilderFactory:     s.mockBuilderFactory,
		BuildStore:         s.buildStore,
		GitProviderService: s.mockGitProviderService,
	})

	err := s.buildStore.Save(mocks.MockBuild)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *PollerTestSuite) TestStart() {
	s.mockScheduler.On("AddFunc", mock.Anything, mock.Anything).Return(nil)
	s.mockScheduler.On("Start").Return()

	require := s.Require()

	err := s.Poller.Start()
	require.NoError(err)

	s.mockScheduler.AssertExpectations(s.T())
}

func (s *PollerTestSuite) TestStop() {
	s.mockScheduler.On("Stop").Return()

	s.Poller.Stop()

	s.mockScheduler.AssertExpectations(s.T())
}

func (s *PollerTestSuite) TestPoll() {
	gpc := gitprovider.GitProviderConfig{}

	s.mockGitProviderService.On("GetConfigForUrl", mock.Anything).Return(&gpc, nil)
	s.mockBuilderFactory.On("Create", mock.Anything, gpc).Return(s.mockBuilder, nil)
	s.mockBuilder.On("Build").Return(mocks.MockBuild, nil)
	s.mockBuilder.On("Publish").Return(nil)
	s.mockBuilder.On("CleanUp").Return(nil)

	s.Poller.Poll()

	time.Sleep(100 * time.Millisecond)

	s.mockGitProviderService.AssertExpectations(s.T())
}
