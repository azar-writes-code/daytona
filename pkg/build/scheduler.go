// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package build

import "github.com/robfig/cron/v3"

type IScheduler interface {
	Start()
	Stop()
	AddFunc(interval string, cmd func()) error
}

type Scheduler struct {
	cron *cron.Cron
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		cron: cron.New(cron.WithSeconds()),
	}
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}

func (s *Scheduler) AddFunc(interval string, cmd func()) error {
	_, err := s.cron.AddFunc(interval, cmd)
	return err
}
