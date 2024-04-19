// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package middlewares

import (
	"context"
	"time"

	"github.com/daytonaio/daytona/internal"
	"github.com/daytonaio/daytona/pkg/telemetry"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func TelemetryMiddleware(telemetryService telemetry.TelemetryService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if telemetryService == nil {
			ctx.Next()
			return
		}

		if ctx.GetHeader(telemetry.ENABLED_HEADER) != "true" {
			ctx.Next()
			return
		}

		telemetryCtx := context.WithValue(ctx.Request.Context(), telemetry.ENABLED_CONTEXT_KEY, true)
		ctx.Request = ctx.Request.WithContext(telemetryCtx)

		sessionId := ctx.GetHeader(telemetry.SESSION_ID_HEADER)
		if sessionId == "" {
			sessionId = internal.SESSION_ID
		}

		source := ctx.GetHeader(telemetry.SOURCE_HEADER)

		reqMethod := ctx.Request.Method
		reqUri := ctx.FullPath()

		query := ctx.Request.URL.RawQuery

		err := telemetryService.TrackServerEvent(telemetry.ServerEventApiRequestStarted, sessionId, map[string]interface{}{
			"method": reqMethod,
			"URI":    reqUri,
			"query":  query,
			"source": source,
		})
		if err != nil {
			log.Trace(err)
		}

		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()
		execTime := endTime.Sub(startTime)
		statusCode := ctx.Writer.Status()

		properties := map[string]interface{}{
			"method":         reqMethod,
			"URI":            reqUri,
			"query":          query,
			"status":         statusCode,
			"source":         source,
			"exec time (µs)": execTime.Microseconds(),
		}

		if len(ctx.Errors) > 0 {
			properties["error"] = ctx.Errors.String()
		}

		err = telemetryService.TrackServerEvent(telemetry.ServerEventApiResponseSent, sessionId, properties)
		if err != nil {
			log.Trace(err)
		}

		ctx.Next()
	}
}
