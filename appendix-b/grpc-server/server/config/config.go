package config

import (
	"practical/appendix-b/grpc-server/server/telemetry"
	"github.com/rs/zerolog"
)

type AppConfig struct {
	Logger  zerolog.Logger
	Metrics telemetry.MetricReporter
}
