package config

import (
	"context"
	"database/sql"

	users "practical/appendix-b/grpc-server/service"
	"practical/appendix-b/http-server/telemetry"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
	"gocloud.dev/blob"
)

type AppConfig struct {
	PackageBucket *blob.Bucket
	Db            *sql.DB
	UsersSvc      users.UsersClient

	// telemetry
	Logger   zerolog.Logger
	Metrics  telemetry.MetricReporter
	Trace    trace.Tracer
	TraceCtx context.Context
	Span     trace.Span
	SpanCtx  context.Context
}
