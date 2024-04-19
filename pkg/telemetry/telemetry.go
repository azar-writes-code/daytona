// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package telemetry

import (
	"context"
	"io"

	"github.com/daytonaio/daytona/internal"
)

const ENABLED_HEADER = "X-Daytona-Telemetry-Enabled"
const SESSION_ID_HEADER = "X-Daytona-Session-Id"
const SOURCE_HEADER = "X-Daytona-Source"

type TelemetryContextKey string

var (
	ENABLED_CONTEXT_KEY TelemetryContextKey = "telemetry-enabled"
)

type TelemetryService interface {
	io.Closer
	TrackCliEvent(event CliEvent, sessionId string, properties map[string]interface{}) error
	TrackServerEvent(event ServerEvent, sessionId string, properties map[string]interface{}) error
	SetCommonProps(properties map[string]interface{})
}

func TelemetryEnabled(ctx context.Context) bool {
	enabled, ok := ctx.Value(ENABLED_CONTEXT_KEY).(bool)
	if !ok {
		return false
	}

	return enabled
}

type AbstractTelemetryService struct {
	daytonaVersion string
	TelemetryService
}

func NewAbstractTelemetryService() *AbstractTelemetryService {
	return &AbstractTelemetryService{
		daytonaVersion: internal.Version,
	}
}

func (t *AbstractTelemetryService) SetCommonProps(properties map[string]interface{}) {
	properties["daytona_version"] = t.daytonaVersion
}
